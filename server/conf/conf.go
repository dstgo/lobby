package conf

import (
	"log/slog"
	"time"
)

// App is configuration for the whole application
type App struct {
	Server Server    `toml:"server" comment:"http server configuration"`
	Log    Log       `toml:"log" comment:"server log configuration"`
	DB     DB        `toml:"db" comment:"database connection configuration"`
	Redis  Redis     `toml:"redis" comment:"redis connection configuration"`
	Email  Email     `toml:"email" comment:"email smtp client configuration"`
	Jwt    Jwt       `toml:"jwt" comment:"jwt secret configuration"`
	Limit  RateLimit `toml:"limit" comment:"rate limit configuration"`
	Dst    Dst       `toml:"dst" comment:"dst configuration'"`

	Author    string `toml:"-" mapstructure:"-"`
	Version   string `toml:"-" mapstructure:"-"`
	BuildTime string `toml:"-" mapstructure:"-"`
}

// Server is configuration for the http server
type Server struct {
	Address      string        `toml:"address" comment:"server bind address"`
	ReadTimeout  time.Duration `toml:"readTimeout" comment:"the maximum duration for reading the entire request"`
	WriteTimeout time.Duration `toml:"writeTimeout" comment:"the maximum duration before timing out writes of the response"`
	IdleTimeout  time.Duration `toml:"idleTimeout" comment:"the maximum amount of time to wait for the next request when keep-alives are enabled"`
	MultipartMax int64         `toml:"multipartMax" comment:"value of 'maxMemory' param that is given to http.Request's ParseMultipartForm"`
	Pprof        bool          `toml:"pprof" comment:"enabled pprof program profiling"`
	TLS          TLS           `toml:"tls" comment:"tls certificate"`
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
	Driver             string        `toml:"driver" comment:"sqlite | mysql | postgresql"`
	Address            string        `toml:"address" comment:"db server host"`
	User               string        `toml:"user" comment:"db username"`
	Password           string        `toml:"password" comment:"db password"`
	Database           string        `toml:"database" comment:"database name"`
	Params             string        `toml:"param" comment:"connection params"`
	MaxIdleConnections int           `toml:"maxIdleConnections" comment:"max idle connections limit"`
	MaxOpenConnections int           `toml:"maxOpenConnections" comment:"max opening connections limit"`
	MaxLifeTime        time.Duration `toml:"maxLifeTime" comment:"max connection lifetime"`
	MaxIdleTime        time.Duration `toml:"maxIdleTime" comment:"max connection idle time"`
}

// Redis is configuration for redis server
type Redis struct {
	Address      string        `toml:"address" comment:"host address"`
	Password     string        `toml:"password" comment:"redis auth"`
	WriteTimeout time.Duration `toml:"writeTimeout" comment:"Timeout for socket writes."`
	ReadTimeout  time.Duration `toml:"readTimeout" comment:"Timeout for socket reads."`
}

// Jwt is configuration for jwt signing
type Jwt struct {
	Issuer  string       `toml:"issuer" comment:"jwt issuer"`
	Access  AccessToken  `toml:"access" comment:"access token configuration"`
	Refresh RefreshToken `toml:"refresh" comment:"refresh token configuration"`
}

type AccessToken struct {
	Expire time.Duration `toml:"expire" comment:"duration to expire access token"`
	Delay  time.Duration `toml:"delay" comment:"delay duration after expiration"`
	Key    string        `toml:"key" comment:"access token signing key"`
}

type RefreshToken struct {
	Expire time.Duration `toml:"expire" comment:"duration to expire refresh token"`
	Key    string        `toml:"key" comment:"refresh token signing key"`
}

type RateLimit struct {
	Public struct {
		Limit  int           `toml:"limit"`
		Window time.Duration `toml:"window"`
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
	MaxLen    int64    `toml:"maxLen" comment:"max length of topic"`
	BatchSize int64    `toml:"batchSize" comment:"max batch size of per reading"`
	Group     string   `toml:"group" comment:"consumer group"`
	Consumers []string `toml:"consumers" comment:"how many consumer in groups, must >=1."`
}

type VerifyCode struct {
	TTL      time.Duration `toml:"ttl" comment:"lifetime for verification code"`
	RetryTTL time.Duration `toml:"retry" comment:"max wait time before asking for another new verification code"`
}

type Dst struct {
	ProxyUrl    string `toml:"proxyURL" comment:"proxy URL for http client"`
	SteamKey    string `toml:"steamKey" comment:"steam web api key"`
	KeliToken   string `toml:"kleiToken" comment:"klei cluster server token"`
	CollectCron string `json:"collect_cron" comment:"collect cron for server collecting job"`
}
