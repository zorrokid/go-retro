package model

import "gorm.io/gorm"

type EmulatorConfig struct {
	gorm.Model
	EmulatorId uint
	FileType   string
}
