package gfx

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/roman-mazur/icfpc-2021/data"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

// This is deprecated, Visualizer should be used.

var (
	k         = 6.0 // temp scale
	marginTop = 6.0
	maxFPS    = 30

	flipVertically = pixel.IM.ScaledXY(
		pixel.V(0, 0),
		pixel.V(1, -1),
	)

	colors = []pixel.RGBA{
		pixel.RGB(0.8, 0.8, 0.8),
		pixel.RGB(1, 0, 0),
		pixel.RGB(0, 0, 1),
		pixel.RGB(0, 1, 0),
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
			flipVertically.Chained(
				pixel.IM.Moved(
					pixel.V(0, win.Bounds().H()+marginTop),
				),
			),
		)

		atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
		txt := text.New(pixel.V(10, 10), atlas)

		for !win.Closed() && !win.JustReleased(pixelgl.KeyEscape) {
			startTime := time.Now()

			win.Clear(colornames.Gray)
			drawFunc(win)
			txt.Draw(win, pixel.IM.ScaledXY(txt.Bounds().Center(), pixel.V(1, -1)))
			win.Update()

			elapsed := time.Since(startTime)
			toSleep := time.Duration(time.Second.Nanoseconds()/int64(maxFPS)) - elapsed

			txt.Clear()
			fmt.Fprintf(txt, "%d", time.Second/elapsed)

			if toSleep > 0 {
				time.Sleep(toSleep)
			}
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
			drawEdges(imd, e, 2)
			imd.Draw(win)
		}
		if len(edges) > 1 {
			drawEdgeNums(win, edges[1])
		}
	})
}

func newDraw() *imdraw.IMDraw {
	return imdraw.New(nil)
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
		pos := e.PLine().Center().Scaled(k)
		txt := text.New(pos, atlas)

		fmt.Fprintf(txt, "%d", i)
		txt.Draw(win, pixel.IM.ScaledXY(pos, pixel.V(1, -1)))
	}
}
