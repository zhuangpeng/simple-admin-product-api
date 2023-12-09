package svc

import (
	"github.com/agui-coder/simple-admin-product-api/api/internal/config"
	"github.com/agui-coder/simple-admin-product-api/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
	}
}
