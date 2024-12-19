package converter

import (
	"encoding/json"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/shiron-dev/arcanum-hue/internal/converter/model"
	"github.com/shiron-dev/arcanum-hue/templates/vscode"
)

const WriteFilePerm = 0o755

func GetVSCodeTheme(cfgPath string, outPath string) error {
	cfg, err := LoadConfigPath(cfgPath)
	if err != nil {
		return err
	}

	themes := []model.VSCodeThemeModel{}

	for _, color := range []model.ColorTheme{cfg.Light, cfg.Dark} {
		themes = append(themes, model.VSCodeThemeModel{
			Name:    color.Name,
			UITheme: color.VSCodeUITheme,
			Colors:  *model.VSCodeColorsModelToJSONModel(colorsModelToVSCodeColorsModel(&color.Model)),
		})
	}

	vsCodeExt := &model.VSCodeThemeExtension{
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

//nolint:cyclop
func makeVSCodeThemeExtension(outPath string, ext model.VSCodeThemeExtension) error {
	if err := os.MkdirAll(path.Join(outPath, "/themes"), WriteFilePerm); err != nil {
		return err
	}

	if err := os.MkdirAll(path.Join(outPath, "/.vscode"), WriteFilePerm); err != nil {
		return err
	}

	for _, theme := range ext.Themes {
		themeJSON, err := json.MarshalIndent(theme, "", "  ")
		if err != nil {
			return err
		}

		if err := os.WriteFile(
			path.Join(outPath, "/themes/"+toKebabCase(theme.Name)+"-color-theme.json"),
			themeJSON, WriteFilePerm); err != nil {
			return err
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

	readme, err := os.Create(outPath)
	if err != nil {
		return err
	}

	return tmpl.Execute(readme, obj)
}

func colorsModelToVSCodeColorsModel(colors *model.ColorsModel) *model.VSCodeColorsModel {
	return &model.VSCodeColorsModel{
		Foreground:          colors.Foreground,
		HiForeground:        colors.HiForeground,
		SecondaryForeground: colors.SecondaryForeground,
		ForegroundAccent:    colors.ForegroundAccent,
		WarningForeground:   colors.WarningForeground,

		Background:          colors.Background,
		EditorBackground:    colors.EditorBackground,
		SecondaryBackground: colors.SecondaryBackground,
		BackgroundAccent:    colors.BackgroundAccent,
		HiBackgroundAccent:  colors.HiBackgroundAccent,
		MatchBackground:     colors.MatchBackground,

		Border: colors.Border,

		TerminalAnsi: colors.TerminalAnsi,
	}
}
