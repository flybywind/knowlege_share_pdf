package pdfview

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"

	"flybywind.me/fyne/config"
)

func NewPdfViewer() fyne.CanvasObject {
	rect := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	rect.Resize(fyne.NewSize(config.WindSize.Width/2., config.WindSize.Height/2.))
	tabs := container.NewAppTabs(
		container.NewTabItem("pdf viewer0",
			container.NewWithoutLayout(rect)),
		container.NewTabItem("pdf viewer1",
			container.NewWithoutLayout(rect)))

	tabs.OnSelected = func(_ *container.TabItem) {
		r, g, _, _ := rect.FillColor.RGBA()
		rect.FillColor = color.RGBA{uint8(g), uint8(r), 0, 255}
	}
	return tabs
}
