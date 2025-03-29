package pens

import (
	"github.com/cwdot/go-stdlib/colors"
)

type RainbowOpts struct {
	Colors []colors.Color
}

func Colors(newColors []colors.Color) func(opts *RainbowOpts) {
	return func(opts *RainbowOpts) {
		opts.Colors = newColors
	}
}

func NewRainbowMarker(opts ...func(opt *RainbowOpts)) *RainbowMarker {
	paints := []colors.Color{
		colors.Red,
		colors.Green,
		colors.Yellow,
		colors.Blue,
		colors.Magenta,
		colors.Cyan,
	}
	o := &RainbowOpts{Colors: paints}
	for _, opt := range opts {
		opt(o)
	}
	return &RainbowMarker{
		colors: o.Colors,
	}
}

type RainbowMarker struct {
	colors   []colors.Color
	position int
}

func (m *RainbowMarker) Mark(value string) string {
	if m.position >= len(m.colors) {
		m.position = 0
	}
	c := m.colors[m.position]
	m.position++
	return c.It(value)
}
