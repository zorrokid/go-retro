package model

import "gorm.io/gorm"

type Title struct {
	gorm.Model
	Name     string
	Releases []Release
}
