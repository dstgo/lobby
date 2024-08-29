package conf

import (
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestReadFrom(t *testing.T) {
	filename := "testdata/conf.toml"
	cfg := App{Server: Server{Address: "127.0.0.1:8080"}, Log: Log{Level: slog.LevelDebug}}
	err := WriteTo(filename, cfg)
	assert.NoError(t, err)
	app, err := ReadFrom(filename)
	assert.NoError(t, err)
	assert.Equal(t, app.Server.Address, cfg.Server.Address)
}

func TestRevise(t *testing.T) {
	cfg := App{Server: Server{Address: "127.0.0.1:8080"}, Log: Log{Level: slog.LevelDebug}}
	reviseConf, err := Revise(cfg)
	assert.NoError(t, err)
	t.Log(reviseConf)
}
