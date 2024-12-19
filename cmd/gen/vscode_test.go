package gen

import (
	"testing"
)

func TestVSCodeCmd(t *testing.T) {
	t.Parallel()

	configFilePath, tempDir := genTestDir(t)

	type args struct {
		args []string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "2 args",
			args: args{args: []string{configFilePath, tempDir}},
		},
		{
			name: "1 arg",
			args: args{args: []string{configFilePath}},
		},
		{
			name: "0 args",
			args: args{args: []string{}},
		},
		{
			name: "reversed args",
			args: args{args: []string{tempDir, configFilePath}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			vscodeCmd.Run(vscodeCmd, tt.args.args)
		})
	}
}
