package database

import "github.com/jinzhu/gorm"

type SqlHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
}
