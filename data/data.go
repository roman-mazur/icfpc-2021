package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"unsafe"

	"github.com/faiface/pixel"
)

type Tristate int

const (
	TristateUnset Tristate = iota
	TristateTrue
	TristateFalse
)

type VertexMetadata struct {
	Score    float64
	IsInHole Tristate
}

func (m *VertexMetadata) Reset() {
	m.Score = 0
	m.IsInHole = TristateUnset
}

type Vertex struct {
	X, Y     float64
	Metadata VertexMetadata
}

func (v Vertex) String() string {
	return fmt.Sprintf("(%d,%d)", int(v.X), int(v.Y))
}

func (v Vertex) MarshalJSON() ([]byte, error) {
	enc := []float64{v.X, v.Y}
	return json.Marshal(enc)
}

type Edge struct {
	Index int // For debugging only.
	A, B  *Vertex
}

func (e Edge) String() string {
	return fmt.Sprintf("%d[%s->%s]", e.Index, e.A, e.B)
}

func (e Edge) OtherV(v *Vertex) *Vertex {
	if e.A == v {
		return e.B
	}
	if e.B == v {
		return e.A
	}
	panic("vertex does not belong to the edge")
}

type Hole struct {
	Vertices []Vertex
	Edges    []*Edge

	center *Vertex
}

func (h *Hole) Center() Vertex {
	if h.center != nil {
		return *h.center
	}
	res := Vertex{}
	for _, v := range h.Vertices {
		res.X += v.X
		res.Y += v.Y
	}
	res.X /= float64(len(h.Vertices))
	res.Y /= float64(len(h.Vertices))
	h.center = &res
	res.Metadata.Reset()
	return res
}

type Figure struct {
	Vertices []Vertex
	Edges    []*Edge
}

func (f *Figure) FindV(v Vertex) *Vertex {
	for i, fv := range f.Vertices {
		if fv.X == v.X && fv.Y == v.Y {
			return &f.Vertices[i]
		}
	}
	return nil
}

func (f *Figure) GetConnectedEdges(e *Edge) []*Edge {
	// TODO: Review later. This may need to be optimized.
	var res []*Edge
	for _, edge := range f.Edges {
		if e == edge {
			continue
		}
		if edge.A == e.A || edge.B == e.A || edge.B == e.B || edge.A == e.B {
			res = append(res, edge)
		}
	}
	return res
}

func (f *Figure) GetConnectedVertices(v *Vertex) []*Vertex {
	// TODO: Review later. This may need to be optimized.
	var res []*Vertex
	for _, edge := range f.Edges {
		if edge.A == v {
			res = append(res, edge.B)
		} else if edge.B == v {
			res = append(res, edge.A)
		}
	}
	return res
}

type Problem struct {
	Hole    *Hole   `json:"hole"`
	Figure  *Figure `json:"figure"`
	Epsilon int     `json:"epsilon"`
}

type Solution struct {
	Vertices []Vertex `json:"vertices"`
}

func (f *Figure) Solution() Solution {
	return Solution{
		Vertices: f.Vertices,
	}
}

func (e Edge) SqLength() float64 {
	x := e.A.X - e.B.X
	y := e.A.Y - e.B.Y
	return x*x + y*y
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
	aCopy := *e.A
	bCopy := *e.B
	return Edge{A: &aCopy, B: &bCopy, Index: e.Index}
}

func (v *Vertex) UnmarshalJSON(b []byte) error {
	var rawVtx []float64
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
				Index: 0,
				A:     &h.Vertices[verticesCount-1],
				B:     &h.Vertices[0],
			}
			continue
		}
		h.Edges[i] = &Edge{
			Index: i,
			A:     &h.Vertices[i-1],
			B:     &h.Vertices[i],
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

	f.Edges = make([]*Edge, len(edges))
	for i, e := range edges {
		f.Edges[i] = &Edge{
			Index: i,
			A:     &vertices[int(e.X)],
			B:     &vertices[int(e.Y)],
		}
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
		c.Edges[i] = &Edge{
			Index: i,
			A:     (*Vertex)(unsafe.Pointer(uintptr(unsafe.Pointer(e.A)) - uintptr(unsafe.Pointer(&f.Vertices[0])) + uintptr(unsafe.Pointer(&c.Vertices[0])))),
			B:     (*Vertex)(unsafe.Pointer(uintptr(unsafe.Pointer(e.B)) - uintptr(unsafe.Pointer(&f.Vertices[0])) + uintptr(unsafe.Pointer(&c.Vertices[0])))),
		}
	}

	return c
}

// pixel library data binding

func (v *Vertex) PVec() pixel.Vec {
	return pixel.V(float64(v.X), float64(v.Y))
}

func (e *Edge) PLine() pixel.Line {
	return pixel.L(e.A.PVec(), e.B.PVec())
}
