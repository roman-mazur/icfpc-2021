package data

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
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
			a, b, c := (&Edge{A: &Vertice{X: tcase.x1, Y: tcase.y1}, B: &Vertice{X: tcase.x2, Y: tcase.y2}}).Line()
			assert.Equal(t, a, tcase.a)
			assert.Equal(t, b, tcase.b)
			assert.Equal(t, c, tcase.c)
		})
	}
}
