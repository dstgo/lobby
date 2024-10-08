package conf

import (
	"github.com/246859/duration"
	"log/slog"
)

// App is configuration for the whole application
type App struct {
	Server  Server        `toml:"server" comment:"http server configuration"`
	Log     Log           `toml:"log" comment:"server log configuration"`
	Elastic Elasticsearch `toml:"elastic" comment:"elasticsearch configuration"`
	DB      DB            `toml:"db" comment:"database connection configuration"`
	Redis   Redis         `toml:"redis" comment:"redis connection configuration"`
	Email   Email         `toml:"email" comment:"email smtp client configuration"`
	Jwt     Jwt           `toml:"jwt" comment:"jwt secret configuration"`
	Dst     Dst           `toml:"dst" comment:"dst configuration'"`
	Job     Job           `toml:"job" comment:"cron job configuration'"`

	Author    string `toml:"-" mapstructure:"-"`
	Version   string `toml:"-" mapstructure:"-"`
	BuildTime string `toml:"-" mapstructure:"-"`
}

// Server is configuration for the http server
type Server struct {
	Address      string            `toml:"address" comment:"server bind address"`
	BasePath     string            `toml:"basepath" comment:"base path for api"`
	ReadTimeout  duration.Duration `toml:"readTimeout" comment:"the maximum duration for reading the entire request"`
	WriteTimeout duration.Duration `toml:"writeTimeout" comment:"the maximum duration before timing out writes of the response"`
	IdleTimeout  duration.Duration `toml:"idleTimeout" comment:"the maximum amount of time to wait for the next request when keep-alives are enabled"`
	MultipartMax int64             `toml:"multipartMax" comment:"value of 'maxMemory' param that is given to http.Request's ParseMultipartForm"`
	Pprof        bool              `toml:"pprof" comment:"enabled pprof program profiling"`
	TLS          TLS               `toml:"tls" comment:"tls certificate"`
}

type TLS struct {
	Cert string `toml:"cert" comment:"tls certificate"`
	Key  string `toml:"key" comment:"tls key"`
}

// Log is configuration for logging
type Log struct {
	Filename string     `toml:"filename" comment:"log output file"`
	Prompt   string     `toml:"-"`
	Level    slog.Level `toml:"level" comment:"support levels: DEBUG | INFO | WARN | ERROR"`
	Format   string     `toml:"format" comment:"TEXT or JSON"`
	Source   bool       `toml:"source" comment:"whether to show source file in logs"`
	Color    bool       `toml:"color" comment:"enable color log"`
}

// DB is configuration for database
type DB struct {
	Driver             string            `toml:"driver" comment:"sqlite | mysql | postgresql"`
	Address            string            `toml:"address" comment:"db server host"`
	User               string            `toml:"user" comment:"db username"`
	Password           string            `toml:"password" comment:"db password"`
	Database           string            `toml:"database" comment:"database name"`
	Params             string            `toml:"param" comment:"connection params"`
	MaxIdleConnections int               `toml:"maxIdleConnections" comment:"max idle connections limit"`
	MaxOpenConnections int               `toml:"maxOpenConnections" comment:"max opening connections limit"`
	MaxLifeTime        duration.Duration `toml:"maxLifeTime" comment:"max connection lifetime"`
	MaxIdleTime        duration.Duration `toml:"maxIdleTime" comment:"max connection idle time"`
}

// Redis is configuration for redis server
type Redis struct {
	Address      string            `toml:"address" comment:"host address"`
	Password     string            `toml:"password" comment:"redis auth"`
	WriteTimeout duration.Duration `toml:"writeTimeout" comment:"Timeout for socket writes."`
	ReadTimeout  duration.Duration `toml:"readTimeout" comment:"Timeout for socket reads."`
}

// Jwt is configuration for jwt signing
type Jwt struct {
	Issuer  string       `toml:"issuer" comment:"jwt issuer"`
	Access  AccessToken  `toml:"access" comment:"access token configuration"`
	Refresh RefreshToken `toml:"refresh" comment:"refresh token configuration"`
}

type AccessToken struct {
	Expire duration.Duration `toml:"expire" comment:"duration to expire access token"`
	Delay  duration.Duration `toml:"delay" comment:"delay duration after expiration"`
	Key    string            `toml:"key" comment:"access token signing key"`
}

type RefreshToken struct {
	Expire duration.Duration `toml:"expire" comment:"duration to expire refresh token"`
	Key    string            `toml:"key" comment:"refresh token signing key"`
}

type RateLimit struct {
	Public struct {
		Limit  int               `toml:"limit"`
		Window duration.Duration `toml:"window"`
	} `toml:"public"`
}

type Email struct {
	Host     string     `toml:"host" comment:"smtp server host"`
	Port     int        `toml:"port" comment:"smtp server port"`
	Username string     `toml:"username" comment:"smtp user name"`
	Password string     `toml:"password" comment:"password to authenticate"`
	MQ       EmailMq    `toml:"-"`
	Code     VerifyCode `toml:"code" comment:"email verification code configuration"`
}

type EmailMq struct {
	Topic     string   `toml:"topic" comment:"email mq topic"`
	BatchSize int64    `toml:"batchSize" comment:"max batch size of per reading"`
	Group     string   `toml:"group" comment:"consumer group"`
	Consumers []string `toml:"consumers" comment:"how many consumer in groups, must >=1."`
}

type VerifyCode struct {
	TTL      duration.Duration `toml:"ttl" comment:"lifetime for verification code"`
	RetryTTL duration.Duration `toml:"retry" comment:"max wait time before asking for another new verification code"`
}

type Dst struct {
	ProxyUrl  string `toml:"proxyURL" comment:"proxy URL for http client"`
	SteamKey  string `toml:"steamKey" comment:"steam web api key"`
	KeliToken string `toml:"kleiToken" comment:"klei cluster server token"`
}

type Job struct {
	Collect Collect `toml:"collect" comment:"lobby collect jobs configuration"`
	Clean   Clean   `toml:"clean" comment:"lobby clean jobs configuration"`
}

type Collect struct {
	Cron      string `toml:"cron" comment:"collect jobs cron expression"`
	Limit     int    `toml:"limit" comment:"max goroutine limit"`
	BatchSize int    `toml:"batchSize" comment:"batch size for insertion"`
}

type Clean struct {
	Cron      string            `toml:"cron" comment:"clean jobs cron expression"`
	Expired   duration.Duration `toml:"expired" comment:"max lifetime for records"`
	BatchSize int               `toml:"batchSize" comment:"batch size for clean"`
}

type Elasticsearch struct {
	Enabled       bool   `toml:"enabled" comment:"whether to enable elasticsearch"`
	Address       string `toml:"address" comment:"address to connect to elasticsearch"`
	CAFingerprint string `toml:"caFingerprint" comment:"SHA256 hex fingerprint given by Elasticsearch"`
	Username      string `toml:"username" comment:"user for elasticsearch"`
	Password      string `toml:"password" comment:"password for user"`
}
