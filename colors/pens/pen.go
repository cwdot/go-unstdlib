package pens

import (
	"github.com/cwdot/go-stdlib/colors"
)

func NewPen(trueC colors.Color, falseC colors.Color) *Pen {
	return &Pen{
		trueC:  trueC,
		falseC: falseC,
	}
}

type Pen struct {
	trueC  colors.Color
	falseC colors.Color
}

func (p *Pen) Ternary(value bool, trueT string, falseT string) string {
	if value {
		return it(p.trueC, trueT)
	}
	return it(p.falseC, falseT)
}

func (p *Pen) Mark(value bool, text string) string {
	if value {
		return it(p.trueC, text)
	}
	return it(p.falseC, text)
}

func it(value colors.Color, text string) string {
	if value == "" {
		return text
	}
	return value.It(text)
}
