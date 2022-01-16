package model

import "gorm.io/gorm"

type FileContainer struct {
	gorm.Model
	FileId    uint
	Files     []File `gorm:"many2many:filecontainer_files;"`
	ReleaseId uint
}
