package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/zorrokid/go-retro/database"
	"github.com/zorrokid/go-retro/database/model"
	"github.com/zorrokid/go-retro/ui/services"
)

var db *database.Database
var titleService *services.TitleService
var data []model.Title
var list *widget.List

func main() {
	db = database.NewDatabase()
	db.InitDB()
	titleService = services.NewTitleService(db)
	data = titleService.GetTitles()

	app := app.New()
	mainWindow := app.NewWindow("Go-Retro!")
	mainWindow.SetMainMenu(makeMenu(app, mainWindow))
	mainWindow.SetContent(makeListTab(mainWindow))
	mainWindow.ShowAndRun()
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("New", nil)
	newItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Title", func() {
			openAddTitleDialog(w)
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

func openAddTitleDialog(win fyne.Window) {
	titlename := widget.NewEntry()
	items := []*widget.FormItem{
		widget.NewFormItem("Title name", titlename),
	}

	dialog.ShowForm("Add title", "Add", "Cancel", items, func(b bool) {
		if !b {
			return
		}
		log.Println("Add title", titlename.Text)
		titleService.AddTitle(titlename.Text)
		refreshData()
	}, win)
}

func refreshData() {
	data = titleService.GetTitles()
	list.Refresh()
}

func makeListTab(_ fyne.Window) fyne.CanvasObject {

	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List")
	hbox := container.NewHBox(icon, label)

	list = widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(data[id].Name)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		label.SetText(data[id].Name)
		icon.SetResource(theme.DocumentIcon())
	}
	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
	}

	list.Select(125)

	return container.NewHSplit(list, container.NewCenter(hbox))
}
