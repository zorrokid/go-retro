package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/zorrokid/go-retro/ui/services"
)

type GoRetroUi struct {
	app          fyne.App
	titleService *services.TitleService
	list         *TitleList
	titleDialog  *TitleDialog
	window       fyne.Window
}

func NewGoRetroUi() *GoRetroUi {
	titleService := services.NewTitleService()
	app := app.New()
	window := app.NewWindow("Go-Retro!")
	list := NewTitleList(titleService, &window)
	titleDialog := NewTitleDialog(titleService)

	ui := &GoRetroUi{
		app:          app,
		titleService: titleService,
		list:         list,
		titleDialog:  titleDialog,
		window:       window,
	}
	return ui
}

func (ui *GoRetroUi) InitAndRun() {
	ui.window.SetMainMenu(ui.makeMenu(ui.app, ui.window))
	ui.window.SetContent(ui.list.MakeList())
	ui.window.Resize(fyne.NewSize(640, 460))
	ui.window.ShowAndRun()
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
