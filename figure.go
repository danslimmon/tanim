package tanim

import (
	"github.com/gdamore/tcell/v2"
)

// Figure is an object which can be drawn on the canvas.
type Figure interface {
	DrawCell(Dim) (bool, rune, tcell.Style)
	Extent() Dim
	OnTick(int) bool
	//OnTouch
	Origin() Dim
	SetOrigin(Dim)
}
