package gfx

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/data"
	"golang.org/x/image/colornames"
)

var k = 6.0 // temp scale

func DrawProblem(cfg pixelgl.WindowConfig, pb *data.Problem) {
	pixelgl.Run(func() {
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		for !win.Closed() {
			win.Clear(colornames.Gray)

			drawHole(pb.Hole).Draw(win)
			drawFigure(pb.Figure).Draw(win)

			win.Update()
		}
	})
}

func drawHole(hole *data.Hole) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.SetColorMask(pixel.RGB(0.8, 0.8, 0.8))
	drawEdges(imd, hole.Edges, 2)
	return imd
}

func drawFigure(f *data.Figure) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.SetColorMask(pixel.RGB(1.0, 0, 0))
	drawEdges(imd, f.Edges, 5)
	return imd
}

func drawEdges(imd *imdraw.IMDraw, edges []*data.Edge, thickness float64) {
	for _, e := range edges {
		imd.Push(
			pixel.V(
				float64(e.A.X)*k,  // x1
				float64(e.A.Y)*k), // y1
			pixel.V(
				float64(e.B.X)*k,  // x2
				float64(e.B.Y)*k), // y2
		)
		imd.Line(thickness)
		imd.Reset()
	}
}
