package converter

import (
	"os"
	"path"
)

func GetAllTheme(cfgPath string, outPath string) error {
	cfg, err := LoadConfigPath(cfgPath)
	if err != nil {
		return err
	}

	outPath = path.Join(outPath, toKebabCase(cfg.Name))
	if err := os.MkdirAll(outPath, WriteFilePerm); err != nil {
		return err
	}

	err = GetItermTheme(cfgPath, outPath)
	if err != nil {
		return err
	}

	err = GetVSCodeTheme(cfgPath, outPath)
	if err != nil {
		return err
	}

	return nil
}
