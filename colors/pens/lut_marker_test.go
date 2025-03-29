package pens

import (
	"fmt"
	"testing"

	"github.com/cwdot/go-stdlib/colors"
)

func TestLutMarker(t *testing.T) {
	r := NewLutMarker()
	r.Set("hello", colors.Red)
	r.Set("there", colors.Green)
	tests := []struct {
		name  string
		value string
		color colors.Color
	}{
		{"red", "hello", colors.Red},
		{"green", "there", colors.Green},
		{"default", "billy", colors.Normal},
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
