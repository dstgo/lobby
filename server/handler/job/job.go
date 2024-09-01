package job

import (
	"context"
	"errors"
	"github.com/dstgo/lobby/server/data/ent"
	"github.com/dstgo/lobby/server/data/repo"
	"github.com/dstgo/lobby/server/jobs"
	"github.com/dstgo/lobby/server/types"
	"github.com/ginx-contribs/ginx/pkg/resp/statuserr"
)

func NewJobHandler(jobrepo *repo.JobRepo, cronjob *jobs.CronJob) *JobHandler {
	return &JobHandler{jobRepo: jobrepo, cronjob: cronjob}
}

type JobHandler struct {
	jobRepo *repo.JobRepo
	cronjob *jobs.CronJob
}

func (j *JobHandler) List(ctx context.Context, page int, size int, search string) (types.JobPageList, error) {
	list, total, err := j.jobRepo.ListByPage(ctx, page, size, search)
	if err != nil {
		return types.JobPageList{}, statuserr.InternalError(err)
	}
	return types.JobPageList{
		Total: total,
		List:  types.EntJobToJobInfoBatch(list),
	}, err
}

func (j *JobHandler) GetOne(ctx context.Context, name string) (types.JobInfo, error) {
	one, err := j.jobRepo.QueryOne(ctx, name)
	if err != nil {
		return types.JobInfo{}, statuserr.InternalError(err)
	}
	return types.EntJobToJobInfo(one), nil
}

// Stop removes the job from the future scheduled jobs list
func (j *JobHandler) Stop(ctx context.Context, name string) error {
	job, e := j.cronjob.GetJob(name)
	if !e {
		return errors.New("job not found")
	}
	// remove the job from the future scheduler
	j.cronjob.DelJob(name)
	// update information
	err := j.jobRepo.UpsertOne(ctx, &ent.CronJob{
		Cron:    job.Cron(),
		EntryID: int(job.ID),
		Prev:    job.Prev.UnixMicro(),
		Next:    -1,
	})
	if err != nil {
		return err
	}
	return nil
}

// Start add job into future scheduled jobs list
func (j *JobHandler) Start(ctx context.Context, name string) error {
	err := j.cronjob.ContinueJob(name)
	if err != nil {
		return err
	}
	job, _ := j.cronjob.GetJob(name)

	// update information
	err = j.jobRepo.UpsertOne(ctx, &ent.CronJob{
		Cron:    job.Cron(),
		EntryID: int(job.ID),
		Prev:    job.Prev.UnixMicro(),
		Next:    job.Next.UnixMicro(),
	})
	if err != nil {
		return err
	}
	return nil
}
