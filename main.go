package main

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"flybywind.me/fyne/config"
	"flybywind.me/fyne/view/menu"
	"flybywind.me/fyne/view/pdfview"
)

func main() {
	a := app.New()
	w := a.NewWindow(config.Title)
	w.SetMainMenu(makeMenu(a, w))
	w.SetMaster()
	w.Resize(config.WindSize)
	w.SetContent(container.NewVBox(menu.NewHomeMenu(w), pdfview.NewPdfViewer()))
	w.ShowAndRun()
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	headMenu := fyne.NewMenu(config.Title,
		fyne.NewMenuItem("language", func() { fmt.Println("language") }),
	)
	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://developer.fyne.io")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItem("Support", func() {
			u, _ := url.Parse("https://fyne.io/support/")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Sponsor", func() {
			u, _ := url.Parse("https://fyne.io/sponsor/")
			_ = a.OpenURL(u)
		}))

	return fyne.NewMainMenu(
		headMenu,
		helpMenu,
	)
}
