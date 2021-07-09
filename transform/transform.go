package transform

import (
	"github.com/roman-mazur/icfpc-2021/data"
	"math"
)

type FoldDirection byte

const (
	FoldRight FoldDirection = iota
	FoldLeft
)

const eps = 0.000001

// Fold transforms the figure mutating its state.
// TODO: Figure out how to copy figures.
func Fold(figure data.Figure, edge *data.Edge, dir FoldDirection) {
	a, b, c := edge.Line()
	for _, v := range figure.Vertices {
		eq := a*float64(v.X) + b*float64(v.Y) + c
		if math.Abs(eq) < eps {
			// On the line, not touching.
			continue
		}
		actual := FoldRight
		if eq > 0 {
			actual = FoldLeft
		}
		if actual != dir {
			// It has to be flipped.
			// Perpendicular: -b * x + a * y + d = 0.
			d := b*float64(v.X) - a*float64(v.Y)
			// Projected point.
			var ix, iy float64
			if b == 0 {
				ix = -c / a
				iy = float64(v.Y) // keep Y.
			} else {
				ix = (a*c - d) / (a*a + b*b)
				iy = (-a*ix - c) / b
			}

			// Finally flip the point.
			v.X = int(2*ix - float64(v.X))
			v.Y = int(2*iy - float64(v.Y))
		}
	}
}

func Rotate(figure data.Figure, edge *data.Edge, angle float64) data.Figure {
	// TODO
	return figure
}
