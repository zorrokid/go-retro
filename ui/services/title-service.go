package services

import (
	"fmt"

	"github.com/zorrokid/go-retro/database"
	"github.com/zorrokid/go-retro/database/model"
	"github.com/zorrokid/go-retro/repository"
)

type TitleService struct {
	repo *repository.Repository
}

func NewTitleService() *TitleService {
	db := database.NewDatabase()
	db.InitDB()

	titleService := &TitleService{
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

func (service *TitleService) Update(title *model.Title) {
	fmt.Print("TitleService")
	service.repo.Update(*title)
}
