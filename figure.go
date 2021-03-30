package tanim

import (
	"github.com/gdamore/tcell/v2"
)

// Figure is an object which can be drawn on the canvas.
type Figure interface {
	// DrawCell returns the rune and style to draw  at the given cell.
	//
	// If the value of the returned rune is 0, this indicates that the Figure should not be drawn at
	// the cell. In other words, a rune of 0 means "transparent."
	DrawCell(Dim) (rune, tcell.Style)

	Extent() Dim
	OnTick(int) bool
	//OnTouch
	Origin() Dim
	SetOrigin(Dim)
}
