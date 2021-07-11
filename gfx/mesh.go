package gfx

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"github.com/roman-mazur/icfpc-2021/data"
)

type EdgeMesh struct {
	edges       []*data.Edge
	selected    *data.Edge
	selectable  bool
	thickness   float64
	color       color.Color
	showIndexes bool
}

func (mesh *EdgeMesh) Build() *imdraw.IMDraw {
	imd := imdraw.New(nil)

	for _, e := range mesh.edges {
		imd.Color = mesh.color

		imd.Push(
			e.A.PVec().Scaled(k),
			e.B.PVec().Scaled(k),
		)

		if mesh.selectable && mesh.selected == e {
			imd.Line(mesh.thickness + 2)
		} else {
			imd.Line(mesh.thickness)
		}

		imd.Reset()
	}

	return imd
}

func (mesh *EdgeMesh) BuildLabels(atlas *text.Atlas) []*text.Text {
	var labels []*text.Text

	for i, e := range mesh.edges {
		txt := text.New(e.PLine().Center().Scaled(k), atlas)
		fmt.Fprintf(txt, "%d", i)

		labels = append(labels, txt)
	}

	return labels
}
