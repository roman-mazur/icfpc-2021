package gfx

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/data"
	"golang.org/x/image/colornames"
)

var k = 6.0

func DrawProblem(cfg pixelgl.WindowConfig, pb *data.Problem) {
	pixelgl.Run(func() {
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		for !win.Closed() {
			win.Clear(colornames.Gray)

			drawHole(win, pb.Hole)
			drawFigure(win, pb.Figure.Edges)

			win.Update()
		}
	})
}

func drawHole(win *pixelgl.Window, h *data.Hole) {
	imd := newIMDrawWithEdges(h.Edges)

	imd.SetColorMask(pixel.RGB(0.8, 0.8, 0.8))
	imd.Line(2)
	imd.Draw(win)
}

func drawFigure(win *pixelgl.Window, edges []*data.Edge) {
	imd := newIMDrawWithEdges(edges)

	imd.SetColorMask(pixel.RGB(1.0, 0, 0))
	imd.Line(5)
	imd.Draw(win)
}

func newIMDrawWithEdges(edges []*data.Edge) *imdraw.IMDraw {
	imd := imdraw.New(nil)

	for _, e := range edges {
		imd.Push(
			pixel.V(
				float64(e.A.X)*k,  // x1
				float64(e.A.Y)*k), // y1
			pixel.V(
				float64(e.B.X)*k,  // x2
				float64(e.B.Y)*k), // y2
		)
	}

	return imd
}
