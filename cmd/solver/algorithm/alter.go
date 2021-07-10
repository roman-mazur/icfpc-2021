package algorithm

import (
	"math"
	"math/rand"

	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/transform"
)

const (
	actionFold = iota
	actionRotate

	actionCount
)

var actionList = map[int]func(*data.Figure){
	actionFold:   randomFold,
	actionRotate: randomRotate,
}

func randomFold(f *data.Figure) {
	direction := transform.FoldRight

	if rand.Intn(1) == 0 {
		direction = transform.FoldLeft
	}

	transform.Fold(f, getRandomEdge(f), direction)
}

func randomRotate(f *data.Figure) {
	randomAngle := math.Pi * rand.Float64()

	transform.Rotate(getRandomEdge(f), randomAngle)
}

func getRandomEdge(f *data.Figure) *data.Edge {
	return f.Edges[rand.Intn(len(f.Edges))]
}

func randomAlter(f *data.Figure) {
	actionList[rand.Intn(actionCount)](f)
}
