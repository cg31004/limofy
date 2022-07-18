package binder

import (
	"sync"

	"go.uber.org/dig"
)

var (
	binder *dig.Container
	once   sync.Once
)

func New() *dig.Container {
	once.Do(func() {
		binder = dig.New()

		provideApp(binder)
		provideConfig(binder)
		provideController(binder)
		provideCoreUseCase(binder)
		provideCoreCommon(binder)
		provideRepository(binder)
		provideThirdParty(binder)
	})

	return binder
}
