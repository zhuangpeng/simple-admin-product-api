package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth     rest.AuthConf
	CROSConf config.CROSConf
}
