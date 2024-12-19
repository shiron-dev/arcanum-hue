package model

import (
	"reflect"
	"testing"
)

func TestItermColorsModelToPlistModel(t *testing.T) {
	t.Parallel()
	t.Skip("skipping test")
}

func TestHexToItermColor(t *testing.T) {
	t.Parallel()

	type args struct {
		hex string
	}

	tests := []struct {
		name string
		args args
		want ItermColor
	}{
		{
			name: "Red",
			args: args{hex: "#FF0000"},
			want: ItermColor{
				AlphaComponent: 1,
				BlueComponent:  0,
				ColorSpace:     "P3",
				GreenComponent: 0,
				RedComponent:   1,
			},
		},
		{
			name: "Cyan",
			args: args{hex: "#00FFFF"},
			want: ItermColor{
				AlphaComponent: 1,
				BlueComponent:  1,
				ColorSpace:     "P3",
				GreenComponent: 1,
				RedComponent:   0,
			},
		},
		{
			name: "aquamarine alpha",
			args: args{hex: "#7FFFD470"},
			want: ItermColor{
				AlphaComponent: 0.4392156862745098,
				BlueComponent:  0.8313725490196079,
				ColorSpace:     "P3",
				GreenComponent: 1,
				RedComponent:   0.4980392156862745,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := hexToItermColor(tt.args.hex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexToItermColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
