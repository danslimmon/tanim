package tanim

import (
	"github.com/gdamore/tcell"
)

// Figure is an object which can be drawn on the canvas.
type Figure interface {
	DrawCell(Dim) (bool, rune, tcell.Style)
	Extent() Dim
	Origin() Dim
	SetOrigin(Dim)
}
