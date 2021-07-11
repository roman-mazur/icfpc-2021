package algorithm

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/fitness"
	"github.com/roman-mazur/icfpc-2021/transform"
)

var actionList = []func(*data.Figure, *data.Hole, int) string{
	// randomFold,
	randomRotate,
}

func randomFold(f *data.Figure, h *data.Hole, eps int) string {
	direction := transform.FoldRight

	if rand.Intn(1) == 0 {
		direction = transform.FoldLeft
	}

	edge := getRandomEdge(f, h)
	transform.Fold(f, edge, direction)
	return fmt.Sprintf("randomFold(%s)", edge)
}

func randomRotate(f *data.Figure, h *data.Hole, eps int) string {
	randomAngle := 2*math.Pi * rand.Float64()

	edge := getRandomEdge(f, h)
	transform.Rotate(f, edge, randomAngle, eps)
	return fmt.Sprintf("randomRotate(%s, %f)", edge, 180*randomAngle/math.Pi)
}

func getRandomEdge(f *data.Figure, h *data.Hole) *data.Edge {
	unfits := fitness.ListUnfits(*f, *h)
	nbUnfits := len(unfits)

	if true {
		return f.Edges[rand.Intn(len(f.Edges))]
	}
	return unfits[rand.Intn(nbUnfits)].Edge
}

func randomAlter(f *data.Figure, h *data.Hole, eps int) string {
	return actionList[rand.Intn(len(actionList))](f, h, eps)
}
