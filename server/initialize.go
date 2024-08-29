package server

import (
	entsql "entgo.io/ent/dialect/sql"
	"errors"
	"github.com/dstgo/lobby/assets"
	"github.com/dstgo/lobby/server/conf"
	"github.com/dstgo/lobby/server/data/ent"
	"github.com/dstgo/lobby/server/types"
	"github.com/gin-gonic/gin"
	"github.com/ginx-contribs/dbx"
	"github.com/ginx-contribs/ginx/pkg/resp"
	"github.com/ginx-contribs/logx"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/wneessen/go-mail"
	"golang.org/x/net/context"
	"io"

	// offline time zone database
	_ "time/tzdata"
)

// EnvProvider only use for wire injection
var EnvProvider = wire.NewSet(
	wire.FieldsOf(new(*types.Env), "AppConf"),
	wire.FieldsOf(new(*types.Env), "Ent"),
	wire.FieldsOf(new(*types.Env), "Redis"),
	wire.FieldsOf(new(*types.Env), "Router"),
	wire.FieldsOf(new(*types.Env), "Email"),
	wire.FieldsOf(new(*conf.App), "Jwt"),
	wire.FieldsOf(new(*conf.App), "Email"),
)

// PrintBanner prints the banner into given writer
func PrintBanner(writer io.Writer) error {
	bytes, err := assets.FS.ReadFile("banner.txt")
	if err != nil {
		return err
	}
	_, err = writer.Write(bytes)
	return err
}

// NewLogger returns a new app logger with the given options
func NewLogger(option conf.Log) (*logx.Logger, error) {

	writer, err := logx.NewWriter(&logx.WriterOptions{
		Filename: option.Filename,
	})
	if err != nil {
		return nil, err
	}
	handler, err := logx.NewHandler(writer, &logx.HandlerOptions{
		Level:       option.Level,
		Format:      option.Format,
		Prompt:      option.Prompt,
		Source:      option.Source,
		ReplaceAttr: nil,
		Color:       option.Color,
	})
	if err != nil {
		return nil, err
	}
	logger, err := logx.New(
		logx.WithHandlers(handler),
	)
	if err != nil {
		return nil, err
	}
	return logger, nil
}

// handler to process when params validating failed
func validatePramsHandler(ctx *gin.Context, val any, err error) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		ctx.Error(err)
		resp.Fail(ctx).Error(types.ErrBadParams).JSON()
	}
}

// initialize database with ent
func initializeDB(ctx context.Context, dbConf conf.DB) (*ent.Client, error) {
	sqldb, err := dbx.Open(dbx.Options{
		Driver:             dbConf.Driver,
		Address:            dbConf.Address,
		User:               dbConf.User,
		Password:           dbConf.Password,
		Database:           dbConf.Database,
		Params:             dbConf.Params,
		MaxIdleConnections: dbConf.MaxIdleConnections,
		MaxOpenConnections: dbConf.MaxOpenConnections,
		MaxLifeTime:        dbConf.MaxLifeTime,
		MaxIdleTime:        dbConf.MaxIdleTime,
	})
	if err != nil {
		return nil, err
	}
	entClient := ent.NewClient(
		ent.Driver(entsql.OpenDB(dbConf.Driver, sqldb)),
	)
	// migrate database
	if err := entClient.Schema.Create(ctx); err != nil {
		return nil, err
	}

	return entClient, err
}

// initialize redis connection
func initializeRedis(ctx context.Context, redisConf conf.Redis) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         redisConf.Address,
		Password:     redisConf.Password,
		ReadTimeout:  redisConf.ReadTimeout,
		WriteTimeout: redisConf.WriteTimeout,
	})
	pingResult := redisClient.Ping(ctx)
	if pingResult.Err() != nil {
		return nil, pingResult.Err()
	}

	return redisClient, nil
}

// initialize email client
func initializeEmail(ctx context.Context, emailConf conf.Email) (*mail.Client, error) {
	client, err := mail.NewClient(emailConf.Host,
		mail.WithPort(emailConf.Port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(emailConf.Username),
		mail.WithPassword(emailConf.Password),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}
