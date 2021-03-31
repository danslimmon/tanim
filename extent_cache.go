package tanim

type extentCache struct {
	xAxis []Figure
	yAxis []Figure
}

// Contacts returns a slice of Contact
func (ec *extentCache) Contacts(fig Figure) []Contact {
	origin, extent := fig.Origin(), fig.Extent()
	for i := range ec.xAxis[origin.X:origin.X+extent.X] {
		for _, contactFig := range ec.X
	}
}
