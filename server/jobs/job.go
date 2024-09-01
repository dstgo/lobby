package jobs

import (
	"errors"
	"github.com/google/wire"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/robfig/cron/v3"
	"log/slog"
)

type cronLogger struct {
	logger *slog.Logger
	prefab string
}

func (c cronLogger) Info(msg string, keysAndValues ...interface{}) {
	c.logger.Info(c.prefab+": "+msg, keysAndValues...)
}

func (c cronLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	c.logger.Error(c.prefab+": "+msg, append([]any{"err", err}, keysAndValues...)...)
}

// Job is representation of a cron jobs
type Job interface {
	// Name returns the name of the jobs
	Name() string
	// Cron returns the cron expression for the jobs
	Cron() string
	// Cmd returns the command to be executed by the jobs
	Cmd() func()
}

// RunningJob is representation of a running cron jobs
type RunningJob struct {
	Job
	cron.Entry
}

var Provider = wire.NewSet(
	NewCronJob,
)

func NewCronJob() *CronJob {
	logger := cronLogger{logger: slog.Default(), prefab: "cron-job-runtime-manager"}
	c := cron.New(cron.WithLogger(logger))
	return &CronJob{
		cron:    c,
		record:  cmap.New[Job](),
		running: cmap.New[RunningJob](),
	}
}

// CronJob is a simple cron jobs manager
type CronJob struct {
	cron    *cron.Cron
	record  cmap.ConcurrentMap[string, Job]
	running cmap.ConcurrentMap[string, RunningJob]
}

func (c *CronJob) AddJob(job Job) error {
	_, e := c.GetJob(job.Name())
	if e {
		return errors.New("same job already")
	}
	entryId, err := c.cron.AddFunc(job.Cron(), job.Cmd())
	if err != nil {
		return err
	}
	c.record.Set(job.Name(), job)
	c.running.Set(job.Name(), RunningJob{Entry: cron.Entry{ID: entryId}, Job: job})
	return nil
}

func (c *CronJob) GetJob(name string) (RunningJob, bool) {
	job, e := c.running.Get(name)
	if !e {
		return RunningJob{}, false
	}
	entry := c.cron.Entry(job.ID)
	if (entry != cron.Entry{}) {
		job.Entry = entry
	}
	return job, true
}

func (c *CronJob) DelJob(name string) {
	job, e := c.GetJob(name)
	if !e {
		return
	}
	c.cron.Remove(job.ID)
	c.running.Remove(name)
}

func (c *CronJob) do() {

}

func (c *CronJob) ContinueJob(name string) error {
	job, e := c.record.Get(name)
	if !e {
		return errors.New("job not found")
	}
	entryId, err := c.cron.AddFunc(job.Cron(), job.Cmd())
	if err != nil {
		return err
	}
	c.running.Set(job.Name(), RunningJob{Entry: cron.Entry{ID: entryId}, Job: job})
	return nil
}

func (c *CronJob) Start() {
	c.cron.Start()
}

func (c *CronJob) Stop() {
	c.cron.Stop()
}
