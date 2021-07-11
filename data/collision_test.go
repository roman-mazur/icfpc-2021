package data

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

const (
	padding = 0
	scale   = 1
)

// Collection of pre-defined vertices for testing
var (
	v00 = Vertex{X: 0, Y: 0}
	v01 = Vertex{X: 0, Y: 1}
	v02 = Vertex{X: 0, Y: 2}
	v10 = Vertex{X: 1, Y: 0}
	v11 = Vertex{X: 1, Y: 1}
	v12 = Vertex{X: 1, Y: 2}
	v20 = Vertex{X: 2, Y: 0}
	v21 = Vertex{X: 2, Y: 1}
	v22 = Vertex{X: 2, Y: 2}
)

func newEdge(A, B Vertex) Edge {
	return Edge{A: &A, B: &B}
}

func TestIntersect(t *testing.T) {
	type test struct {
		a, b   Edge
		expect bool
	}

	var suite = []test{
		{newEdge(v00, v22), newEdge(v02, v20), true},
		{newEdge(v00, v12), newEdge(v02, v20), true},
		{newEdge(v00, v20), newEdge(v00, v10), false}, // Touch, not intersect
		{newEdge(v10, v12), newEdge(v00, v20), false}, // Touch, not intersect
		{newEdge(v01, v21), newEdge(v11, v20), false}, // Touch, not intersect
		{newEdge(v01, v21), newEdge(v11, v10), false}, // Touch, not intersect
		{newEdge(v01, v21), newEdge(v11, v12), false}, // Touch, not intersect
		{newEdge(v00, v20), newEdge(v20, v22), false}, // Touch, not intersect
		{newEdge(v00, v10), newEdge(v02, v20), false},
		{newEdge(v01, v10), newEdge(v11, v21), false},
		{newEdge(v10, v20), newEdge(v10, v20), false}, // Collinear
		{newEdge(v10, v10), newEdge(v02, v20), false}, // Collinear
		{newEdge(v10, v10), newEdge(v10, v22), false}, // Collinear
		{newEdge(Vertex{X: -1000000, Y: 0}, v00), newEdge(v12, v01), false},
	}

	for _, test := range suite {
		t.Run(
			fmt.Sprintf("(%f%f-%f%f)x(%f%f-%f%f)_intersect_%v",
				test.a.A.X, test.a.A.Y, test.a.B.X, test.a.B.Y,
				test.b.A.X, test.b.A.Y, test.b.B.X, test.b.B.Y,
				test.expect,
			),
			func(t *testing.T) {
				// Scale points to avoid irregular cases with 0 values
				test.a.A.X = (test.a.A.X + padding) * scale
				test.a.A.Y = (test.a.A.Y + padding) * scale
				test.a.B.X = (test.a.B.X + padding) * scale
				test.a.B.Y = (test.a.B.Y + padding) * scale
				test.b.A.X = (test.b.A.X + padding) * scale
				test.b.A.Y = (test.b.A.Y + padding) * scale
				test.b.B.X = (test.b.B.X + padding) * scale
				test.b.B.Y = (test.b.B.Y + padding) * scale

				assert.Assert(t, test.a.Intersect(test.b) == test.expect)
			})
	}
}

func TestContain(t *testing.T) {
	type test struct {
		hole   Hole
		vertex Vertex
		expect bool
	}

	square := Hole{
		Vertices: []Vertex{
			Vertex{X: 10, Y: 0},
			Vertex{X: 20, Y: 10},
			Vertex{X: 10, Y: 20},
			Vertex{X: 0, Y: 10},
		},
	}
	square.FillEdges()

	concave := Hole{
		Vertices: []Vertex{
			Vertex{X: 0, Y: 0},
			Vertex{X: 30, Y: 30},
			Vertex{X: 60, Y: 0},
			Vertex{X: 60, Y: 60},
			Vertex{X: 0, Y: 60},
		},
	}
	concave.FillEdges()

	var suite = []test{
		test{square, Vertex{X: 10, Y: 10}, true},
		test{square, Vertex{X: 5, Y: 5}, true},
		test{square, Vertex{X: 10, Y: 0}, true},
		test{square, v00, false},

		test{concave, Vertex{X: 6, Y: 20}, true},
		test{concave, Vertex{X: 50, Y: 40}, true},
		test{concave, Vertex{X: 30, Y: 30}, true},
		test{concave, Vertex{X: 0, Y: 0}, true},
		test{concave, Vertex{X: 60, Y: 0}, true},
		test{concave, Vertex{X: 30, Y: 10}, false},
		test{concave, Vertex{X: 10, Y: 75}, false},
		test{concave, Vertex{X: 75, Y: 10}, false},
	}

	for _, test := range suite {
		t.Run(
			fmt.Sprintf("hole-contains-(%f,%f)-%v", test.vertex.X, test.vertex.Y, test.expect),
			func(t *testing.T) {
				assert.Assert(t, test.hole.Contain(test.vertex) == test.expect)
			},
		)
	}
}
