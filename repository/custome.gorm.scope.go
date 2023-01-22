package repository

import (
	"math"

	"github.com/fxmbx/go_practice_pr/utils"
	"gorm.io/gorm"
)

func Paginate(model any, pagination *utils.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var (
		totalRows  int64
		totalPages int32
	)

	db.Model(model).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages = int32(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(pagination.GetOffset())).Limit(int(pagination.GetLimit())).Order(pagination.GetSort())
	}
}
