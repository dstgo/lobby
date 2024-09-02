package mq

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"
)

func NewStreamQueue(client *redis.Client) *StreamQueue {
	withContext, _ := errgroup.WithContext(context.Background())
	return &StreamQueue{
		redis:      client,
		subscribes: make(map[string][]Consumer),
		group:      withContext,
	}
}

// StreamQueue implement Queue interface by Redis Stream
type StreamQueue struct {
	redis *redis.Client
	// ready-only map
	subscribes map[string][]Consumer

	running atomic.Bool
	group   *errgroup.Group
	once    sync.Once
	close   atomic.Bool
}

func (q *StreamQueue) Subscribe(consumer Consumer) error {
	// prevent from concurrent writes after running
	if q.running.Load() {
		return errors.New("consumer subscribe after stream queue already running")
	}
	q.subscribes[consumer.Topic()] = append(q.subscribes[consumer.Topic()], consumer)
	return nil
}

func (q *StreamQueue) Publish(ctx context.Context, topic string, msg any, maxLen int64) (id string, err error) {
	result, err := q.redis.XAdd(ctx, &redis.XAddArgs{
		Stream: topic,
		MaxLen: maxLen,
		Values: msg,
		ID:     "*",
	}).Result()
	return result, err
}

func (q *StreamQueue) Start(ctx context.Context) {
	q.once.Do(func() {
		for _, consumers := range q.subscribes {
			for _, consumer := range consumers {
				q.group.Go(func() error {
					return q.consume(ctx, consumer.Topic(), consumer.Group(), consumer.Name(), consumer.Size(), consumer)
				})
			}
		}
	})
}

func (q *StreamQueue) Close() error {
	q.close.Store(true)
	return q.group.Wait()
}

func (q *StreamQueue) consume(ctx context.Context, topic, group, consumer string, batchSize int64, cb Consumer) error {
	defer func() {
		if err := recover(); err != nil {
			slog.Error("stream panic recovered", slog.Any("error", err))
		}
	}()
	// create the consumer group
	stream := q.redis.XGroupCreateMkStream(ctx, topic, group, "0")
	if stream.Err() != nil && stream.Err().Error() != "BUSYGROUP Consumer Group name already exists" {
		return stream.Err()
	}

	slog.Debug(fmt.Sprintf("consumer %q is running", consumer), slog.String("topic", topic), slog.String("group", group))

	// consume messages in a loop
	for {
		// quit if queue was closed
		if q.close.Load() {
			return nil
		}

		// read the latest message
		if id, err := q.readStream(ctx, topic, group, consumer, ">", batchSize, cb); err != nil {
			errorLog("stream read latest failed", err, id, topic, group, consumer)
		}

		// read the messages that received but not ack
		if id, err := q.readStream(ctx, topic, group, consumer, "1", batchSize, cb); err != nil {
			errorLog("stream read not-ack failed", err, id, topic, group, consumer)
		}

		// clear dead messages in pending list
		if err := q.clearDead(ctx, topic, group, time.Minute*5, 10); err != nil {
			slog.Error("stream clear dead failed", slog.Any("error", err))
		}

		time.Sleep(1)
	}
}

func (q *StreamQueue) readStream(ctx context.Context, topic, group, consumer, id string, batchSize int64, cb Consumer) (errorId string, err error) {
	// read from specified stream in specified group
	result, err := q.redis.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    group,
		Consumer: consumer,
		Streams:  []string{topic, id},
		Count:    batchSize,
	}).Result()

	if err != nil {
		return "", err
	}

	for _, stream := range result {
		topic := stream.Stream
		for _, message := range stream.Messages {
			if err := cb.Consume(ctx, message.ID, message.Values); err != nil {
				return message.ID, err
			} else { // make sure message is consumed if callback executed successfully
				if err := q.redis.XAck(ctx, topic, group, message.ID).Err(); err != nil {
					return message.ID, err
				}

				// del it if ack ok
				if err := q.redis.XDel(ctx, topic, message.ID).Err(); err != nil {
					return message.ID, err
				}
			}
		}
	}

	return "", nil
}

// clear dead msg that idle timeout
func (q *StreamQueue) clearDead(ctx context.Context, topic, group string, idle time.Duration, count int64) error {
	pel, err := q.redis.XPendingExt(ctx, &redis.XPendingExtArgs{
		Stream: topic,
		Group:  group,
		Idle:   idle,
		Start:  "-",
		End:    "+",
		Count:  count,
	}).Result()

	if err != nil {
		return err
	}

	var ids []string
	for _, pending := range pel {
		ids = append(ids, pending.ID)
	}

	if len(ids) == 0 {
		return nil
	}

	/// delete msg
	if _, err := q.redis.XDel(ctx, topic, ids...).Result(); err != nil {
		return err
	}

	return nil
}

func errorLog(msg string, err error, id, topic, group, consumer string) {
	slog.Error(msg,
		slog.String("error", err.Error()),
		slog.String("msg-id", id),
		slog.String("topic", topic),
		slog.String("group", group),
		slog.String("consumer", consumer),
	)
}
