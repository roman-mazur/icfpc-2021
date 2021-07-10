package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"unsafe"
)

type Vertex struct {
	X, Y int
}

func (v Vertex) String() string {
	return fmt.Sprintf("(%d,%d)", v.X, v.Y)
}

type Edge struct {
	A, B *Vertex
}

func (e Edge) String() string {
	return fmt.Sprintf("[%s->%s]", e.A, e.B)
}

type Hole struct {
	Vertices []Vertex
	Edges    []*Edge
}

type Figure struct {
	Vertices []Vertex
	Edges    []*Edge
}

type Problem struct {
	Hole    *Hole   `json:"hole"`
	Figure  *Figure `json:"figure"`
	Epsilon int     `json:"epsilon"`
}

func (e *Edge) SqLength() float64 {
	return math.Pow(float64(e.A.X-e.B.X), 2) + math.Pow(float64(e.A.Y-e.B.Y), 2)
}

// Line of this edge: a*x + b*y + c = 0.
func (e *Edge) Line() (a, b, c float64) {
	if e.A.Y == e.B.Y {
		if e.A.X == e.B.X {
			panic(fmt.Errorf("it's a point %v", e))
		}
		a = 0
		b = 1
		c = -float64(e.A.Y)
	} else if e.A.X == e.B.X {
		a = 1
		b = 0
		c = -float64(e.A.X)
	} else {
		a = float64(e.A.Y-e.B.Y) / float64(e.A.X-e.B.X)
		b = -1
		c = float64(e.A.Y) - a*float64(e.A.X)
	}
	return
}

func (e Edge) Copy() Edge {
	return NewEdgeFromCopy(*e.A, *e.B)
}

func (v *Vertex) UnmarshalJSON(b []byte) error {
	var rawVtx []int
	if err := json.Unmarshal(b, &rawVtx); err != nil {
		return err
	}

	if len(rawVtx) != 2 {
		return errors.New("invalid vertex")
	}

	v.X = rawVtx[0]
	v.Y = rawVtx[1]

	return nil
}

func (h *Hole) UnmarshalJSON(b []byte) error {
	var vertices []Vertex
	if err := json.Unmarshal(b, &vertices); err != nil {
		return err
	}

	h.Vertices = vertices
	h.FillEdges()

	return nil
}

func (h *Hole) FillEdges() {
	verticesCount := len(h.Vertices)
	h.Edges = make([]*Edge, verticesCount)

	for i := range h.Vertices {
		if i == 0 {
			h.Edges[0] = &Edge{
				A: &h.Vertices[verticesCount-1],
				B: &h.Vertices[0],
			}
			continue
		}
		h.Edges[i] = &Edge{
			A: &h.Vertices[i-1],
			B: &h.Vertices[i],
		}
	}
}

func (f *Figure) UnmarshalJSON(b []byte) error {
	var rawFigure map[string][]Vertex
	if err := json.Unmarshal(b, &rawFigure); err != nil {
		return err
	}

	edges, ok := rawFigure["edges"]
	if !ok {
		return errors.New("invalid edges")
	}

	vertices, ok := rawFigure["vertices"]
	if !ok {
		return errors.New("invalid vertices")
	}

	for _, e := range edges {
		f.Edges = append(f.Edges, &Edge{
			A: &vertices[e.X],
			B: &vertices[e.Y],
		})
	}
	f.Vertices = vertices
	return nil
}

// Copy makes a deep copy of the original Figure. No pointer are overlaping from f to c afterwards.
func (f Figure) Copy() (c Figure) {
	c.Vertices = make([]Vertex, len(f.Vertices))
	c.Edges = make([]*Edge, len(f.Edges))

	for i, v := range f.Vertices {
		c.Vertices[i] = v
	}

	for i, e := range f.Edges {
		// Translate target vertices addresses from f.Vertices array address to c.Vertices array
		// It will point to the same index in the new array
		edge := Edge{
			A: (*Vertex)(unsafe.Pointer(uintptr(unsafe.Pointer(e.A)) - uintptr(unsafe.Pointer(&f.Vertices[0])) + uintptr(unsafe.Pointer(&c.Vertices[0])))),
			B: (*Vertex)(unsafe.Pointer(uintptr(unsafe.Pointer(e.B)) - uintptr(unsafe.Pointer(&f.Vertices[0])) + uintptr(unsafe.Pointer(&c.Vertices[0])))),
		}

		c.Edges[i] = &edge
	}

	return c
}
