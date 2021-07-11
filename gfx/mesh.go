package gfx

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"github.com/roman-mazur/icfpc-2021/data"
	"golang.org/x/image/colornames"
)

// I need to rework that
type EdgeMesh struct {
	origEdges   []*data.Edge
	edges       []*data.Edge
	selectedIdx int
	selectable  bool
	thickness   float64
	color       color.Color
	showIndexes bool
}

func (mesh *EdgeMesh) Build() *imdraw.IMDraw {
	imd := imdraw.New(nil)

	for i, e := range mesh.edges {
		imd.Color = mesh.color

		if len(mesh.origEdges) > 0 {
			ε := 150000 // Still need to pass the real epsilon here

			if data.LengthRatio(mesh.origEdges[i], e) > float64(ε)/float64(1000000) {
				// fmt.Printf("Ratio: %f\n", data.LengthRatio(mesh.origEdges[i], e))
				// fmt.Printf("Threshhold: %f\n", float64(ε)/float64(1000000))

				// not valid
				imd.Color = colornames.Black
			}
		}

		imd.Push(
			e.A.PVec().Scaled(k),
			e.B.PVec().Scaled(k),
		)

		if mesh.selectable && (mesh.selectedIdx >= 0 && mesh.edges[mesh.selectedIdx] == e) {
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
