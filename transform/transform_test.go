package transform

import (
	"github.com/roman-mazur/icfpc-2021/data"
	"gotest.tools/assert"
	"testing"
)

func TestFold(t *testing.T) {
	twoSquares := `
{
  "vertices": [[0, 0], [1, 0], [0, 1], [1, 1], [0, 2], [1, 2]],
  "edges": [[0, 1], [1, 3], [3, 2], [2, 0], [2, 4], [4, 5], [5, 3]]
}
`
	var f data.Figure
	_ = f.UnmarshalJSON([]byte(twoSquares))
	t.Log(f)
	Fold(f, &data.Edge{A: f.Vertices[3], B: f.Vertices[2]}, FoldRight)
	t.Log(f)
	assert.Equal(t, f.Vertices[4].X, 0)
	assert.Equal(t, f.Vertices[4].Y, 0)
	assert.Equal(t, f.Vertices[5].X, 1)
	assert.Equal(t, f.Vertices[5].Y, 0)
}
