package transform

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/roman-mazur/icfpc-2021/data"
)

type FoldDirection byte

const (
	FoldRight FoldDirection = iota
	FoldLeft
)

const eps = 0.000001

const debug = false

func log(format string, args ...interface{}) {
	if debug {
		fmt.Printf(format+"\n", args...)
	}
}

func side(edge *data.Edge, v data.Vertex) (side FoldDirection, onTheLine bool) {
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
		onTheLine = true
	}
	side = FoldLeft
	if diff < 0 {
		side = FoldRight
	}
	return
}

func flipVertex(edge *data.Edge, v *data.Vertex, dir FoldDirection) bool {
	if _, onTheLine := side(edge, *v); onTheLine {
		return false
	}

	log("FLIP %s over %s", v, edge)

	// It has to be flipped.
	// Perpendicular: -b * x + a * y + d = 0.
	a, b, c := edge.Line()
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
	v.X = 2*ix - float64(v.X)
	v.Y = 2*iy - float64(v.Y)
	// Reset its score
	v.Metadata.Reset()
	return true
}

type folder struct {
	visited  map[*data.Edge]struct{}
	flipped  map[*data.Vertex]struct{}
	figure   *data.Figure
	foldEdge *data.Edge
	dir      FoldDirection
}

func (f *folder) visit(edge *data.Edge) {
	f.visited[edge] = struct{}{}
}

func (f *folder) fold(edge *data.Edge) {
	f.visit(edge)
	for _, candidate := range f.figure.GetConnectedEdges(edge) {
		if _, beenThere := f.visited[candidate]; beenThere {
			continue
		}
		log("work with %s", candidate)

		_, aFlipped := f.flipped[candidate.A]
		_, bFlipped := f.flipped[candidate.B]

		if !aFlipped {
			aFlipped = flipVertex(f.foldEdge, candidate.A, f.dir)
			f.flipped[candidate.A] = struct{}{}
		}
		if !bFlipped {
			bFlipped = flipVertex(f.foldEdge, candidate.B, f.dir)
			f.flipped[candidate.B] = struct{}{}
		}
		if aFlipped || bFlipped {
			f.fold(candidate)
		}
	}
}

// Fold transforms the figure mutating its state.
func Fold(figure *data.Figure, edge *data.Edge, dir FoldDirection, excludes ...*data.Edge) {
	f := &folder{
		visited:  make(map[*data.Edge]struct{}),
		flipped:  make(map[*data.Vertex]struct{}),
		figure:   figure,
		foldEdge: edge,
		dir:      dir,
	}
	f.visit(edge)
	for _, candidate := range figure.GetConnectedEdges(edge) {
		sideA, onLineA := side(edge, *candidate.A)
		sideB, onLineB := side(edge, *candidate.B)
		if (onLineA || sideA == dir) && (onLineB || sideB == dir) {
			// Make sure we don't touch first edges that are on the target side.
			log("exclude %s", candidate)
			f.visit(candidate)
		} else {
			log("accept %s", candidate)
		}
	}
	for _, exclude := range excludes {
		log("explicitly exclude %s", exclude)
		f.visit(exclude)
	}
	f.fold(edge)
}

// Rotate changes edge.B position rotating it by the delta.
func Rotate(figure *data.Figure, edge *data.Edge, Δ float64, ε int) {
	Matrix(figure, edge.B, pixel.IM.Rotated(edge.A.PVec(), Δ), ε)
}
