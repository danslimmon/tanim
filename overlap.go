package tanim

// Contact represents a the cells of one Figure that are in contact with another Figure.
//
// A Contact is returned by extentCache.Overlap.
type Contact struct {
	Figure Figure
	Cells  DimRange
}
