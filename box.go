package tanim

import (
	"github.com/gdamore/tcell/v2"
)

type BoxChars struct {
	TopLeft     rune
	Top         rune
	TopRight    rune
	Right       rune
	BottomRight rune
	Bottom      rune
	BottomLeft  rune
	Left        rune
}

func DefaultBoxChars() BoxChars {
	return BoxChars{
		TopLeft:     tcell.RuneULCorner,
		Top:         tcell.RuneHLine,
		TopRight:    tcell.RuneURCorner,
		Right:       tcell.RuneVLine,
		BottomRight: tcell.RuneLRCorner,
		Bottom:      tcell.RuneHLine,
		BottomLeft:  tcell.RuneLLCorner,
		Left:        tcell.RuneVLine,
	}
}

// Box draws a box with the given dimensions around another Figure.
type Box struct {
	// The characters to use to draw the box.
	BoxChars BoxChars
	// The X and Y dimensions of the rectangle. It is X cells wide and Y cells tall.
	Style tcell.Style
	// The Figure to draw a box around
	Wrapped Figure

	origin Dim
}

// Origin returns the position of the bottom left corner of a rectangle that contains the figure.
//
// The origin of a Box is one cell down and to the left of Wrapped's origin.
func (fig *Box) Origin() Dim {
	return fig.origin
}

func (fig *Box) SetOrigin(d Dim) {
	fig.origin = d
}

// Extent returns the dimensions of the rectangular region that fig will be drawn in.
//
// We draw in a single-cell border around Wrapped's extent.
func (fig *Box) Extent() Dim {
	wrappedExtent := fig.Wrapped.Extent()
	return wrappedExtent.Add(Dim{2, 2})
}

// pickChar returns the character that should appear at the given position.
//
// It returns a bool indicating whether a box character should be drawn. If this is false, drawing
// should be delegated to Wrapped.
func (fig *Box) pickChar(pos Dim) (bool, rune) {
	extent := fig.Extent()
	if pos.X == 0 && pos.Y == 0 {
		return true, fig.BoxChars.BottomLeft
	} else if pos.X == 0 && pos.Y == extent.Y {
		return true, fig.BoxChars.TopLeft
	} else if pos.X == extent.X && pos.Y == 0 {
		return true, fig.BoxChars.BottomRight
	} else if pos.X == extent.X && pos.Y == extent.Y {
		return true, fig.BoxChars.TopRight
	} else if pos.X == 0 {
		return true, fig.BoxChars.Left
	} else if pos.X == extent.X {
		return true, fig.BoxChars.Right
	} else if pos.Y == 0 {
		return true, fig.BoxChars.Bottom
	} else if pos.Y == extent.Y {
		return true, fig.BoxChars.Top
	}
	return false, ' '
}

// DrawCell returns what to draw at the given position relative to Origin.
func (fig *Box) DrawCell(pos Dim) (rune, tcell.Style) {
	draw, char := fig.pickChar(pos)
	if !draw {
		return fig.Wrapped.DrawCell(pos.Sub(Dim{1, 1}))
	}
	return char, fig.Style
}

func (fig *Box) OnTick(t int) bool {
	return true
}
