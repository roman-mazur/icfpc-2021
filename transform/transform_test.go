package transform

import (
	"encoding/json"
	"math"
	"os"
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
		Fold(&f, &data.Edge{A: &f.Vertices[3], B: &f.Vertices[2]}, FoldRight)
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
		Fold(&f, &data.Edge{A: &f.Vertices[1], B: &f.Vertices[3]}, FoldRight)
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

	twoSquaresWithHand := `
{
  "vertices": [[0, 0], [1, 0], [0, 1], [1, 1], [0, 2], [1, 2], [2, 2], [2, 0]],
  "edges": [[0, 1], [1, 3], [3, 2], [2, 0], [2, 4], [4, 5], [5, 3], [5, 6], [6, 7]]
}
`
	t.Run("2 squares with hand", func(t *testing.T) {
		var f data.Figure
		_ = f.UnmarshalJSON([]byte(twoSquaresWithHand))
		t.Log(f)
		Fold(&f, &data.Edge{A: &f.Vertices[1], B: &f.Vertices[3]}, FoldLeft)
		t.Log(f)
		// Hand should not be touched.
		assert.Equal(t, f.Vertices[6].X, 2)
		assert.Equal(t, f.Vertices[6].Y, 2)
		assert.Equal(t, f.Vertices[7].X, 2)
		assert.Equal(t, f.Vertices[7].Y, 0)
	})

	square := `
{
  "vertices": [[0, 0], [1, 0], [0, 1], [1, 1]],
  "edges": [[0, 1], [1, 3], [3, 2], [2, 0]]
}
`
	t.Run("square - diagonal fold", func(t *testing.T) {
		var f data.Figure
		_ = f.UnmarshalJSON([]byte(square))
		t.Log(f)
		Fold(&f, &data.Edge{A: &f.Vertices[0], B: &f.Vertices[3]}, FoldRight)
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

func TestRotate(t *testing.T) {
	t.Run("line", func(t *testing.T) {
		line := `
{
  "vertices": [[1, 1], [5, 1]],
  "edges": [[0, 1]]
}
`
		var f data.Figure
		_ = f.UnmarshalJSON([]byte(line))
		Rotate(f.Edges[0], math.Pi/2)
		assert.Equal(t, f.Vertices[0].X, 1)
		assert.Equal(t, f.Vertices[0].Y, 1)
		assert.Equal(t, f.Vertices[1].X, 1)
		assert.Equal(t, f.Vertices[1].Y, 5)
		Rotate(f.Edges[0], math.Pi/2)
		assert.Equal(t, f.Vertices[0].X, 1)
		assert.Equal(t, f.Vertices[0].Y, 1)
		assert.Equal(t, f.Vertices[1].X, -3)
		assert.Equal(t, f.Vertices[1].Y, 1)
		Rotate(f.Edges[0], -math.Pi/2)
		assert.Equal(t, f.Vertices[0].X, 1)
		assert.Equal(t, f.Vertices[0].Y, 1)
		assert.Equal(t, f.Vertices[1].X, 1)
		assert.Equal(t, f.Vertices[1].Y, 5)
		Rotate(f.Edges[0], math.Pi)
		assert.Equal(t, f.Vertices[0].X, 1)
		assert.Equal(t, f.Vertices[0].Y, 1)
		assert.Equal(t, f.Vertices[1].X, 1)
		assert.Equal(t, f.Vertices[1].Y, -3)
	})
}

func TestFoldAnt(t *testing.T) {
	f, err := os.Open("testdata/ant.problem")
	if err != nil {
		t.Fatal(err)
	}
	var p data.Problem
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		t.Fatal(err)
	}
	original := p.Figure.Copy()
	Fold(p.Figure, p.Figure.Edges[37], FoldRight)
	assert.Equal(t, p.Figure.IsValid(original, p.Epsilon), true)
	Fold(p.Figure, p.Figure.Edges[4], FoldLeft)
	assert.Equal(t, p.Figure.IsValid(original, p.Epsilon), true)
}

func TestFoldKorgi(t *testing.T) {
	f, err := os.Open("testdata/korgi.problem")
	if err != nil {
		t.Fatal(err)
	}
	var p data.Problem
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		t.Fatal(err)
	}
	original := p.Figure.Copy()
	Fold(p.Figure, p.Figure.Edges[27], FoldRight)
	assert.Equal(t, p.Figure.IsValid(original, p.Epsilon), true)
	Fold(p.Figure, p.Figure.Edges[3], FoldLeft)
	assert.Equal(t, p.Figure.IsValid(original, p.Epsilon), true)
	Fold(p.Figure, p.Figure.Edges[30], FoldLeft)
	assert.Equal(t, p.Figure.IsValid(original, p.Epsilon), true)
}
