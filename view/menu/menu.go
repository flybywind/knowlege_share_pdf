package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"flybywind.me/fyne/config"
)

func NewHomeMenu() fyne.CanvasObject {
	ln := config.GetLanguage()
	openMenu := widget.NewButton(config.MenuLangMap[ln][config.MenuOpenTxt], tapOpen)
	shareMenu := widget.NewButton(config.MenuLangMap[ln][config.MenuShareTxt], tapShare)
	commentMenu := widget.NewButton(config.MenuLangMap[ln][config.MenuCommentTxt], tapComment)
	return container.NewHBox(openMenu, shareMenu, commentMenu)
}

func tapOpen() {
	// todo

}
func tapShare() {
	// todo

}
func tapComment() {
	// todo

}
