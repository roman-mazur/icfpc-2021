package fitness

import (
	"fmt"
	"testing"

	"github.com/roman-mazur/icfpc-2021/data"
	"gotest.tools/assert"
)

func TestFigureFits(t *testing.T) {
	type test struct {
		hole   data.Hole
		figure data.Figure
		expect bool
	}

	bigHole := data.Hole{
		Vertices: []data.Vertex{
			data.Vertex{0, 0},
			data.Vertex{1000, 0},
			data.Vertex{1000, 1000},
			data.Vertex{0, 1000},
		},
	}
	bigHole.FillEdges()

	smallHole := data.Hole{
		Vertices: []data.Vertex{
			data.Vertex{0, 0},
			data.Vertex{1, 0},
			data.Vertex{1, 1},
			data.Vertex{0, 1},
		},
	}
	smallHole.FillEdges()

	line := []data.Vertex{data.Vertex{10, 10}, data.Vertex{-10, -10}}
	lineFigure := data.Figure{Vertices: line, Edges: []*data.Edge{&data.Edge{A: &line[0], B: &line[1]}}}

	smallFigure := data.Figure{Vertices: smallHole.Vertices, Edges: smallHole.Edges}
	bigFigure := data.Figure{Vertices: bigHole.Vertices, Edges: bigHole.Edges}

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
