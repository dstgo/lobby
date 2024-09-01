package test

import (
	"context"
	"github.com/dstgo/lobby/server"
	"github.com/dstgo/lobby/server/data/ent"
	"github.com/dstgo/lobby/server/data/repo"
	"github.com/dstgo/lobby/test/testuitl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newJobRepo() (*repo.JobRepo, error) {
	ctx := context.Background()
	dbConf, err := testuitl.ReadDBConf()
	if err != nil {
		return nil, err
	}
	db, err := server.InitializeDB(ctx, dbConf)
	if err != nil {
		return nil, err
	}
	return repo.NewJobRepo(db.Debug()), nil
}

func TestUpsertOne(t *testing.T) {
	jobRepo, err := newJobRepo()
	if !assert.NoError(t, err) {
		return
	}
	samples := []*ent.CronJob{
		{Name: "1", Cron: "* * * 1 *", EntryID: 2, Prev: 3, Next: 4},
		{Name: "2", Cron: "* * * 1 *", EntryID: 4, Prev: 6, Next: 7},
		{Name: "3", Cron: "* * * 1 *", EntryID: 8, Prev: 9, Next: 10},
	}
	for _, sample := range samples {
		err = jobRepo.UpsertOne(context.Background(), sample)
		if !assert.NoError(t, err) {
			return
		}
	}

	samples2 := []*ent.CronJob{
		{Name: "1", Cron: "* * * 1 *", EntryID: 1, Prev: 2, Next: 3},
		{Name: "2", Cron: "* * * 1 *", EntryID: 1, Prev: 2, Next: 3},
		{Name: "3", Cron: "* * * 1 *", EntryID: 1, Prev: 2, Next: 3},
	}
	for _, sample := range samples2 {
		err = jobRepo.UpsertOne(context.Background(), sample)
		if !assert.NoError(t, err) {
			return
		}
	}

	// count of records must be not changed
	count, err := jobRepo.Ent.CronJob.Query().Count(context.Background())
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, len(samples), count)
}

func TestListByPage(t *testing.T) {
	ctx := context.Background()
	jobRepo, err := newJobRepo()
	if !assert.NoError(t, err) {
		return
	}
	list, total, err := jobRepo.ListByPage(ctx, 1, 10, "")
	if !assert.NoError(t, err) {
		return
	}
	t.Log(total)
	t.Log(list)
}
