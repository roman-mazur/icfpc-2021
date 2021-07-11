package algorithm

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/roman-mazur/icfpc-2021/cmd/solver/search"
	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/fitness"
	"github.com/roman-mazur/icfpc-2021/transform"
)

type action struct {
	Probability int
	Function    func(*data.Figure, *data.Hole, int) string
}

var actionList = []action{
	{100, randomFold},
	{50, randomRotateSmallAngle},
	{30, randomRotate},
	{5, moveToCenter},
	//{20, searchForEdges}, // TODO: Causes too many problems and crashes.
	{1, shortMoveTopLeft},
	{1, shortMoveTop},
	{1, shortMoveTopRight},
	{1, shortMoveLeft},
	{1, shortMoveRight},
	{1, shortMoveBottomLeft},
	{1, shortMoveBottom},
	{1, shortMoveBottomRight},
}

var actionMaxProbability = 0

func init() {
	for _, a := range actionList {
		actionMaxProbability += a.Probability
	}
}

func shortMove(f *data.Figure, vec data.Vertex) string {
	transform.Move(f, vec)

	return fmt.Sprintf("shortMove(%d, %d)", int(vec.X), int(vec.Y))
}

func shortMoveTopLeft(f *data.Figure, h *data.Hole, ε int) string {
	return shortMove(f, data.Vertex{X: -1, Y: 1})
}
func shortMoveTop(f *data.Figure, h *data.Hole, ε int) string {
	return shortMove(f, data.Vertex{X: 0, Y: 1})
}
func shortMoveTopRight(f *data.Figure, h *data.Hole, ε int) string {
	return shortMove(f, data.Vertex{X: 1, Y: 1})
}
func shortMoveLeft(f *data.Figure, h *data.Hole, ε int) string {
	return shortMove(f, data.Vertex{X: -1, Y: 0})
}
func shortMoveRight(f *data.Figure, h *data.Hole, ε int) string {
	return shortMove(f, data.Vertex{X: 1, Y: 0})
}
func shortMoveBottomLeft(f *data.Figure, h *data.Hole, ε int) string {
	return shortMove(f, data.Vertex{X: -1, Y: -1})
}
func shortMoveBottom(f *data.Figure, h *data.Hole, ε int) string {
	return shortMove(f, data.Vertex{X: 0, Y: -1})
}
func shortMoveBottomRight(f *data.Figure, h *data.Hole, ε int) string {
	return shortMove(f, data.Vertex{X: 1, Y: -1})
}

func moveToCenter(f *data.Figure, h *data.Hole, ε int) string {
	target := h.Center()
	v := &f.Vertices[rand.Intn(len(f.Vertices))]
	transform.Matrix(f, v, pixel.IM.Moved(pixel.V(target.X-v.X, target.Y-v.Y)), ε)
	return fmt.Sprintf("moveToCenter(%s)", target)
}

func randomShortMove(f *data.Figure, h *data.Hole, eps int) string {
	var movements = []data.Vertex{
		{X: -1, Y: -1},
		{X: -1, Y: 0},
		{X: -1, Y: 1},
		{X: 0, Y: -1},
		{X: 0, Y: 1},
		{X: 1, Y: -1},
		{X: 1, Y: 0},
		{X: 1, Y: 1},
	}

	return shortMove(f, movements[rand.Intn(len(movements))])
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

func randomRotateSmallAngle(f *data.Figure, h *data.Hole, eps int) string {
	smallAngle := math.Pi / 6.0
	randomAngle := rand.Float64()*2.0*smallAngle - smallAngle

	edge := getRandomEdge(f, h)
	transform.Rotate(f, edge, randomAngle, eps)
	return fmt.Sprintf("randomRotateSmall(%s, %f)", edge, 180*randomAngle/math.Pi)
}

func randomRotate(f *data.Figure, h *data.Hole, eps int) string {
	randomAngle := 2 * math.Pi * rand.Float64()

	edge := getRandomEdge(f, h)
	transform.Rotate(f, edge, randomAngle, eps)
	return fmt.Sprintf("randomRotate(%s, %f)", edge, 180*randomAngle/math.Pi)
}

func getRandomEdge(f *data.Figure, h *data.Hole) *data.Edge {
	unfits := fitness.ListUnfits(*f, *h)
	nbUnfits := len(unfits)

	if len(f.Edges) == 0 {
		fmt.Println(f)
	}
	if true {
		return f.Edges[rand.Intn(len(f.Edges))]
	}
	return unfits[rand.Intn(nbUnfits)].Edge
}

func searchForEdges(f *data.Figure, h *data.Hole, eps int) string {
	var v *data.Vertex
	unfits := fitness.ListUnfits(*f, *h)
	if len(unfits) > 0 {
		vertices := unfits[rand.Intn(len(unfits))].Vertices
		if len(vertices) > 0 {
			v = vertices[rand.Intn(len(vertices))]
			res := search.Solve(f, h, v, eps)
			*f = *res
			if len(f.Edges) == 0 {
				panic("bad search result")
			}
		}
	}
	return fmt.Sprintf("searchForEdges(%s)", v)
}

func randomAlter(f *data.Figure, h *data.Hole, eps int) string {
	prob := rand.Intn(actionMaxProbability)
	for _, a := range actionList {
		prob -= a.Probability
		if prob <= 0 {
			return a.Function(f, h, eps)
		}
	}
	return actionList[0].Function(f, h, eps)
}
