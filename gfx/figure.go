package gfx

import (
	"fmt"
	"image/color"
	"math"

	goColor "github.com/gerow/go-color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"github.com/roman-mazur/icfpc-2021/data"
	"golang.org/x/image/colornames"
)

// I need to rework that
type FigureEntity struct {
	origFig *data.Figure
	fig     *data.Figure

	ε int

	selectable  bool
	selectedIdx int
	selectedVtx *data.Vertex

	showIndexes bool
	thickness   float64
	color       color.Color
}

func (fe *FigureEntity) Build() *imdraw.IMDraw {
	return BuildEdges(EdgeBuilder{
		edges:          fe.fig.Edges,
		baseColor:      colornames.Red,
		customRenderer: fe.render,
	})
}

func (fe *FigureEntity) BuildLabels(atlas *text.Atlas, m *pixel.Matrix) []*text.Text {
	var labels []*text.Text

	for i, e := range fe.fig.Edges {
		txt := text.New(m.Unproject(e.PLine().Center().Scaled(k)), atlas)
		txtContent := i

		if e.A == fe.selectedVtx || e.B == fe.selectedVtx {
			r, rMax := getLengthRatioWithMax(e, fe.origFig.Edges[i], fe.ε)

			txtContent = int(pixel.Clamp(r/rMax*100, 0, 100))
			txt.Color = getColorFromLengthRatio(r, rMax)
		}

		fmt.Fprintf(txt, "%d", txtContent)
		labels = append(labels, txt)
	}

	return labels
}

func (fe *FigureEntity) render(imd *imdraw.IMDraw, e *data.Edge, i int) {
	if len(fe.origFig.Edges) > 0 {
		r, rMax := getLengthRatioWithMax(e, fe.origFig.Edges[i], fe.ε)

		c := pixel.Clamp(r/rMax, 0, 1)
		if c >= 1 {
			imd.Color = colornames.Black
		} else {
			hsl := goColor.HSL{
				H: (1.0 / 3.0) - (c / 3),
				S: 1,
				L: 0.5,
			}

			rgb := hsl.ToRGB()
			imd.Color = pixel.RGB(rgb.R, rgb.G, rgb.B)
		}
	}

	v1 := e.A.PVec().Scaled(k)
	v2 := e.B.PVec().Scaled(k)

	imd.Push(v1, v2)

	if fe.selectable && (fe.selectedIdx >= 0 && fe.fig.Edges[fe.selectedIdx] == e) {
		imd.Line(fe.thickness + 2)
	} else {
		imd.Line(fe.thickness)
	}

	imd.Reset()

	if !fe.selectable {
		return
	}

	imd.Push(v1)
	if fe.selectedVtx == e.A {
		imd.Circle(5, 0)
	} else {
		imd.Circle(3, 0)
	}

	imd.Push(v2)
	if fe.selectedVtx == e.B {
		imd.Circle(5, 0)
	} else {
		imd.Circle(3, 0)
	}

	imd.Reset()
}

func (fe *FigureEntity) mvSelectedVtx(mousePos pixel.Vec) {
	if fe.selectedVtx == nil {
		return
	}

	fe.selectedVtx.X = math.Round(mousePos.X)
	fe.selectedVtx.Y = math.Round(mousePos.Y)
	fe.selectedVtx.Metadata.Reset()
}

func getLengthRatioWithMax(edge *data.Edge, origEdge *data.Edge, ε int) (float64, float64) {
	r := data.LengthRatio(origEdge, edge)
	rMax := float64(ε) / float64(1000000)

	return r, rMax
}

func getColorFromLengthRatio(r float64, rMax float64) color.Color {
	c := pixel.Clamp(r/rMax, 0, 1)
	if c >= 1 {
		return colornames.Black
	} else {
		hsl := goColor.HSL{
			H: (1.0 / 3.0) - (c / 3),
			S: 1,
			L: 0.5,
		}

		rgb := hsl.ToRGB()
		return pixel.RGB(rgb.R, rgb.G, rgb.B)
	}
}
