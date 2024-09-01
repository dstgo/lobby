package testutil

import (
	"github.com/dstgo/lobby/server/conf"
	"log/slog"
	"os"
	"time"
)

func init() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)
}

const configFile = "testdata/conf.toml"

func ReadConf() (conf.App, error) {
	appConf, err := conf.ReadFrom(configFile)
	if err != nil {
		return conf.App{}, err
	}
	return appConf, err
}

// ReadDBConf returns the test configuration
func ReadDBConf() (conf.DB, error) {
	appConf, err := conf.ReadFrom(configFile)
	if err != nil {
		return conf.DB{}, err
	}
	return appConf.DB, err
}

// Timer is helper to calculate cost-time
type Timer struct {
	start time.Time
}

func (t Timer) Start() {
	t.start = time.Now()
}

func (t Timer) Stop() time.Duration {
	return t.start.Sub(time.Now())
}

func (t Timer) Reset() {
	t.start = time.Time{}
}
