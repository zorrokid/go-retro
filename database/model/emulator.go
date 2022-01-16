package model

import "gorm.io/gorm"

type Emulator struct {
	gorm.Model
	SystemId             uint
	ExternalExecutableId uint
}
