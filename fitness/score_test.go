package fitness

import (
	"testing"

	"github.com/roman-mazur/icfpc-2021/data"
	"gotest.tools/assert"
)

func TestFitScore(t *testing.T) {
	_, bigFigure := genBigShape()
	smallHole, smallFigure := genSmallShape()

	line := []data.Vertex{data.Vertex{10, 10}, data.Vertex{-10, -10}}
	lineFigure := data.Figure{Vertices: line, Edges: []*data.Edge{&data.Edge{A: &line[0], B: &line[1]}}}

	poorScore := FitScore(bigFigure, smallHole)
	niceScore := FitScore(lineFigure, smallHole)
	bestScore := FitScore(smallFigure, smallHole)

	assert.Assert(t, niceScore < poorScore)
	assert.Assert(t, bestScore < niceScore)
}

func TestFitScore_Fits(t *testing.T) {
	bigHole, _ := genBigShape()
	_, smallFigure := genSmallShape()

	score := FitScore(smallFigure, bigHole)

	assert.Assert(t, score < 0)
}
