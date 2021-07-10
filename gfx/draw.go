package gfx

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/roman-mazur/icfpc-2021/data"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

var (
	k         = 6.0 // temp scale
	posMatrix = pixel.IM.Moved(pixel.V(0, -5))

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
		drawEdgeNums(win, pb.Figure.Edges)
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
	res.SetMatrix(posMatrix)
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
			e.A.PVec().Scaled(k),
			e.B.PVec().Scaled(k),
		)

		imd.Line(thickness)
		imd.Reset()
	}
}

func drawEdgeNums(win *pixelgl.Window, edges []*data.Edge) {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	for i, e := range edges {
		txt := text.New(e.PLine().Center().Scaled(k), atlas)

		fmt.Fprintf(txt, "%d", i)
		txt.Draw(win, pixel.IM)
	}
}
