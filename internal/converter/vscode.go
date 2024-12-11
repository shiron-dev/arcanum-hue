package converter

import (
	"encoding/json"
	"os"
)

const WriteFilePerm = 0o644

func GetVSCodeTheme(cfgPath string, outPath string) error {
	cfg, err := LoadConfigPath(cfgPath)
	if err != nil {
		return err
	}

	vsCodeTheme := colorsModelToVSCodeColorsModel(&cfg.Colors[0].Model)
	vsCodeColorsJSONModel := VSCodeColorsModelToJSONModel(vsCodeTheme)
	vsCodeThemeModel := &VSCodeThemeModel{
		Name:   cfg.Name,
		Type:   cfg.Colors[0].Type,
		Colors: *vsCodeColorsJSONModel,
	}

	jsonData, err := json.MarshalIndent(vsCodeThemeModel, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(outPath, jsonData, WriteFilePerm)
	if err != nil {
		return err
	}

	return nil
}

func colorsModelToVSCodeColorsModel(model *ColorsModel) *VSCodeColorsModel {
	return &VSCodeColorsModel{
		Foreground:                       model.Foreground,
		SecondaryForeground:              model.SecondaryForeground,
		Background:                       model.Background,
		SecondaryBackground:              model.SecondaryBackground,
		BackgroundHighlight:              model.BackgroundHighlight,
		PeekHighlightBackground:          model.PeekHighlightBackground,
		BorderForeground:                 model.BorderForeground,
		BorderBackground:                 model.BorderBackground,
		EditorForeground:                 model.EditorForeground,
		EditorBackground:                 model.EditorBackground,
		DeleteForeground:                 model.DeleteForeground,
		Border2:                          model.Border2,
		TextLinkForeground:               model.TextLinkForeground,
		ActivityBarInactiveForeground:    "#868686",
		BadgeForeground:                  "#F8F8F8",
		ButtonHoverBackground:            "#026EC1",
		ChatSlashCommandForeground:       "#40A6FF",
		ChatEditedFileForeground:         "#E2C08D",
		EditorFindMatchBackground:        "#9E6A03",
		EditorGutterAddedBackground:      "#2EA043",
		EditorLineNumberForeground:       "#6E7681",
		EditorOverviewRulerBorder:        "#010409",
		EditorWidgetBackground:           "#202020",
		InputPlaceholderForeground:       "#989898",
		InputOptionActiveBackground:      "#2489DB82",
		InputOptionActiveBorder:          "#2488DB",
		QuickInputBackground:             "#222222",
		StatusBarItemProminentBackground: "#6E768166",
		TabSelectedBorderTop:             "#6caddf",
		TextPreformatForeground:          "#D0D0D0",
		TextSeparatorForeground:          "#21262D",
	}
}
