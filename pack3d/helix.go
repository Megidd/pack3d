package pack3d

import (
	"math"
	"math/rand"

	"github.com/fogleman/fauxgl"
)

type Helix struct {
	Radius float64 //
	Pitch  float64 // Height of one complete helix turn, measured parallel to helix axis.
	Step   float64 // Space between coords on the curve.
}

func NewHelix(r, p, s float64) *Helix {
	h := &Helix{
		Radius: r,
		Pitch:  p,
		Step:   s,
	}
	return h
}

func (h *Helix) Coord(i int) fauxgl.Vector {
	// pitch == 2πb
	b := h.Pitch / 2.0 / math.Pi

	t := float64(i) * h.Step

	x := h.Radius * math.Cos(t)
	y := h.Radius * math.Sin(t)
	z := b * t

	return fauxgl.Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

func (h *Helix) Tangent(i int) fauxgl.Vector {
	b := h.Pitch / (2.0 * math.Pi)
	t := float64(i) * h.Step

	// The derivatives of x, y, and z with respect to t
	dx := -h.Radius * math.Sin(t)
	dy := h.Radius * math.Cos(t)
	dz := b

	return fauxgl.Vector{
		X: dx,
		Y: dy,
		Z: dz,
	}
}

func (h *Helix) TangentRandom(i int) fauxgl.Vector {
	tangentPos := h.Tangent(i).Normalize()
	tangentNeg := tangentPos.MulScalar(-1)

	zeroOrOne := rand.Intn(2)

	if zeroOrOne == 0 {
		return tangentPos
	} else if zeroOrOne == 1 {
		return tangentNeg
	} else {
		return fauxgl.Vector{}
	}
}

func (h *Helix) MulStep(float64) {
	h.Step *= 1.2
}
