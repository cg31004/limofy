package repository

import (
	"context"

	"gorm.io/gorm"

	"simon/limofy/service/internal/model/po"
	"simon/limofy/service/internal/thirdparty/errortool"
	"simon/limofy/service/internal/utils/timelogger"
)

type IExampleDao interface {
	List(ctx context.Context, db *gorm.DB, cond *po.ExampleCondByGet, pager *po.Pager) ([]*po.Example, error)
	ListPager(ctx context.Context, db *gorm.DB, cond *po.ExampleCondByGet, pager *po.Pager) (*po.PagingResult, error)
	ListWithoutPager(ctx context.Context, db *gorm.DB, cond *po.ExampleCondByGet) ([]*po.Example, error)
}

type exampleDao struct {
	in repositoryIn
}

func newExampleDao(in repositoryIn) IExampleDao {
	return &exampleDao{in: in}
}

func (dao *exampleDao) List(ctx context.Context, db *gorm.DB, cond *po.ExampleCondByGet, pager *po.Pager) ([]*po.Example, error) {
	defer timelogger.LogTime(ctx)()

	var example []*po.Example

	if err := db.
		Scopes(dao.list(cond, pager)).
		Scan(&example).
		Error; err != nil {
		return nil, errortool.ConvertDB(err)
	}

	return example, nil
}

func (dao *exampleDao) ListPager(ctx context.Context, db *gorm.DB, cond *po.ExampleCondByGet, pager *po.Pager) (*po.PagingResult, error) {
	defer timelogger.LogTime(ctx)()

	var count int64

	if err := db.
		Scopes(dao.list(cond, pager)).
		Count(&count).
		Error; err != nil {
		return nil, errortool.ConvertDB(err)
	}

	return po.NewPagerResult(pager, count), nil
}
func (dao *exampleDao) ListWithoutPager(ctx context.Context, db *gorm.DB, cond *po.ExampleCondByGet) ([]*po.Example, error) {
	defer timelogger.LogTime(ctx)()

	var example []*po.Example

	if err := db.
		Scopes(dao.list(cond, nil)).
		Scan(&example).
		Error; err != nil {
		return nil, errortool.ConvertDB(err)
	}

	return example, nil
}

func (dao *exampleDao) list(cond *po.ExampleCondByGet, pager *po.Pager) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if cond.ID != "" {
			db = db.Where("example.id = ?", cond.ID)
		}

		db = db.Model(&po.Example{})

		// page condition
		if pager != nil {
			db = db.Scopes(parsePaging(pager))
		}

		return db
	}
}
