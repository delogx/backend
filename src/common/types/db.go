package types

import (
	"database/sql"

	"gorm.io/gorm"
)

type DB interface {
	Where(query interface{}, arg ...interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Select(query interface{}, args ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Begin(opts ...*sql.TxOptions) *gorm.DB
	Model(value interface{}) *gorm.DB
	Association(column string) *gorm.Association
}
