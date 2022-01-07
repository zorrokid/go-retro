package services

import (
	"github.com/zorrokid/go-retro/database"
	"github.com/zorrokid/go-retro/database/model"
	"github.com/zorrokid/go-retro/repository"
)

type TitleService struct {
	db   *database.Database
	repo *repository.Repository
}

func NewTitleService(db *database.Database) *TitleService {
	titleService := &TitleService{
		db:   db,
		repo: repository.NewRepository(db),
	}
	return titleService
}

func (service *TitleService) AddTitle(titleName string) {
	service.repo.AddTitle(titleName)
}

func (service *TitleService) GetTitles() []model.Title {
	return service.repo.GetTitles()
}
