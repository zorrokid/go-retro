package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

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

			imageOpened(reader)
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

			readContent(reader)
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

func readContent(f fyne.URIReadCloser) {

	if f == nil {
		log.Println("Cancelled")
		return
	}
	defer f.Close()

	fmt.Println(f.URI().Name())

	fmt.Println(f.URI().Path())

	// Open a zip archive for reading.
	r, err := zip.OpenReader(f.URI().Path())
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}

}
