package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var cfg config

func init() {
	font()

}

func main() {

	// create a fyne app
	a := app.New()
	// create a window for the app
	win := a.NewWindow("Markdown")

	//get the user interface
	edit, preview := cfg.makeUI()
	cfg.createMenuItems(win)

	// set the content of the window
	win.SetContent(container.NewHSplit(edit, preview))

	//show window add run app
	win.Resize(fyne.Size{Width: 800, Height: 500})

	//focus on the ceter window
	win.CenterOnScreen()
	win.ShowAndRun()

}

func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	app.EditWidget = edit
	app.PreviewWidget = preview

	//listener
	edit.OnChanged = preview.ParseMarkdown
	return edit, preview
}

// setting the munuItem
func (app *config) createMenuItems(win fyne.Window) {
	//set the Items
	openMenuItem := fyne.NewMenuItem("Open...", func() {})
	saveMenuItem := fyne.NewMenuItem("Save", func() {})
	saveAsMenuItem := fyne.NewMenuItem("Save as...", app.saveAsFunc(win))
	//show the item on window
	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)
	// Set MainItem contains
	menu := fyne.NewMainMenu(fileMenu)
	win.SetMainMenu(menu)
}

func (app *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if write == nil {
				// user cancelled
				return
			}
			// save file
			write.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = write.URI()

			defer write.Close()

			win.SetTitle(win.Title() + "-" + write.URI().Name())

			app.SaveMenuItem.Disabled = false

		}, win)
		saveDialog.Show()

	}
}
