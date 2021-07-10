package gfx

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/data"
	"golang.org/x/image/colornames"
)

var (
	k         = 6.0 // temp scale
	marginTop = 6.0

	colors = []pixel.RGBA{
		pixel.RGB(0.8, 0.8, 0.8),
		pixel.RGB(1, 0, 0),
		pixel.RGB(0, 1, 0),
		pixel.RGB(0, 0, 1),
		pixel.RGB(0, 1, 1),
		pixel.RGB(1, 1, 0),
		pixel.RGB(0, 0, 0),
	}
)

func drawInWindow(cfg pixelgl.WindowConfig, drawFunc func(window *pixelgl.Window)) {
	pixelgl.Run(func() {
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		// Trick to invert origin
		win.SetMatrix(
			pixel.IM.ScaledXY(
				pixel.V(0, 0),
				pixel.V(1, -1),
			).Chained(
				pixel.IM.Moved(
					pixel.V(0, win.Bounds().H()+marginTop),
				),
			),
		)

		for !win.Closed() {
			win.Clear(colornames.Gray)
			drawFunc(win)
			win.Update()
		}
	})
}

func DrawProblem(cfg pixelgl.WindowConfig, pb *data.Problem) {
	drawInWindow(cfg, func(win *pixelgl.Window) {
		drawHole(pb.Hole).Draw(win)
		drawFigure(pb.Figure).Draw(win)
	})
}

func DrawEdges(cfg pixelgl.WindowConfig, edges ...[]*data.Edge) {
	drawInWindow(cfg, func(win *pixelgl.Window) {
		for i, e := range edges {
			imd := newDraw()
			imd.SetColorMask(colors[i%len(colors)])
			thickness := 5.0
			if i == 0 {
				thickness = 2.0
			}
			drawEdges(imd, e, thickness)
			imd.Draw(win)
		}
	})
}

func newDraw() *imdraw.IMDraw {
	res := imdraw.New(nil)
	return res
}

func drawHole(hole *data.Hole) *imdraw.IMDraw {
	imd := newDraw()
	imd.SetColorMask(pixel.RGB(0.8, 0.8, 0.8))
	drawEdges(imd, hole.Edges, 2)
	return imd
}

func drawFigure(f *data.Figure) *imdraw.IMDraw {
	imd := newDraw()
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
