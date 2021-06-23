// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/jursonmo/geektask/week04/blog/internal/biz"
	"github.com/jursonmo/geektask/week04/blog/internal/dao"
	"github.com/jursonmo/geektask/week04/blog/internal/server"
	"github.com/jursonmo/geektask/week04/blog/internal/service"
)
//newApp -->依赖-->server.ProviderSet-->依赖-->service.ProviderSet.....
func initApp(logger log.Logger) (*kratos.App, error) {
	panic(wire.Build(newApp, server.ProviderSet, service.ProviderSet, biz.ProviderSet, dao.ProvideSet))
}
