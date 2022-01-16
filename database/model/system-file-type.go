package model

import "gorm.io/gorm"

type SystemFileType struct {
	gorm.Model
	SystemId        uint
	FileTypePattern string
}
