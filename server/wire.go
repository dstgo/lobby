//go:build wireinject

// The build tag makes sure the stub is not built in the final build
package server

import (
	"github.com/dstgo/lobby/server/api"
	"github.com/dstgo/lobby/server/data"
	"github.com/dstgo/lobby/server/handler"
	"github.com/dstgo/lobby/server/jobs"
	"github.com/dstgo/lobby/server/svc"
	"github.com/dstgo/lobby/server/types"
	"github.com/google/wire"
)

// initialize and setup app environment
func setup(ctx types.Context) (svc.Context, error) {
	panic(wire.Build(ContextProvider, data.Provider, handler.Provider, api.Provider, jobs.Provider, svc.Provider))
}
