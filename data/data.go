package data

import (
	"encoding/json"
	"errors"
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

type Problem struct {
	Hole    *Hole   `json:"hole"`
	Figure  *Figure `json:"figure"`
	Epsilon int     `json:"epsilon"`
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
		k = float64(e.A.Y-e.B.Y) / float64(e.A.X-e.B.X)
		b = float64(e.A.Y) - k*float64(e.A.X)
	}
	return
}

func (v *Vertice) UnmarshalJSON(b []byte) error {
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
	var vertices []*Vertice
	if err := json.Unmarshal(b, &vertices); err != nil {
		return err
	}

	h.Vertices = vertices
	return nil
}

func (f *Figure) UnmarshalJSON(b []byte) error {
	var rawFigure map[string][]*Vertice
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

	for i := 0; i < len(edges); i++ {
		if i%2 == 0 {
			f.Edges = append(f.Edges, &Edge{
				A: edges[i],
			})
		} else {
			f.Edges[i/2].B = edges[i]
		}
	}

	f.Vertices = vertices
	return nil
}
