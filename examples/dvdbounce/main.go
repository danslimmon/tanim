package main

import (
	"math"
	"time"

	"github.com/danslimmon/tanim"
	"github.com/gdamore/tcell/v2"
)

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

func (fig *Translator) DrawCell(pos tanim.Dim) (char rune, style tcell.Style) {
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
			Vx: 1,
			Vy: 1,
			Wrapped: &tanim.Box{
				BoxChars: tanim.DefaultBoxChars(),
				Style:    tcell.StyleDefault.Foreground(tcell.ColorYellow).Background(tcell.ColorReset),
				Wrapped: &tanim.Text{
					Text:  "DVD Player",
					Style: tcell.StyleDefault,
				},
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
