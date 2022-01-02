package database

import (
	"github.com/zorrokid/go-retro/database/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("go-retro.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

type Database struct {
	Connection *gorm.DB
}

func NewDatabase() *Database {
	db := &Database{Connection: connect()}
	return db
}

func (db *Database) InitDB() {
	db.Connection.AutoMigrate(&model.Title{})
}
