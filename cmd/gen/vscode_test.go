package gen_test

import (
	"io"
	"os"
	"path"
	"testing"

	"github.com/shiron-dev/arcanum-hue/cmd/gen"
)

var originalArgs = os.Args

func setArgs(arg string) {
	//nolint:gocritic
	os.Args = append(originalArgs, arg)
}

func resetArgs() {
	os.Args = originalArgs
}

func Test_vscodeCmd(t *testing.T) {
	t.Parallel()

	setArgs("sub")

	defer resetArgs()

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

	gen.VSCodeCmd.Run(gen.VSCodeCmd, []string{configFilePath, tempDir})
}
