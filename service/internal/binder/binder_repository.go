package binder

import (
	"go.uber.org/dig"

	"simon/limofy/service/internal/repository"
)

func provideRepository(binder *dig.Container) {
	if err := binder.Provide(repository.NewRepository); err != nil {
		panic(err)
	}
}
