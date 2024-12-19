package gen

import (
	"testing"
)

func TestVSCodeCmd(t *testing.T) {
	t.Parallel()

	configFilePath, tempDir := genTestDir(t)
	vscodeCmd.Run(vscodeCmd, []string{configFilePath, tempDir})
}
