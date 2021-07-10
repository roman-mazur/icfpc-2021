package transform

import (
	"math"

	"github.com/roman-mazur/icfpc-2021/data"
)

type FoldDirection byte

const (
	FoldRight FoldDirection = iota
	FoldLeft
)

const eps = 0.000001

// Fold transforms the figure mutating its state.
// TODO: Figure out how to copy figures.
func Fold(figure *data.Figure, edge *data.Edge, dir FoldDirection) {
	a, b, c := edge.Line()
	for i, v := range figure.Vertices {
		var diff float64
		if a != 0 {
			diff = (b*float64(v.Y)+c)/(-a) - float64(v.X)
		} else {
			// Horizontal line.
			diff = float64(v.Y) + c // Bottom == Right.
		}
		if math.Abs(diff) < eps {
			// On the line, not touching.
			continue
		}
		actual := FoldLeft
		if diff < 0 {
			actual = FoldRight
		}
		if actual != dir {
			// It has to be flipped.
			// Perpendicular: -b * x + a * y + d = 0.
			d := b*float64(v.X) - a*float64(v.Y)
			// Projected point (intersection with the perpendicular).
			var ix, iy float64
			if b == 0 {
				// Vertical line.
				ix = -c / a
				iy = float64(v.Y) // keep Y.
			} else if a == 0 {
				// Horizontal line.
				ix = float64(v.X) // keep X.
				iy = -c / b
			} else {
				k1, m1 := -a/b, -c/b
				k2, m2 := b/a, -d/a
				ix = (m2 - m1) / (k1 - k2)
				iy = k1*ix + m1
			}

			// Finally flip the point.
			figure.Vertices[i].X = int(math.Round(2*ix - float64(v.X)))
			figure.Vertices[i].Y = int(math.Round(2*iy - float64(v.Y)))
		}
	}
}

func Rotate(edge *data.Edge, Δ float64) {
	x := float64(edge.B.X - edge.A.X) * math.Cos(Δ) - float64(edge.B.Y - edge.A.Y) * math.Sin(Δ) + float64(edge.A.X)
	y := float64(edge.B.X - edge.A.X) * math.Sin(Δ) + float64(edge.B.Y - edge.A.Y) * math.Cos(Δ) + float64(edge.A.Y)
	edge.B.X, edge.B.Y = int(math.Round(x)), int(math.Round(y))
}
