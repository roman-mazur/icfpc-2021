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

var (
	v00 = Vertex{X: 0, Y: 0}
	v01 = Vertex{X: 0, Y: 1}
	v11 = Vertex{X: 1, Y: 1}
	v10 = Vertex{X: 1, Y: 0}
	v22 = Vertex{X: 2, Y: 2}
	v02 = Vertex{X: 0, Y: 2}
	v20 = Vertex{X: 2, Y: 0}
	v21 = Vertex{X: 2, Y: 1}
	v12 = Vertex{X: 1, Y: 2}
)

func newEdge(A, B Vertex) Edge {
	return Edge{&A, &B}
}

func TestIntersect(t *testing.T) {
	type test struct {
		a, b   Edge
		expect bool
	}

	var suite = []test{
		test{newEdge(v00, v22), newEdge(v02, v20), true},
		test{newEdge(v00, v12), newEdge(v02, v20), true},
		test{newEdge(v10, v10), newEdge(v10, v22), true},
		test{newEdge(v10, v12), newEdge(v00, v20), true},
		//test{newEdge(v01, v21), newEdge(v11, v20), true},
		//test{newEdge(v01, v21), newEdge(v11, v10), true},
		//test{newEdge(v01, v21), newEdge(v11, v12), true},
		test{newEdge(v00, v10), newEdge(v02, v20), false},
		test{newEdge(v00, v20), newEdge(v20, v22), false},
		test{newEdge(v01, v10), newEdge(v11, v21), false},
		test{newEdge(v10, v10), newEdge(v02, v20), false}, // Considers parallel
	}

	for _, test := range suite {
		t.Run(
			fmt.Sprintf("(%d%d-%d%d)x(%d%d-%d%d)_intersect_%v",
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

func TestContains(t *testing.T) {
	type test struct {
		hole   Hole
		vertex Vertex
		expect bool
	}

	hole := Hole{
		Vertices: []Vertex{v10, v11, v12, v01},
	}
	hole.FillEdges()

	var suite = []test{
		//test{hole, v11, true},
		test{hole, v00, false},
	}

	for _, test := range suite {
		t.Run(
			fmt.Sprintf("hole-contains-%d%d-%v", test.vertex.X, test.vertex.Y, test.expect),
			func(t *testing.T) {
				assert.Assert(t, test.hole.Contains(test.vertex) == test.expect)
			},
		)
	}
}
