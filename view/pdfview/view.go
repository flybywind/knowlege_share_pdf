package pdfview

import (
	"image/color"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"

	"rsc.io/pdf"

	"flybywind.me/fyne/config"
)

type pdfMeta struct {
	reader   *pdf.Reader
	position float32
	canvas   fyne.CanvasObject
}

var (
	tabs    *container.AppTabs
	curPdf  *pdfMeta
	pdfList []*pdfMeta = make([]*pdfMeta, 0)
)

func NewPdfViewer() fyne.CanvasObject {
	rect := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})
	rect.Resize(fyne.NewSize(config.WindSize.Width/2., config.WindSize.Height/2.))
	tabs = container.NewAppTabs()

	tabs.OnSelected = func(_ *container.TabItem) {
		curPdf = pdfList[tabs.SelectedIndex()]
	}
	return tabs
}

func OpenPDF(fullpath string, parent fyne.Window) error {
	pdf, err := pdf.Open(fullpath)
	if err != nil {
		dialog.NewError(err, parent)
		return err
	}
	pdfList = append(pdfList, &pdfMeta{
		reader:   pdf,
		position: 0,
		canvas:   container.NewWithoutLayout(),
	})
	curPdf = pdfList[len(pdfList)-1]
	_, fileName := path.Split(fullpath)
	tabs.Items = append(tabs.Items, container.NewTabItem(fileName, curPdf.canvas))
	tabs.SelectIndex(len(tabs.Items) - 1)

	return nil
}
