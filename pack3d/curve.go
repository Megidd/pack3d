package pack3d

import "github.com/fogleman/fauxgl"

var h *Helix

type Curve interface {
	Coord(int) fauxgl.Vector
	OffsetRandom(int) fauxgl.Vector
	MulStep(float64)
}

// Any 3D model translation is on a curve.
// To limit the packing result to a shape.
func NewCurve() Curve {
	h = NewHelix(50, 40, 10)
	return h
}
