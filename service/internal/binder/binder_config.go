package binder

import (
	"go.uber.org/dig"

	"simon/limofy/service/internal/config"
)

func provideConfig(binder *dig.Container) {
	if err := binder.Provide(config.NewAppConfig); err != nil {
		panic(err)
	}

	if err := binder.Provide(config.NewServiceConfig); err != nil {
		panic(err)
	}
}
