package example

import (
	"go.uber.org/dig"

	"simon/limofy/service/internal/config"
	"simon/limofy/service/internal/repository"
	"simon/limofy/service/internal/thirdparty/localcache"
	"simon/limofy/service/internal/thirdparty/logger"
	"simon/limofy/service/internal/thirdparty/mysqlcli"
)

func NewExample(in digIn) digOut {
	self := &packet{
		in: in,
		digOut: digOut{
			GetExampleCommon: newGetExampleCommon(in),
		},
	}

	return self.digOut
}

type digIn struct {
	dig.In
	// 套件
	DB         mysqlcli.IMySQLClient
	AppConf    config.IAppConfig
	OpsConf    config.ServiceConfig
	AppLogger  logger.ILogger `name:"appLogger"`
	LocalCache localcache.ILocalCache

	// DB
	ExampleDao repository.IExampleDao

	// 業務邏輯
}

type digOut struct {
	GetExampleCommon IGetExampleCommon
}

type packet struct {
	in digIn

	digOut
}
