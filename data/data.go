package data

import "math"

type Vertice struct {
	X, Y int
}

type Edge struct {
	A, B *Vertice
}

type Hole struct {
	Vertices []Vertice
}

type Figure struct {
	Vertices []Vertice
	Edges    []Edge
}

func (e Edge) SqLength() float64 {
	return math.Pow(float64(e.A.X-e.B.X), 2) + math.Pow(float64(e.A.Y-e.B.Y), 2)
}
