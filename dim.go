package tanim

type Dim struct {
	X, Y int
}

// Add returns the result of adding extent's dimensions to those of origin.
func (origin Dim) Add(extent Dim) Dim {
	return Dim{origin.X + extent.X, origin.Y + extent.Y}
}

func (d Dim) Each(fn func(Dim)) {
	for x := 0; x < d.X; x++ {
		for y := 0; y < d.Y; y++ {
			fn(Dim{x, y})
		}
	}
}

type dimRange struct {
	Origin Dim
	Extent Dim
}

func (dr dimRange) Each(fn func(Dim)) {
	for x := dr.Origin.X; x <= dr.Extent.X; x++ {
		for y := dr.Origin.Y; y <= dr.Extent.Y; y++ {
			logger.WithField("x", x).WithField("y", y).Info("Each")
			fn(Dim{x, y})
		}
	}
}

func (dr dimRange) IsMember(d Dim) bool {
	if d.X < dr.Origin.X || d.X > dr.Extent.X || d.Y < dr.Origin.Y || d.Y > dr.Extent.Y {
		return false
	}
	return true
}

// Sub returns all points that are members of dr0 but not of dr1.
func (dr0 dimRange) Sub(dr1 dimRange) []Dim {
	rslt := make([]Dim, 0)
	dr0.Each(func(d Dim) {
		if !dr1.IsMember(d) {
			rslt = append(rslt, d)
		}
	})
	return rslt
}
