package binder

import (
	"go.uber.org/dig"

	"simon/limofy/service/internal/core/usecase/example"
)

func provideCoreUseCase(binder *dig.Container) {
	if err := binder.Provide(example.NewExample); err != nil {
		panic(err)
	}
}
