package example

import (
	"context"

	"simon/limofy/service/internal/model/bo"
)

type IGetExampleUseCase interface {
	Handle(ctx context.Context, cond *bo.ExampleGet) ([]*bo.Example, error)
}

func newGetExampleUseCase(in digIn) IGetExampleUseCase {
	return &getExampleUseCase{in: in}
}

type getExampleUseCase struct {
	in digIn
}

func (hd *getExampleUseCase) Handle(ctx context.Context, cond *bo.ExampleGet) ([]*bo.Example, error) {
	return hd.in.GetExampleCommon.Handle(ctx, cond)
}
