package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	FileName        string
	CheckSum        []byte
	FileSizeInBytes int64
	FileType        string
}
