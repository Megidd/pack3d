package pack3d

import "github.com/fogleman/fauxgl"

// Local coordinate system origin is moved to bottom-center of bounding box.
func Bottom(m *fauxgl.Mesh) fauxgl.Matrix {
	return m.MoveTo(fauxgl.Vector{}, fauxgl.Vector{X: 0.5, Y: 0.5, Z: 0.0})
}
