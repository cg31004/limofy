package binder

import (
	"go.uber.org/dig"

	"simon/limofy/service/internal/thirdparty/localcache"
	"simon/limofy/service/internal/thirdparty/logger"
	"simon/limofy/service/internal/thirdparty/mysqlcli"
	"simon/limofy/service/internal/thirdparty/redisclient"
)

func provideThirdParty(binder *dig.Container) {
	if err := binder.Provide(logger.NewAppLogger, dig.Name("appLogger")); err != nil {
		panic(err)
	}

	if err := binder.Provide(logger.NewSysLogger, dig.Name("sysLogger")); err != nil {
		panic(err)
	}

	if err := binder.Provide(mysqlcli.NewDBClient); err != nil {
		panic(err)
	}

	if err := binder.Provide(redisclient.NewRedisClient); err != nil {
		panic(err)
	}

	if err := binder.Provide(localcache.NewDefault); err != nil {
		panic(err)
	}

}
