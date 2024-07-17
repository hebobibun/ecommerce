package interfaces

import "gorm.io/gorm"

type Database interface {
	OpenDB() *gorm.DB
	CloseDB(db *gorm.DB)
}
