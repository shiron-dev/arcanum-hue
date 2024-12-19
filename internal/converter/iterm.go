package converter

import (
	"encoding/xml"
	"os"
	"path"

	"github.com/shiron-dev/arcanum-hue/internal/converter/model"
	"github.com/shiron-dev/arcanum-hue/internal/util"
)

const xmlHeader = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
`
const xmlFooter = `</plist>`

func GetItermTheme(cfgPath string, outPath string) error {
	cfg, err := LoadConfigPath(cfgPath)
	if err != nil {
		return err
	}

	outPath = path.Join(outPath, cfg.Name)
	if err := os.MkdirAll(outPath, WriteFilePerm); err != nil {
		return err
	}

	for _, color := range cfg.Colors {
		plist := *model.ItermColorsModelToPlistModel(colorsModelToItermColorsModel(&color.Model))

		xml, err := xml.MarshalIndent(util.StringOrDictMarshaler{Value: plist}, "", "  ")
		if err != nil {
			return err
		}

		plistFile, err := os.Create(path.Join(outPath, cfg.Name+"_"+color.Name+".itermcolors"))
		if err != nil {
			return err
		}
		defer plistFile.Close()

		if _, err := plistFile.WriteString(xmlHeader); err != nil {
			return err
		}

		if _, err := plistFile.Write(xml); err != nil {
			return err
		}

		if _, err := plistFile.WriteString(xmlFooter); err != nil {
			return err
		}
	}

	return nil
}

func colorsModelToItermColorsModel(colors *model.ColorsModel) *model.ItermColorsModel {
	return &model.ItermColorsModel{
		TerminalAnsi: colors.TerminalAnsi,
	}
}
