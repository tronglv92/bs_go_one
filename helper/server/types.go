package server

import (
	"go_one/helper/locale"
	"go_one/helper/server/core"

	"github.com/zeromicro/go-zero/rest"
)

type RestHandler interface {
	Register(svr *rest.Server)
}

func Providers() []core.Service {
	return []core.Service{
		locale.NewLocalizer(),
	}
}
