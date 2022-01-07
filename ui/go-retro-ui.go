package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/zorrokid/go-retro/database"
	"github.com/zorrokid/go-retro/ui/services"
)

type GoRetroUi struct {
	app          fyne.App
	db           *database.Database
	titleService *services.TitleService
	list         *TitleList
}

func NewGoRetroUi() *GoRetroUi {
	db := database.NewDatabase()
	db.InitDB()
	titleService := services.NewTitleService(db)
	app := app.New()
	list := NewTitleList(titleService)

	ui := &GoRetroUi{
		app:          app,
		db:           db,
		titleService: titleService,
		list:         list,
	}
	return ui
}

func (ui *GoRetroUi) InitAndRun() {
	mainWindow := ui.app.NewWindow("Go-Retro!")
	mainWindow.SetMainMenu(ui.makeMenu(ui.app, mainWindow))
	mainWindow.SetContent(ui.list.MakeList())
	mainWindow.ShowAndRun()
}

func (ui *GoRetroUi) makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", nil)
	newItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Title", func() {
			ui.openAddTitleDialog(w)
		}),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Info", func() {
			fmt.Println("Add info")
		}),
	)

	file := fyne.NewMenu("File", newItem)
	return fyne.NewMainMenu(
		file,
		helpMenu,
	)
}

func (ui *GoRetroUi) openAddTitleDialog(win fyne.Window) {
	titlename := widget.NewEntry()
	items := []*widget.FormItem{
		widget.NewFormItem("Title name", titlename),
	}

	dialog.ShowForm("Add title", "Add", "Cancel", items, func(b bool) {
		if !b {
			return
		}
		log.Println("Add title", titlename.Text)
		ui.titleService.AddTitle(titlename.Text)
		ui.list.Update()
	}, win)
}
