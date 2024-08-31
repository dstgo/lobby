package testuitl

import (
	"github.com/dstgo/lobby/server/conf"
	"log/slog"
	"os"
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
