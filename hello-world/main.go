package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

type App struct {
	output *widget.Label
}

var myApp App

func main() {
	a := app.New()
	w := a.NewWindow("Hello, World!")

	//	w.SetContent(widget.NewLabel("Hello, World!"))

	output, entry, btn := myApp.makeUI()
	w.SetContent(container.NewVBox(output, entry, btn))
	// Resize the windows
	w.Resize(fyne.Size{Width: 500, Height: 500})

	w.ShowAndRun()

	//same as following
	//w.Show()
	//a.Run()
	//when the windows off run tidy_function
	tidy()

}

func tidy() {
	fmt.Println("Would tidy up")
}

func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello, World!")
	entry := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		app.output.SetText(entry.Text)
	})
	//set the btn color
	btn.Importance = widget.HighImportance
	app.output = output
	return output, entry, btn
}
