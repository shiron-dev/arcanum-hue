package converter

import (
	"encoding/json"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/shiron-dev/arcanum-hue/templates/vscode"
)

const WriteFilePerm = 0o755

func GetVSCodeTheme(cfgPath string, outPath string) error {
	cfg, err := LoadConfigPath(cfgPath)
	if err != nil {
		return err
	}

	themes := []VSCodeThemeModel{}

	for _, color := range cfg.Colors {
		themes = append(themes, VSCodeThemeModel{
			Name:    color.Name,
			UITheme: color.VSCodeUITheme,
			Type:    color.Type,
			Colors:  *VSCodeColorsModelToJSONModel(colorsModelToVSCodeColorsModel(&color.Model)),
		})
	}

	vsCodeExt := &VSCodeThemeExtension{
		Name:        cfg.Name,
		Description: cfg.Description,
		Version:     "0.0.1",
		Themes:      themes,
	}

	err = makeVSCodeThemeExtension(path.Join(outPath, toKebabCase(cfg.Name)), *vsCodeExt)
	if err != nil {
		return err
	}

	return nil
}

func makeVSCodeThemeExtension(outPath string, ext VSCodeThemeExtension) error {
	if err := os.MkdirAll(path.Join(outPath, "/themes"), WriteFilePerm); err != nil {
		return err
	}

	if err := os.MkdirAll(path.Join(outPath, "/.vscode"), WriteFilePerm); err != nil {
		return err
	}

	for _, theme := range ext.Themes {
		if themeJson, err := json.MarshalIndent(theme, "", "  "); err != nil {
			return err
		} else {
			if err := os.WriteFile(
				path.Join(outPath, "/themes/"+toKebabCase(theme.Name)+"-color-theme.json"),
				themeJson, WriteFilePerm); err != nil {
				return err
			}
		}
	}

	if err := templateWithInterface(path.Join(outPath, "/package.json"), vscode.PackageJSON, ext); err != nil {
		return err
	}

	if err := templateWithInterface(path.Join(outPath, "/README.md"), vscode.README, ext); err != nil {
		return err
	}

	if err := templateWithInterface(path.Join(outPath, "/.vscodeignore"), vscode.VSCodeIgnore, nil); err != nil {
		return err
	}

	if err := templateWithInterface(path.Join(outPath, "/CHANGELOG.md"), vscode.CHANGELOG, nil); err != nil {
		return err
	}

	if err := templateWithInterface(path.Join(outPath, "/.vscode/launch.json"), vscode.LaunchJSON, nil); err != nil {
		return err
	}

	return nil
}

func toKebabCase(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}

func templateWithInterface(outPath string, parse string, obj interface{}) error {
	funcMap := template.FuncMap{
		"kebabcase": toKebabCase,
		"sub":       func(a, b int) int { return a - b },
	}

	tmpl, err := template.New("tmpl").Funcs(funcMap).Parse(parse)
	if err != nil {
		return err
	}

	if readme, err := os.Create(outPath); err != nil {
		return err
	} else {
		if err := tmpl.Execute(readme, obj); err != nil {
			return err
		}
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
