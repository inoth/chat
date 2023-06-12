package dbutil

import (
	"fmt"

	"gorm.io/gorm"
)

func Select[T any](db *gorm.DB, filter T) (*T, error) {
	var res T
	err := db.Model(filter).Where(filter).First(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func SelectList[T any](db *gorm.DB, filter T, funcs ...func(*gorm.DB) *gorm.DB) ([]T, error) {
	var res []T
	err := db.Model(filter).Where(filter).Scopes(funcs...).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SelectListAndTotal[T any](db *gorm.DB, filter T, funcs ...func(*gorm.DB) *gorm.DB) ([]T, int64, error) {
	var res []T
	q := db.Model(filter).Where(filter)
	var total int64
	q.Count(&total)
	if total <= 0 {
		return nil, 0, fmt.Errorf("not found data")
	}
	err := q.Scopes(funcs...).Scan(&res).Error
	if err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func Update[T any](db *gorm.DB, filter T, updates map[string]interface{}) error {
	return db.Model(filter).Where(filter).Updates(updates).Error
}

func Delete[T any](db *gorm.DB, filter T) error {
	return db.Model(filter).Where(filter).Delete(nil).Error
}

func Paginate(page int, limit int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
