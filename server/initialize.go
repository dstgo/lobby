package server

import (
	entsql "entgo.io/ent/dialect/sql"
	"errors"
	"github.com/dstgo/lobby/server/conf"
	"github.com/dstgo/lobby/server/data/ent"
	"github.com/dstgo/lobby/server/types"
	"github.com/gin-gonic/gin"
	"github.com/ginx-contribs/dbx"
	"github.com/ginx-contribs/ginx"
	"github.com/ginx-contribs/ginx/pkg/resp"
	"github.com/ginx-contribs/logx"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/wneessen/go-mail"
	"golang.org/x/net/context"
	"strings"

	// offline time zone database
	_ "time/tzdata"
)

// EnvProvider only use for wire injection
var EnvProvider = wire.NewSet(
	wire.FieldsOf(new(*types.Context), "AppConf"),
	wire.FieldsOf(new(*types.Context), "Ent"),
	wire.FieldsOf(new(*types.Context), "Redis"),
	wire.FieldsOf(new(*types.Context), "Router"),
	wire.FieldsOf(new(*types.Context), "Email"),
	wire.FieldsOf(new(*types.Context), "Lobby"),
	wire.FieldsOf(new(*conf.App), "Jwt"),
	wire.FieldsOf(new(*conf.App), "Email"),
)

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

// InitializeDB initialize database with ent
func InitializeDB(ctx context.Context, dbConf conf.DB) (*ent.Client, error) {
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

// InitializeRedis initialize redis connection
func InitializeRedis(ctx context.Context, redisConf conf.Redis) (*redis.Client, error) {
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

// InitializeEmail initialize email client
func InitializeEmail(ctx context.Context, emailConf conf.Email) (*mail.Client, error) {
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

func setupHumanizedValidator() error {
	v := validator.New()
	v.SetTagName("binding")
	englishValidator, err := ginx.EnglishValidator(v, validateParams)
	if err != nil {
		return err
	}
	ginx.SetValidator(englishValidator)
	ginx.SetValidateHandler(englishValidator.HandleError)
	return nil
}

func validateParams(ctx *gin.Context, val any, err error, translator ut.Translator) {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		var errorMsg []string
		for _, validateErr := range validationErrors {
			errorMsg = append(errorMsg, validateErr.Translate(translator))
		}
		verr := errors.New(strings.Join(errorMsg, ", "))
		resp.Fail(ctx).Code(types.ErrBadParams.Code).Error(verr).JSON()
		return
	}
	resp.Fail(ctx).Code(types.ErrBadParams.Code).ErrorMsg("params validate failed").JSON()
}
