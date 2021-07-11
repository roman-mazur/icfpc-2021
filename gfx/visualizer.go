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

type Visualizer struct {
	camPos  pixel.Vec
	camZoom float64

	winCfg     pixelgl.WindowConfig
	win        *pixelgl.Window
	asciiAtlas *text.Atlas

	pb        *data.Problem
	figures   []*FigureEntity
	miscEdges [][]*data.Edge

	targetFigure *FigureEntity
}

func NewVisualizer(cfg pixelgl.WindowConfig, pb *data.Problem) *Visualizer {
	return &Visualizer{
		winCfg:     cfg,
		camPos:     pixel.ZV,
		camZoom:    1.0,
		figures:    make([]*FigureEntity, 0),
		asciiAtlas: text.NewAtlas(basicfont.Face7x13, text.ASCII),
		pb:         pb,
	}
}

func (vis *Visualizer) Start() {
	pixelgl.Run(func() {
		win, err := pixelgl.NewWindow(vis.winCfg)
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

		vis.win = win
		txt := text.New(pixel.V(10, 10), vis.asciiAtlas)

		hole := BuildEdges(EdgeBuilder{
			edges:     vis.pb.Hole.Edges,
			thickness: 2,
		})

		grid := buildGrid(win.Bounds().Max.Scaled(1 / k))

		builtMiscEdges := []*imdraw.IMDraw{}
		for i, e := range vis.miscEdges {
			builtMiscEdges = append(builtMiscEdges, BuildEdges(EdgeBuilder{
				edges:     e,
				baseColor: colors[i+2%len(colors)],
				thickness: 2,
			}))
		}

		for !vis.win.Closed() && !vis.win.JustReleased(pixelgl.KeyEscape) {
			startTime := time.Now()

			// cam := pixel.IM.Scaled(vis.camPos, vis.camZoom).Moved(win.Bounds().Center().Sub(vis.camPos))
			// cam := pixel.IM.ScaledXY(pixel.ZV, pixel.V(1, -1)).Moved(win.Bounds().Center().Sub(vis.camPos))
			// vis.win.SetMatrix(cam)

			vis.win.Clear(colornames.Gray)
			vis.updateInputs()

			grid.Draw(win)
			hole.Draw(win)

			for _, be := range builtMiscEdges {
				be.Draw(win)
			}

			for _, f := range vis.figures {
				f.Build().Draw(vis.win)

				if f.showIndexes {
					for _, t := range f.BuildLabels(vis.asciiAtlas) {
						t.Draw(vis.win, pixel.IM.ScaledXY(t.Bounds().Center(), pixel.V(1, -1)))
					}
				}
			}

			txt.Draw(win, pixel.IM.ScaledXY(txt.Bounds().Center(), pixel.V(1, -1)))
			vis.win.Update()

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

func (vis *Visualizer) PushFigure(fig *data.Figure, selectable bool, thickness float64, showIndexes bool) *Visualizer {
	vis.figures = append(vis.figures, &FigureEntity{
		origFig:     vis.pb.Figure,
		fig:         fig,
		selectable:  selectable,
		selectedIdx: -1,
		thickness:   thickness,
		color:       colors[len(vis.figures)],
		showIndexes: showIndexes,
		ε:           vis.pb.Epsilon,
	})

	return vis
}

// func (vis *Visualizer) PushSolution( *data.Solution)

func (vis *Visualizer) PushEdges(edges ...[]*data.Edge) *Visualizer {

	vis.miscEdges = append(vis.miscEdges, edges...)
	return vis
}

// Dirty mouse selection
func (vis *Visualizer) updateInputs() {
	if !vis.win.MouseInsideWindow() {
		return
	}

	mousePos := vis.win.MousePosition().ScaledXY(pixel.V(1, -1))
	mousePos.Y += vis.win.Bounds().H() + marginTop

	if vis.win.MouseScroll() != pixel.ZV {
		vis.camZoom += vis.win.MouseScroll().Y * 0.2
	}

	if vis.win.Pressed(pixelgl.MouseButton3) && vis.win.MousePosition() != vis.win.MousePreviousPosition() {
		test := vis.win.MousePreviousPosition().Sub(vis.win.MousePosition()).Normal()
		vis.camPos.X += test.Y
		vis.camPos.Y -= test.X
	}

	if vis.win.JustPressed(pixelgl.MouseButton1) {
		vis.findTargetFigure(mousePos)
		return
	}

	if vis.win.JustReleased(pixelgl.MouseButton1) {
		vis.targetFigure = nil
		return
	}

	if vis.win.MousePosition() != vis.win.MousePreviousPosition() && vis.targetFigure != nil {
		vis.targetFigure.mvSelectedVtx(mousePos.Scaled(1 / k))
	}
}

func (vis *Visualizer) findTargetFigure(mousePos pixel.Vec) {
	found := false

	for _, f := range vis.figures {
		f.selectedIdx = -1
		f.selectedVtx = nil

		if !f.selectable || found {
			continue
		}

		for _, e := range f.fig.Edges {
			v1 := e.A.PVec().Scaled(k)
			v2 := e.B.PVec().Scaled(k)

			if pixel.C(v1, 5).Contains(mousePos) {
				f.selectedVtx = e.A
				found = true
			} else if pixel.C(v2, 5).Contains(mousePos) {
				f.selectedVtx = e.B
				found = true
			}

			if found {
				vis.targetFigure = f
				return
			}
		}
	}
}
