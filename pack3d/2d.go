package pack3d

import (
	"math/rand"

	"github.com/fogleman/fauxgl"
)

// Random unit vector inside the XY plane.
// Its Z component is zero.
func RanUnitVecXY() fauxgl.Vector {
	for {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		z := float64(0)
		if x*x+y*y+z*z > 1 {
			continue
		}
		return fauxgl.Vector{X: x, Y: y, Z: z}.Normalize()
	}
}
