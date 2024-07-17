package database

import (
	"log"
	"os"

	IF "github.com/hebobibun/go-ecommerce/interfaces"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct{}

func NewDB() IF.Database {
	return Database{}
}

func (Database) OpenDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DNS")), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	log.Println("Database connection established successfully")

	return db
}

func (Database) CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	_ = sqlDB.Close()
}
