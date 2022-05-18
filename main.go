package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"flybywind.me/fyne/config"
	"flybywind.me/fyne/view/menu"
	"flybywind.me/fyne/view/pdfview"
)

func main() {
	a := app.New()
	w := a.NewWindow(config.Title)
	w.Resize(config.WindSize)
	w.SetContent(container.NewVBox(menu.NewHomeMenu(), pdfview.NewPdfViewer()))
	w.ShowAndRun()
}
