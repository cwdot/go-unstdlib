package colors

import (
	"fmt"
)

type Color string

func (c Color) It(text any) string {
	return fmt.Sprintf("%s%v%s", c, text, Reset)
}

func (c Color) Wrap(text any) string {
	if enabled {
		return fmt.Sprintf("%s%v%s", c, text, Reset)
	}
	return fmt.Sprintf("%v", text)
}

func It(c Color, text any) string {
	return fmt.Sprintf("%s%v%s", c, text, Reset)
}

func Wrap(c Color, text any) string {
	if enabled {
		return fmt.Sprintf("%s%v%s", c, text, Reset)
	}
	return fmt.Sprintf("%v", text)
}
