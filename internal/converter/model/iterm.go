package model

import (
	"strconv"
	"strings"
)

type ItermColor struct {
	AlphaComponent float64 `xml:"Alpha Component"`
	BlueComponent  float64 `xml:"Blue Component"`
	ColorSpace     string  `xml:"Color Space"`
	GreenComponent float64 `xml:"Green Component"`
	RedComponent   float64 `xml:"Red Component"`
}

//nolint:govet
type ItermPlistModel struct {
	Ansi0Color           ItermColor `xml:"Ansi 0 Color"`
	Ansi1Color           ItermColor `xml:"Ansi 1 Color"`
	Ansi2Color           ItermColor `xml:"Ansi 2 Color"`
	Ansi3Color           ItermColor `xml:"Ansi 3 Color"`
	Ansi4Color           ItermColor `xml:"Ansi 4 Color"`
	Ansi5Color           ItermColor `xml:"Ansi 5 Color"`
	Ansi6Color           ItermColor `xml:"Ansi 6 Color"`
	Ansi7Color           ItermColor `xml:"Ansi 7 Color"`
	Ansi8Color           ItermColor `xml:"Ansi 8 Color"`
	Ansi9Color           ItermColor `xml:"Ansi 9 Color"`
	Ansi10Color          ItermColor `xml:"Ansi 10 Color"`
	Ansi11Color          ItermColor `xml:"Ansi 11 Color"`
	Ansi12Color          ItermColor `xml:"Ansi 12 Color"`
	Ansi13Color          ItermColor `xml:"Ansi 13 Color"`
	Ansi14Color          ItermColor `xml:"Ansi 14 Color"`
	Ansi15Color          ItermColor `xml:"Ansi 15 Color"`
	BackgroundColor      ItermColor `xml:"Background Color"`
	BadgeColor           ItermColor `xml:"Badge Color"`
	BoldColor            ItermColor `xml:"Bold Color"`
	CursorColor          ItermColor `xml:"Cursor Color"`
	CursorGuideColor     ItermColor `xml:"Cursor Guide Color"`
	CursorTextColor      ItermColor `xml:"Cursor Text Color"`
	ForegroundColor      ItermColor `xml:"Foreground Color"`
	LinkColor            ItermColor `xml:"Link Color"`
	MatchBackgroundColor ItermColor `xml:"Match Background Color"`
	SelectedTextColor    ItermColor `xml:"Selected Text Color"`
	SelectionColor       ItermColor `xml:"Selection Color"`
}

type ItermColorsModel struct {
	TerminalAnsi TerminalAnsiModel
}

func ItermColorsModelToPlistModel(model *ItermColorsModel) *ItermPlistModel {
	return &ItermPlistModel{
		Ansi0Color:  hexToItermColor(model.TerminalAnsi.TerminalAnsiBlack),
		Ansi1Color:  hexToItermColor(model.TerminalAnsi.TerminalAnsiRed),
		Ansi2Color:  hexToItermColor(model.TerminalAnsi.TerminalAnsiGreen),
		Ansi3Color:  hexToItermColor(model.TerminalAnsi.TerminalAnsiYellow),
		Ansi4Color:  hexToItermColor(model.TerminalAnsi.TerminalAnsiBlue),
		Ansi5Color:  hexToItermColor(model.TerminalAnsi.TerminalAnsiMagenta),
		Ansi6Color:  hexToItermColor(model.TerminalAnsi.TerminalAnsiCyan),
		Ansi7Color:  hexToItermColor(model.TerminalAnsi.TerminalAnsiWhite),
		Ansi8Color:  hexToItermColor(model.TerminalAnsi.TerminalAnsiBrightBlack),
		Ansi9Color:  hexToItermColor(model.TerminalAnsi.TerminalAnsiBrightRed),
		Ansi10Color: hexToItermColor(model.TerminalAnsi.TerminalAnsiBrightGreen),
		Ansi11Color: hexToItermColor(model.TerminalAnsi.TerminalAnsiBrightYellow),
		Ansi12Color: hexToItermColor(model.TerminalAnsi.TerminalAnsiBrightBlue),
		Ansi13Color: hexToItermColor(model.TerminalAnsi.TerminalAnsiBrightMagenta),
		Ansi14Color: hexToItermColor(model.TerminalAnsi.TerminalAnsiBrightCyan),
		Ansi15Color: hexToItermColor(model.TerminalAnsi.TerminalAnsiBrightWhite),
	}
}

const hexFF = 255.0

func hexSingleColorToRGBFloat(hex string) (float64, error) {
	num, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return 0, err
	}

	return float64(num) / hexFF, nil
}

func hexToItermColor(hex string) ItermColor {
	hex = strings.TrimPrefix(hex, "#")

	alpha := 1.0
	if len(hex)%3 != 0 {
		alpha, _ = hexSingleColorToRGBFloat(hex[len(hex)-len(hex)/4-1:])
		hex = hex[:len(hex)-len(hex)/4]
	}

	red, err := hexSingleColorToRGBFloat(hex[len(hex)/3:])
	if err != nil {
		red = 0
	}

	hex = hex[:len(hex)-len(hex)/3]

	green, err := hexSingleColorToRGBFloat(hex[len(hex)/2:])
	if err != nil {
		green = 0
	}

	hex = hex[:len(hex)-len(hex)/2]

	blue, err := hexSingleColorToRGBFloat(hex)
	if err != nil {
		blue = 0
	}

	return ItermColor{
		AlphaComponent: alpha,
		BlueComponent:  blue,
		ColorSpace:     "P3",
		GreenComponent: green,
		RedComponent:   red,
	}
}
