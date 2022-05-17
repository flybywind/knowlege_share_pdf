package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"

	"flybywind.me/fyne/config"
)

func main() {
	a := app.New()
	w := a.NewWindow(config.Title)
	w.Resize(config.WindSize)
	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
