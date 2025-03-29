package colors

import (
	"fmt"
	"testing"
)

func TestIt(t *testing.T) {
	colors := map[string]Color{
		"Normal":       Normal,
		"Reset":        Reset,
		"Bold":         Bold,
		"Red":          Red,
		"Green":        Green,
		"Yellow":       Yellow,
		"Blue":         Blue,
		"Magenta":      Magenta,
		"Cyan":         Cyan,
		"BoldRed":      BoldRed,
		"BoldGreen":    BoldGreen,
		"BoldYellow":   BoldYellow,
		"BoldBlue":     BoldBlue,
		"BoldMagenta":  BoldMagenta,
		"BoldCyan":     BoldCyan,
		"FaintRed":     FaintRed,
		"FaintGreen":   FaintGreen,
		"FaintYellow":  FaintYellow,
		"FaintBlue":    FaintBlue,
		"FaintMagenta": FaintMagenta,
		"FaintCyan":    FaintCyan,
		"BgRed":        BgRed,
		"BgGreen":      BgGreen,
		"BgYellow":     BgYellow,
		"BgBlue":       BgBlue,
		"BgMagenta":    BgMagenta,
		"BgCyan":       BgCyan,
		"Faint":        Faint,
		"FaintItalic":  FaintItalic,
		"Reverse":      Reverse,
	}

	for name, c := range colors {
		fmt.Println(name, It(c, "TEST"))
	}
}
