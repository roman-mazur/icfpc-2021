package data

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
