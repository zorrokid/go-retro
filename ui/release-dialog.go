package main

import (
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
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

	fileButton := widget.NewButton("File Open With Filter (.jpg or .png)", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, *win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}

			imageOpened(reader)
		}, *win)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
		fd.Show()
	})

	items := []*widget.FormItem{
		widget.NewFormItem("Edition", edition),
		widget.NewFormItem("Browse", fileButton),
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

func imageOpened(f fyne.URIReadCloser) {
	if f == nil {
		log.Println("Cancelled")
		return
	}
	defer f.Close()

	showImage(f)
}

func loadImage(f fyne.URIReadCloser) *canvas.Image {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fyne.LogError("Failed to load image data", err)
		return nil
	}
	res := fyne.NewStaticResource(f.URI().Name(), data)

	return canvas.NewImageFromResource(res)
}

func showImage(f fyne.URIReadCloser) {
	img := loadImage(f)
	if img == nil {
		return
	}
	img.FillMode = canvas.ImageFillOriginal

	w := fyne.CurrentApp().NewWindow(f.URI().Name())
	w.SetContent(container.NewScroll(img))
	w.Resize(fyne.NewSize(320, 240))
	w.Show()
}
