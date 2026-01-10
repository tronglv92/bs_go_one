package config

import (
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

func Load(file *string) Config {
	flag.Parse()
	var c Config
	conf.MustLoad(*file, &c, conf.UseEnv())
	return c
}

type Config struct {
	Server ServerConfig `json:"server,optional"`
}

func (c Config) ServiceName() string {
	return c.Server.Http.Name
}

type ServerConfig struct {
	Id      int           `json:",default=0,optional"`
	Env     string        `json:",default=production,optional"`
	Http    rest.RestConf `json:"http,optional"`
	StatLog bool          `json:"stat-log,default=false"`
	LoadLog bool          `json:"load-log,default=false"`
}
