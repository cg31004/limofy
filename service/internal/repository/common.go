package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"simon/limofy/service/internal/model/po"
)

func parsePaging(pager *po.Pager) func(dc *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pager != nil {
			db = db.Limit(pager.GetSize()).Offset(pager.GetOffset())
		}

		return db
	}
}

func setForUpdate(b bool) func(dc *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if b {
			db = db.Clauses(clause.Locking{Strength: "UPDATE"})
		}
		return db
	}
}
