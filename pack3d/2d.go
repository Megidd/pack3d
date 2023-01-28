package pack3d

import (
	"math/rand"

	"github.com/fogleman/fauxgl"
)

// Random unit vector inside the XY plane.
// Its Z component is zero.
// Reference:
// https://github.com/fogleman/fauxgl/blob/27cddc103802008bbda73dc74cab49038d96fdf3/vector.go#L16
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
