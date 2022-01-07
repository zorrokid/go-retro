package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/zorrokid/go-retro/database"
	"github.com/zorrokid/go-retro/ui/services"
)

type GoRetroUi struct {
	app          fyne.App
	db           *database.Database
	titleService *services.TitleService
	list         *TitleList
	titleDialog  *TitleDialog
}

func NewGoRetroUi() *GoRetroUi {
	db := database.NewDatabase()
	db.InitDB()
	titleService := services.NewTitleService(db)
	app := app.New()
	list := NewTitleList(titleService)
	titleDialog := NewTitleDialog(titleService)

	ui := &GoRetroUi{
		app:          app,
		db:           db,
		titleService: titleService,
		list:         list,
		titleDialog:  titleDialog,
	}
	return ui
}

func (ui *GoRetroUi) InitAndRun() {
	mainWindow := ui.app.NewWindow("Go-Retro!")
	mainWindow.SetMainMenu(ui.makeMenu(ui.app, mainWindow))
	mainWindow.SetContent(ui.list.MakeList())
	mainWindow.Resize(fyne.NewSize(640, 460))
	mainWindow.ShowAndRun()
}

func (ui *GoRetroUi) makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", nil)
	newItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Title", func() {
			ui.titleDialog.ShowDialog(&w, ui.list.Update)
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
