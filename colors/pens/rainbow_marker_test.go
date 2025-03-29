package pens

import (
	"fmt"
	"testing"

	"github.com/cwdot/go-stdlib/colors"
)

func TestRainbowMarker(t *testing.T) {
	r := NewRainbowMarker()
	tests := []struct {
		name  string
		value string
		color colors.Color
	}{
		{"red", "hello", colors.Red},
		{"green", "hello", colors.Green},
		{"yellow", "hello", colors.Yellow},
		{"blue", "hello", colors.Blue},
		{"magenta", "hello", colors.Magenta},
		{"cyan", "hello", colors.Cyan},
		{"red-loop", "hello", colors.Red},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := r.Mark(tt.value)
			want := fmt.Sprintf("%s%s%s", tt.color, tt.value, colors.Reset)
			if got != want {
				t.Errorf("Mark() = %v, want %v", got, want)
			}
		})
	}
}
