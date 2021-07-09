package transform

import (
	"fmt"
	"github.com/roman-mazur/icfpc-2021/data"
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
	f.UnmarshalJSON([]byte(twoSquares))
	fmt.Println(f)
}
