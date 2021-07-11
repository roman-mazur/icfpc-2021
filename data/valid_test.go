package data

import (
	"testing"

	"gotest.tools/assert"
)

func TestValid(t *testing.T) {
	var original, transformed Figure
	var epsilon = 15000

	assert.Assert(t, transformed.IsValid(original, epsilon))
}

func TestScaleToPreserveRatio(t *testing.T) {
	for _, tcase := range []struct{
		name string
		n, o *Edge
		eps int
		target, tb, te float64
	} {
		{
			name: "equal",
			n: &Edge{A: &Vertex{1, 1}, B: &Vertex{1, 3}},
			o: &Edge{A: &Vertex{2, 2}, B: &Vertex{2, 4}},
			eps: 1,
		},
		{
			name: "too big",
			n: &Edge{A: &Vertex{2, 2}, B: &Vertex{2, 6}},
			o: &Edge{A: &Vertex{1, 1}, B: &Vertex{1, 3}},
			eps: 2e5,
		},
		{
			name: "too small",
			n: &Edge{A: &Vertex{1, 1}, B: &Vertex{1, 3}},
			o: &Edge{A: &Vertex{2, 2}, B: &Vertex{2, 6}},
			eps: 5e5,
		},
	} {
		t.Run(tcase.name, func(t *testing.T) {
			scale := ScaleToPreserveRatio(tcase.o, tcase.n, tcase.eps)
			newEdge := &Edge{A: tcase.n.A, B: &Vertex{
				X: tcase.n.A.X + (tcase.n.B.X - tcase.n.A.X)*scale,
				Y: tcase.n.A.Y + (tcase.n.B.Y - tcase.n.A.Y)*scale,
			}}
			if !GoodRatio(tcase.o, newEdge, tcase.eps) {
				t.Errorf("Scale %f does not work", scale)
			}
		})
	}
}
