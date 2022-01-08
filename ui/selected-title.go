package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/zorrokid/go-retro/database/model"
	"github.com/zorrokid/go-retro/ui/services"
)

type SelectedTitle struct {
	title        *model.Title
	titleLabel   *widget.Label
	titleService *services.TitleService
	window       *fyne.Window
}

func NewSelectedTitle(title *model.Title, window *fyne.Window, titleService *services.TitleService) *SelectedTitle {
	titleLabel := widget.NewLabel(title.Name)
	selectedTitle := &SelectedTitle{
		title:        title,
		titleLabel:   titleLabel,
		window:       window,
		titleService: titleService,
	}
	return selectedTitle
}

func (selected *SelectedTitle) makeSelectedTitleContent() fyne.CanvasObject {
	icon := widget.NewIcon(nil)

	add := widget.NewButton("Add release", func() {
		selected.openReleaseDialog()
	})
	hbox := container.NewHBox(icon, selected.titleLabel, add)

	itemList := widget.NewList(
		func() int {
			return len(selected.title.Releases)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(selected.title.Releases[id].Edition)
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

func (s *SelectedTitle) openReleaseDialog() {
	releaseDialog := NewReleaseDialog(s.titleService)
	releaseDialog.ShowDialog(s.window, s.title, s.update)
}

func (s *SelectedTitle) update() {
	fmt.Println("updated, print releases")
	for _, r := range s.title.Releases {
		fmt.Println(r.Edition)
	}
}
