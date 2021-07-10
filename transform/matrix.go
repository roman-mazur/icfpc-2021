package transform

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/roman-mazur/icfpc-2021/data"
)

type matrixTransformer struct {
	transformed map[*data.Vertex]struct{}
	m           pixel.Matrix
	f           *data.Figure
	epsilon     int
}

func (mt *matrixTransformer) apply(v *data.Vertex) {
	if _, done := mt.transformed[v]; done {
		return
	}

	oldV := *v

	newPos := mt.m.Project(v.PVec())
	v.X = int(math.Round(newPos.X))
	v.Y = int(math.Round(newPos.Y))
	mt.transformed[v] = struct{}{}

	for _, candidate := range mt.f.GetConnectedVertices(v) {
		oldEdge := &data.Edge{A: &oldV, B: candidate}
		newEdge := &data.Edge{A: v, B: candidate}
		if !data.GoodRatio(oldEdge, newEdge, mt.epsilon) {
			mt.apply(candidate)
		}
	}
}

func Matrix(figure *data.Figure, v *data.Vertex, m pixel.Matrix, epsilon int) {
	(&matrixTransformer{
		transformed: make(map[*data.Vertex]struct{}),
		m:           m,
		epsilon:     epsilon,
		f:           figure,
	}).apply(v)
}
