package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/zorrokid/go-retro/database"
	"github.com/zorrokid/go-retro/database/model"
	"github.com/zorrokid/go-retro/ui/services"
)

type GoRetroUi struct {
	app          fyne.App
	db           *database.Database
	titleService *services.TitleService
	data         []model.Title
	list         *widget.List
}

func NewGoRetroUi() *GoRetroUi {
	db := database.NewDatabase()
	db.InitDB()
	titleService := services.NewTitleService(db)
	data := titleService.GetTitles()
	app := app.New()

	ui := &GoRetroUi{
		app:          app,
		db:           db,
		titleService: titleService,
		data:         data,
	}
	return ui
}

func (ui *GoRetroUi) InitAndRun() {
	mainWindow := ui.app.NewWindow("Go-Retro!")
	mainWindow.SetMainMenu(ui.makeMenu(ui.app, mainWindow))
	mainWindow.SetContent(ui.makeListTab(mainWindow))
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
		ui.refreshData()
	}, win)
}

func (ui *GoRetroUi) refreshData() {
	ui.data = ui.titleService.GetTitles()
	ui.list.Refresh()
}

func (ui *GoRetroUi) makeListTab(_ fyne.Window) fyne.CanvasObject {

	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Select An Item From The List")
	hbox := container.NewHBox(icon, label)
	titleContent := ui.makeSelectedTitleContent(1)

	ui.list = widget.NewList(
		func() int {
			return len(ui.data)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(ui.data[id].Name)
		},
	)
	ui.list.OnSelected = func(id widget.ListItemID) {
		label.SetText(ui.data[id].Name)
		icon.SetResource(theme.DocumentIcon())
	}
	ui.list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Select An Item From The List")
		icon.SetResource(nil)
	}

	return container.NewHSplit(ui.list, container.NewVSplit(hbox, titleContent))
}

func (ui *GoRetroUi) makeSelectedTitleContent(titleId int) fyne.CanvasObject {
	icon := widget.NewIcon(nil)
	label := widget.NewLabel("Title")

	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	add := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})
	hbox := container.NewHBox(icon, label, add)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	return container.NewVSplit(hbox, list)
}
