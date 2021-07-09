package data

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestEdge_Line(t *testing.T) {
	for _, tcase := range []struct {
		x1, y1, x2, y2 int
		a, b, c        float64
	}{
		{
			x1: 0, y1: 0, x2: 3, y2: 3,
			a: 1, b: -1, c: 0,
		},
		{
			x1: 0, y1: 0, x2: 1, y2: 2,
			a: 2, b: -1, c: 0,
		},
		{
			x1: 0, y1: 0, x2: -1, y2: -2,
			a: 2, b: -1, c: 0,
		},
		{
			x1: 0, y1: 0, x2: -1, y2: 2,
			a: -2, b: -1, c: 0,
		},
		{
			x1: 0, y1: 0, x2: 10, y2: 0,
			a: 0, b: 1, c: 0,
		},
		{
			x1: 5, y1: 5, x2: 10, y2: 5,
			a: 0, b: 1, c: -5,
		},
		{
			x1: 5, y1: 5, x2: 5, y2: 10,
			a: 1, b: 0, c: -5,
		},
	} {
		t.Run(fmt.Sprintf("%f*x + %f*y + %f = 0", tcase.a, tcase.b, tcase.c), func(t *testing.T) {
			a, b, c := (&Edge{A: &Vertex{X: tcase.x1, Y: tcase.y1}, B: &Vertex{X: tcase.x2, Y: tcase.y2}}).Line()
			assert.Equal(t, a, tcase.a)
			assert.Equal(t, b, tcase.b)
			assert.Equal(t, c, tcase.c)
		})
	}
}

func TestFigure_UnmarshalJSON(t *testing.T) {
	twoSquares := `
{
  "vertices": [[0, 0], [1, 0], [0, 1], [1, 1], [0, 2], [1, 2]],
  "edges": [[0, 1], [1, 3], [3, 2], [2, 0], [2, 4], [4, 5], [5, 3]]
}
`
	var f Figure
	e := f.UnmarshalJSON([]byte(twoSquares))
	if e != nil {
		t.Fatal(e)
	}
	assert.Equal(t, len(f.Vertices), 6)
	assert.Equal(t, len(f.Edges), 7)
}
