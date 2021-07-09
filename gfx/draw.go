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
			drawFigure(win, pb.Figure)

			win.Update()
		}
	})
}

func drawHole(win *pixelgl.Window, h *data.Hole) {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0.8, 0.8, 0.8)

	for _, v := range h.Vertices {
		imd.Push(pixel.V(float64(v.X)*k, float64(v.Y)*k))
	}

	// We add the first point again to close the drawing
	imd.Push(pixel.V(float64(h.Vertices[0].X)*k, float64(h.Vertices[0].Y)*k))
	// imd.Polygon(0)
	imd.Line(2)
	imd.Draw(win)
}

func drawFigure(win *pixelgl.Window, fig *data.Figure) {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(1.0, 0, 0)

	for _, v := range fig.Vertices {
		imd.Push(pixel.V(float64(v.X)*k, float64(v.Y)*k))
	}

	// imd.Line(2)
	imd.Circle(1, 3)
	imd.Draw(win)
}
