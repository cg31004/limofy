package binder

import (
	"go.uber.org/dig"

	"simon/limofy/service/internal/core/common/example"
)

func provideCoreCommon(binder *dig.Container) {
	if err := binder.Provide(example.NewExample); err != nil {
		panic(err)
	}
}
