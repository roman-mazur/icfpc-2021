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

func flipVertex(edge *data.Edge, v *data.Vertex, dir FoldDirection) bool {
	a, b, c := edge.Line()
	var diff float64
	if a != 0 {
		diff = (b*float64(v.Y)+c)/(-a) - float64(v.X)
	} else {
		// Horizontal line.
		diff = float64(v.Y) + c // Bottom == Right.
	}
	if math.Abs(diff) < eps {
		// On the line, not touching.
		return false
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
		v.X = int(math.Round(2*ix - float64(v.X)))
		v.Y = int(math.Round(2*iy - float64(v.Y)))
		return true
	}
	return false
}

type folder struct {
	visited map[*data.Edge]struct{}
	figure *data.Figure
	foldEdge *data.Edge
	dir FoldDirection
}

func (f *folder) fold(edge *data.Edge) {
	f.visited[edge] = struct{}{}
	for _, candidate := range f.figure.GetConnectedEdges(edge) {
		if _, beenThere := f.visited[candidate]; beenThere {
			continue
		}

		aFlipped := flipVertex(f.foldEdge, candidate.A, f.dir)
		bFlipped := flipVertex(f.foldEdge, candidate.B, f.dir)
		if aFlipped || bFlipped {
			f.fold(candidate)
		}
	}
}

// Fold transforms the figure mutating its state.
func Fold(figure *data.Figure, edge *data.Edge, dir FoldDirection) {
	f := &folder{
		visited:  make(map[*data.Edge]struct{}),
		figure:   figure,
		foldEdge: edge,
		dir:      dir,
	}
	f.fold(edge)
}

// Rotate changes edge.B position rotating it by the delta.
func Rotate(edge *data.Edge, Δ float64) {
	x := float64(edge.B.X-edge.A.X)*math.Cos(Δ) - float64(edge.B.Y-edge.A.Y)*math.Sin(Δ) + float64(edge.A.X)
	y := float64(edge.B.X-edge.A.X)*math.Sin(Δ) + float64(edge.B.Y-edge.A.Y)*math.Cos(Δ) + float64(edge.A.Y)
	edge.B.X, edge.B.Y = int(math.Round(x)), int(math.Round(y))
}
