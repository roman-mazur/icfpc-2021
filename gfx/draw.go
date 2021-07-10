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

var k = 6.0 // temp scale

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
	imd, _ := newIMDrawWithEdges(h.Edges, false)

	imd.SetColorMask(pixel.RGB(0.8, 0.8, 0.8))
	imd.Line(2)
	imd.Draw(win)
}

func drawFigure(win *pixelgl.Window, edges []*data.Edge) {
	imd, _ := newIMDrawWithEdges(edges, false)

	imd.SetColorMask(pixel.RGB(1.0, 0, 0))
	imd.Line(5)
	imd.Draw(win)
}

func newIMDrawWithEdges(edges []*data.Edge, idx bool) (*imdraw.IMDraw, *text.Text) {
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

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(pixel.V(50, 50), atlas)

	fmt.Fprintf(txt, "hello, world%s", "!")

	return imd, txt
}
