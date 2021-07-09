package transform

import (
	"testing"

	"github.com/roman-mazur/icfpc-2021/data"
	"gotest.tools/assert"
)

func TestFold(t *testing.T) {
	twoSquares := `
{
  "vertices": [[0, 0], [1, 0], [0, 1], [1, 1], [0, 2], [1, 2]],
  "edges": [[0, 1], [1, 3], [3, 2], [2, 0], [2, 4], [4, 5], [5, 3]]
}
`

	t.Run("2 squares - fold to bottom", func(t *testing.T) {
		var f data.Figure
		_ = f.UnmarshalJSON([]byte(twoSquares))
		t.Log(f)
		Fold(f, &data.Edge{A: &f.Vertices[3], B: &f.Vertices[2]}, FoldRight)
		t.Log(f)
		assert.Equal(t, f.Vertices[4].X, 0)
		assert.Equal(t, f.Vertices[4].Y, 0)
		assert.Equal(t, f.Vertices[5].X, 1)
		assert.Equal(t, f.Vertices[5].Y, 0)
	})

	t.Run("2 squares - fold to right", func(t *testing.T) {
		var f data.Figure
		_ = f.UnmarshalJSON([]byte(twoSquares))
		t.Log(f)
		Fold(f, &data.Edge{A: &f.Vertices[1], B: &f.Vertices[3]}, FoldRight)
		t.Log(f)
		assert.Equal(t, f.Vertices[0].X, 2)
		assert.Equal(t, f.Vertices[0].Y, 0)
		assert.Equal(t, f.Vertices[1].X, 1)
		assert.Equal(t, f.Vertices[1].Y, 0)
		assert.Equal(t, f.Vertices[2].X, 2)
		assert.Equal(t, f.Vertices[2].Y, 1)
		assert.Equal(t, f.Vertices[3].X, 1)
		assert.Equal(t, f.Vertices[3].Y, 1)
		assert.Equal(t, f.Vertices[4].X, 2)
		assert.Equal(t, f.Vertices[4].Y, 2)
		assert.Equal(t, f.Vertices[5].X, 1)
		assert.Equal(t, f.Vertices[5].Y, 2)
	})

	square := `
{
  "vertices": [[0, 0], [1, 0], [0, 1], [1, 1]],
  "edges": [[0, 1], [1, 3], [3, 2], [2, 0]]
}
`
	t.Run("2 squares - diagonal fold", func(t *testing.T) {
		var f data.Figure
		_ = f.UnmarshalJSON([]byte(square))
		t.Log(f)
		Fold(f, &data.Edge{A: &f.Vertices[0], B: &f.Vertices[3]}, FoldRight)
		t.Log(f)
		assert.Equal(t, f.Vertices[0].X, 0)
		assert.Equal(t, f.Vertices[0].Y, 0)
		assert.Equal(t, f.Vertices[1].X, 1)
		assert.Equal(t, f.Vertices[1].Y, 0)
		assert.Equal(t, f.Vertices[2].X, 1)
		assert.Equal(t, f.Vertices[2].Y, 0)
		assert.Equal(t, f.Vertices[3].X, 1)
		assert.Equal(t, f.Vertices[3].Y, 1)
	})
}
