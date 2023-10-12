package query

import (
	"errors"

	"github.com/wardonne/gopi/model/pagination"
	"github.com/wardonne/gopi/support/collection/list"
	"gorm.io/gorm"
)

func First[T any](db *gorm.DB, conditions ...any) (value T) {
	var model T
	if err := db.First(model, conditions...).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		} else {
			panic(err)
		}
	} else {
		return model
	}
}

func FirstOrFail[T any](db *gorm.DB, conditions ...any) (value T) {
	var model T
	if err := db.First(model, conditions...).Error; err != nil {
		panic(err)
	} else {
		return model
	}
}

func FirstOrCreate[T any](db *gorm.DB, conditions ...any) (value T) {
	var model T
	if err := db.FirstOrCreate(model, conditions...).Error; err != nil {
		panic(err)
	} else {
		return model
	}
}

func FirstOrInit[T any](db *gorm.DB, conditions ...any) (value T) {
	var model T
	if err := db.FirstOrInit(model, conditions...); err != nil {
		panic(err)
	} else {
		return model
	}
}

func Last[T any](db *gorm.DB, conditions ...any) (value T) {
	var model T
	if err := db.Last(model, conditions...).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		} else {
			panic(err)
		}
	} else {
		return model
	}
}

func LastOrFail[T any](db *gorm.DB, conditions ...any) (value T) {
	var model T
	if err := db.Last(model, conditions...).Error; err != nil {
		panic(err)
	} else {
		return model
	}
}

func Take[T any](db *gorm.DB, conditions ...any) (value T) {
	var model T
	if err := db.Take(model, conditions...).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		} else {
			panic(err)
		}
	} else {
		return model
	}
}

func TakeOrFail[T any](db *gorm.DB, conditions ...any) (value T) {
	var model T
	if err := db.Take(model, conditions...).Error; err != nil {
		panic(err)
	} else {
		return model
	}
}

func Find[T any](db *gorm.DB, conditions ...any) *list.ArrayList[T] {
	var models = []T{}
	if err := db.Find(&models, conditions...).Error; err != nil {
		panic(err)
	} else {
		return list.NewArrayList[T](models...)
	}
}

func Chunk[T any](db gorm.DB, batchSize int, callback func(tx *gorm.DB, batch int) error) *list.ArrayList[T] {
	var allModels = []T{}
	var models = []T{}
	if err := db.FindInBatches(&models, batchSize, func(tx *gorm.DB, batch int) error {
		if err := callback(tx, batch); err != nil {
			return err
		} else {
			allModels = append(allModels, models...)
			return nil
		}
	}).Error; err != nil {
		panic(err)
	} else {
		return list.NewArrayList[T](allModels...)
	}
}

func Count(db *gorm.DB) int64 {
	var total int64
	if err := db.Count(&total).Error; err != nil {
		panic(err)
	} else {
		return total
	}
}

func Exists(db *gorm.DB, conditions ...any) bool {
	return Count(db) > 0
}

func Paginate[T any](db *gorm.DB, pageSize, page int) *pagination.QueryPaginator[T] {
	return pagination.NewQueryPaginator[T](db, pageSize, page)
}
