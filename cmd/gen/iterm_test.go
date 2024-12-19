package gen

import (
	"testing"
)

func TestItermCmd(t *testing.T) {
	t.Parallel()

	configFilePath, tempDir := genTestDir(t)
	itermCmd.Run(itermCmd, []string{configFilePath, tempDir})
}
