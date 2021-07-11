package transform

import (
	"github.com/faiface/pixel"
	"github.com/roman-mazur/icfpc-2021/data"
)

const million = 1000000.0

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
	v.X = newPos.X
	v.Y = newPos.Y
	v.Metadata.Reset()
	mt.transformed[v] = struct{}{}

	for _, candidate := range mt.f.GetConnectedVertices(v) {
		oldEdge := &data.Edge{A: &oldV, B: candidate}
		newEdge := &data.Edge{A: v, B: candidate}
		if data.LengthRatio(oldEdge, newEdge) > float64(mt.epsilon)/million {
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
