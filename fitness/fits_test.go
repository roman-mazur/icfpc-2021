package fitness

import (
	"fmt"
	"testing"

	"github.com/roman-mazur/icfpc-2021/data"
	"gotest.tools/assert"
)

func genBigShape() (data.Hole, data.Figure) {
	bigHole := data.Hole{
		Vertices: []data.Vertex{
			data.Vertex{X: 0, Y: 0},
			data.Vertex{X: 1000, Y: 0},
			data.Vertex{X: 1000, Y: 1000},
			data.Vertex{X: 0, Y: 1000},
		},
	}
	bigHole.FillEdges()

	bigFigure := data.Figure{Vertices: bigHole.Vertices, Edges: bigHole.Edges}

	return bigHole, bigFigure
}

func genSmallShape() (data.Hole, data.Figure) {
	smallHole := data.Hole{
		Vertices: []data.Vertex{
			data.Vertex{X: 0, Y: 0},
			data.Vertex{X: 1, Y: 0},
			data.Vertex{X: 1, Y: 1},
			data.Vertex{X: 0, Y: 1},
		},
	}
	smallHole.FillEdges()

	smallFigure := data.Figure{Vertices: smallHole.Vertices, Edges: smallHole.Edges}

	return smallHole, smallFigure
}

func TestFigureFits(t *testing.T) {
	type test struct {
		hole   data.Hole
		figure data.Figure
		expect bool
	}

	line := []data.Vertex{data.Vertex{X: 10, Y: 10}, data.Vertex{X: -10, Y: -10}}
	lineFigure := data.Figure{Vertices: line, Edges: []*data.Edge{&data.Edge{A: &line[0], B: &line[1]}}}

	bigHole, bigFigure := genBigShape()
	smallHole, smallFigure := genSmallShape()

	var suite = []test{
		test{bigHole, smallFigure, true},
		test{smallHole, bigFigure, false},
		test{smallHole, lineFigure, false},
	}

	for i, test := range suite {
		t.Run(
			fmt.Sprintf("figure-fits-%d-%v", i, test.expect),
			func(t *testing.T) {
				assert.Assert(t, Fit(test.figure, test.hole) == test.expect)
			},
		)
	}
}
