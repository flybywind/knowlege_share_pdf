package pdfview

import (
	"fmt"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"rsc.io/pdf"
)

type pdfMeta struct {
	pages          []pdf.Page
	loadingPage    int
	pageNum        int
	scrollPosition float32
	content        *container.Scroll
}

var (
	tabs    *container.AppTabs
	curPdf  *pdfMeta
	pdfList []*pdfMeta = make([]*pdfMeta, 0)
)

func NewPdfViewer() fyne.CanvasObject {
	tabs = container.NewAppTabs()

	tabs.OnSelected = func(_ *container.TabItem) {
		curPdf = pdfList[tabs.SelectedIndex()]
	}
	return tabs
}

func OpenPDF(fullpath string, parent fyne.Window) error {
	pdfReader, err := pdf.Open(fullpath)
	if err != nil {
		dialog.NewError(err, parent)
		return err
	}
	pageNum := pdfReader.NumPage()
	pages := make([]pdf.Page, pageNum)
	for i := 0; i < pageNum; i++ {
		pages[i] = pdfReader.Page(i)
	}
	pdfList = append(pdfList, &pdfMeta{
		pages:          pages,
		loadingPage:    0,
		pageNum:        pageNum,
		scrollPosition: 0.0,
		content:        container.NewScroll(container.NewWithoutLayout()),
	})
	curPdf = pdfList[len(pdfList)-1]
	_, fileName := path.Split(fullpath)
	tabs.Items = append(tabs.Items, container.NewTabItem(fileName, curPdf.content))
	tabs.SelectIndex(len(tabs.Items) - 1)

	curPdf.renderPage()
	return nil
}

func (pdf *pdfMeta) renderPage() {
	// try to read out some content:
	contentPdf := curPdf.pages[curPdf.loadingPage].Content()
	textList := contentPdf.Text
	content, ok := pdf.content.Content.(*fyne.Container)
	if !ok {
		fyne.LogError("pdf content type error", fmt.Errorf("convert error"))
	}
	for _, text := range textList {
		t := widget.NewLabel(text.S)
		t.Move(fyne.NewPos(float32(text.X), float32(text.Y)))
		content.Add(t)
	}
	pdf.nextPage()
}

func (pdf *pdfMeta) nextPage() {
	if pdf.loadingPage < pdf.pageNum-1 {
		pdf.loadingPage += 1
	}
}
