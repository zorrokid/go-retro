package model

import "gorm.io/gorm"

type Release struct {
	gorm.Model
	Edition string
	TitleId uint
}
