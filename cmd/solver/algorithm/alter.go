package algorithm

import (
	"fmt"
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

var actionList = map[int]func(*data.Figure, int) string{
	actionFold:   randomFold,
	actionRotate: randomRotate,
}

func randomFold(f *data.Figure, eps int) string {
	direction := transform.FoldRight

	if rand.Intn(1) == 0 {
		direction = transform.FoldLeft
	}

	edge := getRandomEdge(f)
	transform.Fold(f, edge, direction)
	return fmt.Sprintf("randomFold(%s)", edge)
}

func randomRotate(f *data.Figure, eps int) string {
	randomAngle := math.Pi * rand.Float64()

	edge := getRandomEdge(f)
	transform.Rotate(f, edge, randomAngle, eps)
	return fmt.Sprintf("randomRotate(%s, %f)", edge, 180 * randomAngle / math.Pi)
}

func getRandomEdge(f *data.Figure) *data.Edge {
	return f.Edges[rand.Intn(len(f.Edges))]
}

func randomAlter(f *data.Figure, eps int) string {
	return actionList[rand.Intn(actionCount)](f, eps)
}
