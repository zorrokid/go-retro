package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/zorrokid/go-retro/database/model"
	"github.com/zorrokid/go-retro/ui/services"
	"github.com/zorrokid/go-retro/ui/util"
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

	screenshotButton := widget.NewButton("Select screenshot (.jpg or .png)", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, *win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}
			defer reader.Close()

			util.ShowImage(reader)
		}, *win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
		fd.Show()
	})

	archiveFileButton := widget.NewButton("Select archive (.zip or .7z)", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, *win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}

			defer reader.Close()

			util.ReadZip(reader.URI().Path())
		}, *win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".zip", ".7z"}))
		fd.Show()
	})

	items := []*widget.FormItem{
		widget.NewFormItem("Edition", edition),
		widget.NewFormItem("Browse screenshots", screenshotButton),
		widget.NewFormItem("Browse archive files", archiveFileButton),
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
