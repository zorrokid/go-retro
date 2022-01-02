package repository

import (
	"github.com/zorrokid/go-retro/database"
	"github.com/zorrokid/go-retro/database/model"
)

type Repository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *Repository {
	repo := &Repository{db: db}
	return repo
}

// methods for Repository

func (r *Repository) AddTitle(name string) {
	r.db.Connection.Create(&model.Title{Name: name})
}

func (r *Repository) GetTitle(id int) model.Title {
	var title model.Title
	r.db.Connection.First(&title, id)
	return title
}

func (r *Repository) GetTitleByName(name string) model.Title {
	var title model.Title
	r.db.Connection.First(&title, "name = ?", name)
	return title
}
