package data

import (
	"fmt"
	"math"
)

type Vertice struct {
	X, Y int
}

type Edge struct {
	A, B *Vertice
}

type Hole struct {
	Vertices []*Vertice
}

type Figure struct {
	Vertices []*Vertice
	Edges    []*Edge
}

func (e *Edge) SqLength() float64 {
	return math.Pow(float64(e.A.X-e.B.X), 2) + math.Pow(float64(e.A.Y-e.B.Y), 2)
}

// Line of this edge (y=k*x+b)
func (e *Edge) Line() (k, b float64) {
	if e.A.Y == e.B.Y {
		if e.A.X == e.B.X {
			panic(fmt.Errorf("it's a point %v", e))
		}

		k = 0
		b = float64(e.A.Y)
	} else {
		k = float64(e.A.Y - e.B.Y) / float64(e.A.X - e.B.X)
		b = float64(e.A.Y) - k*float64(e.A.X)
	}
	return
}
