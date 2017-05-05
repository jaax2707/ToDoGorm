package access

import "github.com/jinzhu/gorm"

type Db struct {
	DB *gorm.DB
}

func NewDb(DB *gorm.DB) *Db {
	return &Db{DB}
}
