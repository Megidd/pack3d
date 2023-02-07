package pack3d

import (
	"math/rand"

	"github.com/fogleman/fauxgl"
)

// Restricted mode means:
//
// 1. Assume there is a 3D print floor.
// 2. The Z axis is upward.
// 3. The only transformations allowed are moving along X and Y and rotating around Z axis.
// 4. The bottom of all 3D models are aligned on the 3D print floor.
var Restricted bool = true

// Local coordinate system origin is moved to bottom-center of bounding box.
func Bottom(m *fauxgl.Mesh) fauxgl.Matrix {
	return m.MoveTo(fauxgl.Vector{}, fauxgl.Vector{X: 0.5, Y: 0.5, Z: 0.0})
}

// Random unit vector inside XY plane.
// Its Z component is zero.
// Reference:
// https://github.com/fogleman/fauxgl/blob/27cddc103802008bbda73dc74cab49038d96fdf3/vector.go#L16
func RandomUnitVecXY() fauxgl.Vector {
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
