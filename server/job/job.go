package job

import "log/slog"

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

// Job is representation of a cron job
type Job interface {
	// Name returns the name of the job
	Name() string
	// Cron returns the cron expression for the job
	Cron() string
	// Cmd returns the command to be executed by the job
	Cmd() func()
}

var _GlobalJobs = make(globalJobs)

type globalJobs map[string]Job

func (g globalJobs) Register(j Job) {
	g[j.Name()] = j
}
