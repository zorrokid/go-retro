package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/zorrokid/go-retro/database/model"
)

type SelectedTitle struct {
	title      *model.Title
	titleLabel *widget.Label
}

func NewSelectedTitle(title *model.Title) *SelectedTitle {
	titleLabel := widget.NewLabel(title.Name)
	selectedTitle := &SelectedTitle{
		title:      title,
		titleLabel: titleLabel,
	}
	return selectedTitle
}

func (selected *SelectedTitle) makeSelectedTitleContent() fyne.CanvasObject {
	icon := widget.NewIcon(nil)

	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})
	hbox := container.NewHBox(icon, selected.titleLabel, add)

	itemList := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	return container.NewBorder(hbox, nil, nil, nil, itemList)
}

func (selected *SelectedTitle) Update(title *model.Title) {
	selected.title = title
	selected.titleLabel.Text = title.Name
	selected.titleLabel.Refresh()
}

func (selected *SelectedTitle) Clear() {
	selected.title = nil
	selected.titleLabel.Text = ""
	selected.titleLabel.Refresh()
}
