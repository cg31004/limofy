package binder

import (
	"go.uber.org/dig"

	"simon/limofy/service/internal/controller"
	"simon/limofy/service/internal/controller/handler"
	"simon/limofy/service/internal/controller/middleware"
)

func provideController(binder *dig.Container) {
	if err := binder.Provide(handler.NewRequestParse); err != nil {
		panic(err)
	}

	if err := binder.Provide(middleware.NewResponseMiddleware); err != nil {
		panic(err)
	}

	if err := binder.Provide(controller.NewWebRestCtl); err != nil {
		panic(err)
	}
}
