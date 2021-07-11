package gfx

import (
	"image/color"

	"github.com/faiface/pixel/imdraw"
	"github.com/roman-mazur/icfpc-2021/data"
)

type EdgeBuilder struct {
	edges     []*data.Edge
	baseColor color.Color
	thickness float64

	customRenderer func(*imdraw.IMDraw, *data.Edge, int)
}

func BuildEdges(eb EdgeBuilder) *imdraw.IMDraw {
	imd := imdraw.New(nil)

	if eb.baseColor == nil {
		eb.baseColor = colors[0]
	}

	for i, e := range eb.edges {
		imd.Color = eb.baseColor

		if eb.customRenderer == nil {
			imd.Push(
				e.A.PVec().Scaled(k),
				e.B.PVec().Scaled(k),
			)

			imd.Line(eb.thickness)
			imd.Reset()
		} else {
			eb.customRenderer(imd, e, i)
		}
	}

	return imd
}
