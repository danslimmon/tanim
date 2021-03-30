package tanim

import (
	"github.com/gdamore/tcell/v2"
)

// Text is a Figure that renders the given text.
type Text struct {
	Text  string
	Style tcell.Style

	origin Dim
}

// Origin returns the position of the bottom left corner of a rectangle that contains the figure.
func (fig *Text) Origin() Dim {
	return fig.origin
}

func (fig *Text) SetOrigin(d Dim) {
	fig.origin = d
}

// Extent returns the dimensions of the rectangular region that fig will be drawn in.
//
// fig.DrawCell will be called for each cell in this region.
func (fig *Text) Extent() Dim {
	return Dim{len(fig.Text) - 1, 0}
}

func (fig *Text) DrawCell(pos Dim) (char rune, style tcell.Style) {
	return rune(fig.Text[pos.X]), fig.Style
}

func (fig *Text) OnTick(t int) bool {
	return true
}
