package gen

import (
	"testing"
)

func TestGenCmd(t *testing.T) {
	t.Parallel()

	type args struct {
		args []string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{args: []string{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			genCmd.Run(genCmd, tt.args.args)
		})
	}
}
