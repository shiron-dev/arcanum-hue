package gen

import (
	"io"
	"os"
	"path"
	"testing"
)

func TestVSCodeCmd(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()

	configFilePath := path.Join(tempDir, "config.yaml")

	cwDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	testConfig, err := os.Open(path.Join(cwDir, "../../data/config.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	defer testConfig.Close()

	configFile, err := os.Create(configFilePath)
	if err != nil {
		t.Fatal(err)
	}
	defer configFile.Close()

	_, err = io.Copy(configFile, testConfig)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(configFilePath, []byte("name: test"), 0o600)
	if err != nil {
		t.Fatal(err)
	}

	vscodeCmd.Run(vscodeCmd, []string{configFilePath, tempDir})
}
