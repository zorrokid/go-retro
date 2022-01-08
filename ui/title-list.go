package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/zorrokid/go-retro/database/model"
	"github.com/zorrokid/go-retro/ui/services"
)

type TitleList struct {
	titleService  *services.TitleService
	data          []model.Title
	selectedTitle *SelectedTitle
	list          *widget.List
}

func NewTitleList(titleService *services.TitleService, window *fyne.Window) *TitleList {
	data := titleService.GetTitles()
	selectedTitle := NewSelectedTitle(&data[0], window, titleService)
	list := &TitleList{
		titleService:  titleService,
		data:          data,
		selectedTitle: selectedTitle,
	}
	return list
}

func (list *TitleList) MakeList() fyne.CanvasObject {
	titleContent := list.selectedTitle.makeSelectedTitleContent()

	list.list = widget.NewList(
		func() int {
			return len(list.data)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(list.data[id].Name)
		},
	)
	list.list.OnSelected = func(id widget.ListItemID) {
		list.selectedTitle.Update(&list.data[id])
	}
	list.list.OnUnselected = func(id widget.ListItemID) {
		list.selectedTitle.Clear()
	}

	return container.NewHSplit(list.list, titleContent)
}

func (list *TitleList) Update() {
	list.data = list.titleService.GetTitles()
	list.list.Refresh()
}
