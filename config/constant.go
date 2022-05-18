package config

import (
	"fyne.io/fyne/v2"
)

const (
	Title string = "共慧斋"
)

var WindSize = fyne.Size{
	Width:  1024,
	Height: 1024,
}

type Language uint16

const (
	_ Language = iota
	CN
	EN
)
const (
	MenuOpenTxt    = "Open"
	MenuShareTxt   = "Share"
	MenuCommentTxt = "Comment"
)

var MenuLangMap map[Language]map[string]string = map[Language]map[string]string{
	CN: {
		MenuOpenTxt:    "打开",
		MenuShareTxt:   "分享",
		MenuCommentTxt: "备注",
	},
	EN: {
		MenuOpenTxt:    "Open",
		MenuShareTxt:   "Share",
		MenuCommentTxt: "comment",
	},
}
