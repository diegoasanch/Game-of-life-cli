package renderer

import (
	tm "github.com/buger/goterm"
)

func GetConsoleDimensions() (width, height int) {
	width = tm.Width()
	height = tm.Height()
	return
}
