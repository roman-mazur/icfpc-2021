package gfx

import (
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
	cam Camera

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
		winCfg: cfg,
		cam: Camera{
			pos: cfg.Bounds.Center(),
		},
		figures:    make([]*FigureEntity, 0),
		asciiAtlas: text.NewAtlas(basicfont.Face7x13, text.ASCII),
		pb:         pb,
	}
}

// Start returns true if there was an update, false if nothing moved.
func (vis *Visualizer) Start() (update bool) {
	update = false

	pixelgl.Run(func() {
		win, err := pixelgl.NewWindow(vis.winCfg)
		if err != nil {
			panic(err)
		}

		vis.win = win

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

			vis.cam.Update(vis.win)

			vis.win.Clear(colornames.Gray)
			if vis.updateInputs() {
				update = true
			}

			grid.Draw(win)

			hole := BuildEdges(EdgeBuilder{
				edges:     vis.pb.Hole.Edges,
				thickness: 2,
			})
			hole.Draw(win)

			for _, f := range vis.figures {
				f.Build().Draw(vis.win)

				if f.showIndexes {
					for _, t := range f.BuildLabels(vis.asciiAtlas, &vis.cam.matrix) {
						t.Draw(vis.win, vis.cam.matrix)
					}
				}
			}

			for _, be := range builtMiscEdges {
				be.Draw(win)
			}

			vis.win.Update()

			elapsed := time.Since(startTime)
			toSleep := time.Duration(time.Second.Nanoseconds()/int64(maxFPS)) - elapsed

			if toSleep > 0 {
				time.Sleep(toSleep)
			}
		}
	})

	return
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
func (vis *Visualizer) updateInputs() (update bool) {
	update = false
	if !vis.win.MouseInsideWindow() {
		return
	}

	mousePos := vis.cam.matrix.Unproject(vis.win.MousePosition())

	// if vis.win.MouseScroll() != pixel.ZV {
	// 	k += vis.win.MouseScroll().Y * 0.4
	// }

	if (vis.win.Pressed(pixelgl.MouseButton3) || vis.win.Pressed(pixelgl.MouseButton2)) && vis.win.MousePosition() != vis.win.MousePreviousPosition() {
		test := vis.win.MousePreviousPosition().Sub(vis.win.MousePosition()).Normal()
		vis.cam.pos.X += test.Y
		vis.cam.pos.Y -= test.X
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
		update = true
	}

	return
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
