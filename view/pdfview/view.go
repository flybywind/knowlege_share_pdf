package pdfview

import (
	"fmt"
	"log"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"rsc.io/pdf"

	"flybywind.me/fyne/config"
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
	log.Println("page number:", pageNum)
	pages := make([]pdf.Page, pageNum)
	for i := 0; i < pageNum; i++ {
		p := pdfReader.Page(i + 1)
		if p.V.IsNull() {
			fyne.LogError("skip page", fmt.Errorf("%d page is null", i+1))
			continue
		}
		pages[i] = p
	}
	pdfList = append(pdfList, &pdfMeta{
		pages:          pages,
		loadingPage:    0,
		pageNum:        pageNum,
		scrollPosition: 0.0,
		content:        container.NewScroll(container.NewWithoutLayout()),
	})

	curPdf = pdfList[len(pdfList)-1]
	curPdf.content.Resize(config.WindSize)
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
	log.Println("text segments:", len(textList))
	content, ok := pdf.content.Content.(*fyne.Container)
	if !ok {
		fyne.LogError("pdf content type error", fmt.Errorf("convert error"))
	}
	for i, text := range textList {
		t := widget.NewLabel(text.S)
		fmt.Println("text:", t, ", rect:", contentPdf.Rect[i])
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
