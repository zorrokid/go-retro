package model

import "gorm.io/gorm"

type ExternalExecutable struct {
	gorm.Model
	Command string
	Type    string
}
