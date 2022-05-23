package menu

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

	"flybywind.me/fyne/config"
	"flybywind.me/fyne/view/pdfview"
)

func NewHomeMenu(w fyne.Window) fyne.CanvasObject {
	ln := config.GetLanguage()
	openMenu := widget.NewButton(config.MenuLangMap[ln][config.MenuOpenTxt], func() {
		tapOpen(w)
	})
	shareMenu := widget.NewButton(config.MenuLangMap[ln][config.MenuShareTxt], tapShare)
	commentMenu := widget.NewButton(config.MenuLangMap[ln][config.MenuCommentTxt], tapComment)
	return container.NewHBox(openMenu, shareMenu, commentMenu)
}

func tapOpen(win fyne.Window) {
	// todo
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if reader == nil {
			log.Println("Cancelled")
			return
		}
		log.Println("open pdf at path:", reader.URI().Path())
		pdfview.OpenPDF(reader.URI().Path(), win)
	}, win)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))
	fd.Show()
}
func tapShare() {
	// todo

}
func tapComment() {
	// todo

}
