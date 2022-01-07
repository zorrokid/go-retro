package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/zorrokid/go-retro/ui/services"
)

type TitleDialog struct {
	titleService *services.TitleService
}

func NewTitleDialog(titleService *services.TitleService) *TitleDialog {
	dialog := &TitleDialog{
		titleService: titleService,
	}
	return dialog
}

func (td *TitleDialog) ShowDialog(win *fyne.Window, update func()) {
	titlename := widget.NewEntry()
	items := []*widget.FormItem{
		widget.NewFormItem("Title name", titlename),
	}

	dialog.ShowForm("Add title", "Add", "Cancel", items, func(b bool) {
		if !b {
			return
		}
		log.Println("Add title", titlename.Text)
		td.titleService.AddTitle(titlename.Text)
		update()
	}, *win)
}
