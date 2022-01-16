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
	db.Connection.AutoMigrate(&model.System{})
	db.Connection.AutoMigrate(&model.SystemFileType{})
	db.Connection.AutoMigrate(&model.ExternalExecutable{})
	db.Connection.AutoMigrate(&model.Emulator{})
	db.Connection.AutoMigrate(&model.EmulatorConfig{})
	db.Connection.AutoMigrate(&model.File{})
	db.Connection.AutoMigrate(&model.FileContainer{})
	db.Connection.AutoMigrate(&model.Release{})
	db.Connection.AutoMigrate(&model.Title{})
}
