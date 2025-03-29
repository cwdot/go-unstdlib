package pens

import (
	"github.com/cwdot/go-stdlib/colors"
)

type LutOpts struct {
	Lut          map[string]colors.Color
	DefaultColor colors.Color
}

func Lut(lut map[string]colors.Color) func(opts *LutOpts) {
	return func(opts *LutOpts) {
		opts.Lut = lut
	}
}

func DefaultColor(defColor colors.Color) func(opts *LutOpts) {
	return func(opts *LutOpts) {
		opts.DefaultColor = defColor
	}
}

func NewLutMarker(opts ...func(opt *LutOpts)) *LutMarker {
	lut := make(map[string]colors.Color)
	var defaultColor colors.Color
	o := &LutOpts{
		Lut:          lut,
		DefaultColor: defaultColor,
	}
	for _, opt := range opts {
		opt(o)
	}
	return &LutMarker{
		lut:          o.Lut,
		defaultColor: o.DefaultColor,
	}
}

type LutMarker struct {
	lut          map[string]colors.Color
	defaultColor colors.Color
}

func (m *LutMarker) Set(name string, c colors.Color) {
	m.lut[name] = c
}

func (m *LutMarker) Mark(value string) string {
	if c, ok := m.lut[value]; ok {
		return c.It(value)
	}
	return m.defaultColor.It(value)
}
