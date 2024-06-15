package types

import "gorm.io/gorm"

type DB interface {
	Where(query interface{}, arg ...interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
}
