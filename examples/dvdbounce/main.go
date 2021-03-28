package main

import (
	"math"
	"time"

	"github.com/danslimmon/tanim"
	"github.com/gdamore/tcell/v2"
)

type Rectangle struct {
	// The X and Y dimensions of the rectangle. It is X cells wide and Y cells tall.
	Sides tanim.Dim
	Style tcell.Style

	origin tanim.Dim
}

// Origin returns the position of the bottom left corner of a rectangle that contains the figure.
func (fig *Rectangle) Origin() tanim.Dim {
	return fig.origin
}

func (fig *Rectangle) SetOrigin(d tanim.Dim) {
	fig.origin = d
}

// Extent returns the dimensions of the rectangular region that fig will be drawn in.
//
// fig.DrawCell will be called for each cell in this region.
func (fig *Rectangle) Extent() tanim.Dim {
	// The sides are lengths, but the Extent box is drawn inclusively. So we need to subtract 1
	// from each dimension.
	return fig.Sides.Sub(tanim.Dim{1, 1})
}

// DrawCell returns what to draw at the given position relative to Origin.
//
// If draw is true, char will be drawn with the returned style at pos. If draw is false, nothing
// will be drawn at pos.
func (fig *Rectangle) DrawCell(pos tanim.Dim) (draw bool, char rune, style tcell.Style) {
	return true, ' ', fig.Style
}

func (fig *Rectangle) OnTick(t int) bool {
	return true
}

type Translator struct {
	Vx, Vy  float64
	Wrapped tanim.Figure

	origin tanim.Dim
	dx, dy float64
}

func (fig *Translator) Origin() tanim.Dim {
	return fig.Wrapped.Origin()
}

func (fig *Translator) SetOrigin(d tanim.Dim) {
	fig.Wrapped.SetOrigin(d)
}

func (fig *Translator) Extent() tanim.Dim {
	return fig.Wrapped.Extent()
}

func (fig *Translator) DrawCell(pos tanim.Dim) (draw bool, char rune, style tcell.Style) {
	return fig.Wrapped.DrawCell(pos)
}

// OnTick is called at every tick. It returns a bool indicating whether fig should continue to
// exist.
func (fig *Translator) OnTick(t int) bool {
	oldOrigin := fig.Origin()
	newOrigin := oldOrigin

	fig.dx += fig.Vx
	fig.dy += fig.Vy

	if math.Abs(fig.dx) >= 1.0 {
		newOrigin.X += int(math.Floor(fig.dx))
		fig.dx = fig.dx - math.Floor(fig.dx)
	}
	if math.Abs(fig.dy) >= 1.0 {
		newOrigin.Y += int(math.Floor(fig.dy))
		fig.dy = fig.dy - math.Floor(fig.dy)
	}

	fig.SetOrigin(newOrigin)
	return true
}

func main() {
	a, err := tanim.NewAnimation([]tanim.Figure{
		&Translator{
			Vx: 0.01,
			Vy: 0.01,
			Wrapped: &Rectangle{
				Sides: tanim.Dim{1, 1},
				Style: tcell.StyleDefault.Background(tcell.ColorYellow),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	a.TickEvery(100 * time.Millisecond)
	a.Start()
	a.Wait()
}
