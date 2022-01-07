package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/zorrokid/go-retro/database/model"
	"github.com/zorrokid/go-retro/ui/services"
)

type TitleList struct {
	titleService  *services.TitleService
	data          []model.Title
	selectedTitle *model.Title
	list          *widget.List
}

func NewTitleList(titleService *services.TitleService) *TitleList {
	data := titleService.GetTitles()
	list := &TitleList{
		titleService: titleService,
		data:         data,
	}
	return list
}

func (list *TitleList) MakeList() fyne.CanvasObject {
	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List")
	hbox := container.NewHBox(icon, label)
	titleContent := list.makeSelectedTitleContent(1)

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
		label.SetText(list.data[id].Name)
		icon.SetResource(theme.DocumentIcon())
		list.selectedTitle = &list.data[id]
	}
	list.list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
	}

	return container.NewHSplit(list.list, container.NewVSplit(hbox, titleContent))
}

func (list *TitleList) Update() {
	list.data = list.titleService.GetTitles()
	list.list.Refresh()
}

func (list *TitleList) makeSelectedTitleContent(titleId int) fyne.CanvasObject {
	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Title")

	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})
	hbox := container.NewHBox(icon, label, add)

	itemList := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	return container.NewVSplit(hbox, itemList)
}
