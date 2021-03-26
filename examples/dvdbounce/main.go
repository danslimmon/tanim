package main

import (
	"fmt"

	"github.com/danslimmon/tanim"
	"github.com/gdamore/tcell"
)

type Rectangle struct {
	// type Figure struct {
	//     Origin Dim
	// }
	tanim.Figure

	// The X and Y dimensions of the rectangle. It is X cells wide and Y cells tall.
	X, Y  int
	Style tcell.Style
}

// Extent returns the dimensions of the rectangular region that fig will be drawn in.
//
// fig.DrawCell will be called for each cell in this region.
func (fig Rectangle) Extent() tanim.Dim {
	return tanim.Dim{fig.X, fig.Y}
}

// DrawCell returns what to draw at the given position relative to Origin.
//
// If draw is true, char will be drawn with the returned style at pos. If draw is false, nothing
// will be drawn at pos.
func (fig Rectangle) DrawCell(pos tanim.Dim) (draw bool, char rune, style tcell.Style) {
	return true, ' ', fig.Style
}

type Translator struct {
	tanim.Figure

	Vx, Vy  float64
	Wrapped tanim.Figure

	dx, dy float64
}

func (fig Translator) Extent() tanim.Dim {
	return fig.Wrapped.Extent()
}

func (fig Translator) DrawCell(pos tanim.Dim) (draw bool, char rune, style tcell.Style) {
	return fig.Wrapped.DrawCell(pos)
}

// OnTick is called at every tick. It returns a bool indicating whether fig should continue to
// exist.
func (fig Translator) OnTick(t int) bool {
	fig.dx += Vx
	fig.dy += Vy
	if math.Abs(fig.dx) >= 1.0 {
		fig.Origin.X += math.Floor(fig.dx)
		fig.dx = fig.dx - math.Floor(fig.dx)
	}
	if math.Abs(fig.dy) >= 1.0 {
		fig.Origin.Y += math.Floor(fig.dy)
		fig.dy = fig.dy - math.Floor(fig.dy)
	}
	return true
}

func main() {
	fmt.Println("yo")
}
