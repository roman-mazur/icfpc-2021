package transform

import (
	"github.com/faiface/pixel"
	"github.com/roman-mazur/icfpc-2021/data"
)

const million = 1000000.0

type matrixTransformer struct {
	transformed map[*data.Vertex]struct{}
	f           *data.Figure
	epsilon     int
	doScaling	bool
}

func (mt *matrixTransformer) apply(v *data.Vertex, m pixel.Matrix) {
	if _, done := mt.transformed[v]; done {
		return
	}

	oldV := *v

	newPos := m.Project(v.PVec())
	v.X = newPos.X
	v.Y = newPos.Y
	v.Metadata.Reset()
	mt.transformed[v] = struct{}{}
	if mt.doScaling {
		log("apply %s => %s", oldV, v)
	}

	for _, candidate := range mt.f.GetConnectedVertices(v) {
		if _, done := mt.transformed[candidate]; done {
			continue
		}

		oldEdge := &data.Edge{A: &oldV, B: candidate}
		newEdge := &data.Edge{A: v, B: candidate}
		if !data.GoodRatio(oldEdge, newEdge, mt.epsilon) {
			// Candidate point also has to be moved to remain valid.
			// Get the minimal scale.
			scale := data.ScaleToPreserveRatio(oldEdge, newEdge, mt.epsilon)
			log("scale for %s: %f", candidate, scale)
			if mt.doScaling {
				m = m.Scaled(v.PVec(), scale)
			}
			mt.apply(candidate, m)
		}
	}
}

func Matrix(figure *data.Figure, v *data.Vertex, m pixel.Matrix, epsilon int) {
	MatrixScale(figure, v, m, epsilon, false)
}

func MatrixScale(figure *data.Figure, v *data.Vertex, m pixel.Matrix, epsilon int, doScaling bool) {
	(&matrixTransformer{
		transformed: make(map[*data.Vertex]struct{}),
		epsilon:     epsilon,
		f:           figure,
		doScaling: doScaling,
	}).apply(v, m)
}
