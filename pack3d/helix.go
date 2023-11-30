package pack3d

import (
	"math"
	"math/rand"

	"github.com/fogleman/fauxgl"
)

type Helix struct {
	Radius float64         //
	Pitch  float64         // Height of one complete helix turn, measured parallel to helix axis.
	Step   float64         // Space between coords on the curve.
	Coords []fauxgl.Vector // Sample points on the curve.
}

func NewHelix(r, p, s float64) *Helix {
	h := &Helix{
		Radius: r,
		Pitch:  p,
		Step:   s,
		Coords: []fauxgl.Vector{},
	}
	return h
}

func (h *Helix) Coord(i int) fauxgl.Vector {
	// pitch == 2πb
	b := h.Pitch / 2 / math.Pi

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

func (h *Helix) OffsetRandom(i int) fauxgl.Vector {
	next := fauxgl.Vector.Sub(h.Coord(i+1), h.Coord(i))
	back := fauxgl.Vector.Sub(h.Coord(i), h.Coord(i-1))
	options := [2]fauxgl.Vector{next, back}
	zeroOrOne := rand.Intn(2)
	return options[zeroOrOne]
}

func (h *Helix) MulStep(float64) {
	h.Step *= 1.2
}
