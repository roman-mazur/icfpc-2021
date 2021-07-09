package data

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestEdge_Line(t *testing.T) {
	for _, tcase := range []struct {
		x1, y1, x2, y2 int
		k, b float64
	}{
		{
			x1: 0, y1: 0, x2: 3, y2: 3,
			k: 1, b: 0,
		},
		{
			x1: 0, y1: 0, x2: 1, y2: 2,
			k: 2, b: 0,
		},
		{
			x1: 0, y1: 0, x2: -1, y2: -2,
			k: 2, b: 0,
		},
		{
			x1: 0, y1: 0, x2: -1, y2: 2,
			k: -2, b: 0,
		},
		{
			x1: 0, y1: 0, x2: 10, y2: 0,
			k: 0, b: 0,
		},
		{
			x1: 5, y1: 5, x2: 10, y2: 5,
			k: 0, b: 5,
		},
	}{
		t.Run(fmt.Sprintf("y = %f*x + %f", tcase.k, tcase.b), func(t *testing.T) {
			k, b := (&Edge{A: &Vertice{X: tcase.x1, Y:tcase.y1}, B: &Vertice{X: tcase.x2, Y: tcase.y2}}).Line()
			assert.Equal(t, k, tcase.k)
			assert.Equal(t, b, tcase.b)
		})
	}
}
