package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/zorrokid/go-retro/database/model"
	"github.com/zorrokid/go-retro/ui/services"
)

type ReleaseDialog struct {
	titleService *services.TitleService
}

func NewReleaseDialog(titleService *services.TitleService) *ReleaseDialog {
	dialog := &ReleaseDialog{
		titleService: titleService,
	}
	return dialog
}

func (td *ReleaseDialog) ShowDialog(win *fyne.Window, title *model.Title, update func()) {
	edition := widget.NewEntry()
	items := []*widget.FormItem{
		widget.NewFormItem("Edition", edition),
	}

	dialog.ShowForm("Add release", "Add", "Cancel", items, func(b bool) {
		if !b {
			return
		}
		release := &model.Release{
			Edition: edition.Text,
			TitleId: title.ID,
		}
		title.Releases = append(title.Releases, *release)
		log.Println("Update title", edition.Text)
		td.titleService.Update(title)
		update()
	}, *win)
}
