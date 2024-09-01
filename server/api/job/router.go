package job

import (
	"github.com/ginx-contribs/ginx"
)

type Router struct {
	Job *JobAPI
}

func NewRouter(root *ginx.RouterGroup, job *JobAPI) Router {
	group := root.Group("job")
	group.GET("/info", job.Info)
	group.GET("/list", job.List)
	group.GET("/start", job.Start)
	group.GET("/stop", job.Stop)

	return Router{Job: job}
}
