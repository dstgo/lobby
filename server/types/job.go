package types

import "github.com/dstgo/lobby/server/data/ent"

type JobNameOptions struct {
	Name string `form:"name" binding:"required"`
}

type JobPageOption struct {
	Page   int    `form:"page"`
	Size   int    `form:"size"`
	Search string `form:"search"`
}

type JobInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Cron string `json:"cron"`
	Next int64  `json:"next"`
	Prev int64  `json:"prev"`
}

type JobPageList struct {
	Total int       `json:"total"`
	List  []JobInfo `json:"list"`
}

func EntJobToJobInfo(j *ent.CronJob) JobInfo {
	return JobInfo{
		Id:   j.EntryID,
		Name: j.Name,
		Cron: j.Cron,
		Next: j.Next,
		Prev: j.Prev,
	}
}

func EntJobToJobInfoBatch(js []*ent.CronJob) []JobInfo {
	var jis []JobInfo
	for _, j := range js {
		jis = append(jis, EntJobToJobInfo(j))
	}
	return jis
}
