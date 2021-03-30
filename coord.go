package tanim

import (
	"github.com/gdamore/tcell/v2"
)

func tanimToTcell(screen tcell.Screen, d Dim) Dim {
	_, ymax := screen.Size()
	return Dim{d.X, ymax - d.Y}
}
