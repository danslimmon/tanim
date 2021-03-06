package tanim

type Dim struct {
	X, Y int
}

// Add returns the result of adding d1's dimensions to those of d0.
func (d0 Dim) Add(d1 Dim) Dim {
	return Dim{d0.X + d1.X, d0.Y + d1.Y}
}

// Sub returns the result of subtracting d1's dimensions from those of d0.
func (d0 Dim) Sub(d1 Dim) Dim {
	negd1 := Dim{-d1.X, -d1.Y}
	return d0.Add(negd1)
}

type DimRange struct {
	Origin Dim
	Extent Dim
}

func (dr DimRange) Each(fn func(Dim)) {
	for x := dr.Origin.X; x <= dr.Extent.X; x++ {
		for y := dr.Origin.Y; y <= dr.Extent.Y; y++ {
			logger.WithField("x", x).WithField("y", y).Info("Each")
			fn(Dim{x, y})
		}
	}
}

func (dr DimRange) IsMember(d Dim) bool {
	if d.X < dr.Origin.X || d.X > dr.Extent.X || d.Y < dr.Origin.Y || d.Y > dr.Extent.Y {
		return false
	}
	return true
}

// Sub returns all points that are members of dr0 but not of dr1.
func (dr0 DimRange) Sub(dr1 DimRange) []Dim {
	rslt := make([]Dim, 0)
	dr0.Each(func(d Dim) {
		if !dr1.IsMember(d) {
			rslt = append(rslt, d)
		}
	})
	return rslt
}
