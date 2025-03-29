package pens

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cwdot/go-stdlib/colors"
)

func TestPen(t *testing.T) {
	r := NewPen(colors.Red, colors.Green)

	t.Run("TernaryTrue", func(t *testing.T) {
		got := r.Ternary(true, "hello", "there")
		wanted := fmt.Sprintf("%s%s%s", colors.Red, "hello", colors.Reset)
		assert.Equal(t, wanted, got)
	})
	t.Run("TernaryFalse", func(t *testing.T) {
		got := r.Ternary(false, "hello", "there")
		wanted := fmt.Sprintf("%s%s%s", colors.Green, "there", colors.Reset)
		assert.Equal(t, wanted, got)
	})

	t.Run("MarkTrue", func(t *testing.T) {
		got := r.Mark(true, "value")
		wanted := fmt.Sprintf("%s%s%s", colors.Red, "value", colors.Reset)
		assert.Equal(t, wanted, got)
	})
	t.Run("MarkFalse", func(t *testing.T) {
		got := r.Mark(false, "value")
		wanted := fmt.Sprintf("%s%s%s", colors.Green, "value", colors.Reset)
		assert.Equal(t, wanted, got)
	})
}
