package example

import (
	"context"

	"simon/limofy/service/internal/model/bo"
	"simon/limofy/service/internal/model/po"
	"simon/limofy/service/internal/utils/timelogger"
)

type IGetExampleCommon interface {
	Handle(ctx context.Context, cond *bo.ExampleGet) ([]*bo.Example, error)
}

func newGetExampleCommon(in digIn) IGetExampleCommon {
	return &getExampleCommon{in: in}
}

type getExampleCommon struct {
	in digIn
}

func (hd *getExampleCommon) Handle(ctx context.Context, cond *bo.ExampleGet) ([]*bo.Example, error) {
	defer timelogger.LogTime(ctx)()

	db := hd.in.DB.Session()

	result, err := hd.in.ExampleDao.ListWithoutPager(ctx, db, &po.ExampleCondByGet{ID: ""})
	if err != nil {
		hd.in.AppLogger.Error(ctx, err)
		return nil, err
	}

	resp := make([]*bo.Example, 0, len(result))
	for _, r := range result {
		resp = append(resp, &bo.Example{
			ID:        r.ID,
			Name:      r.Name,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}

	return resp, nil
}
