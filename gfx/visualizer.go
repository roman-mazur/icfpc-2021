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
	winCfg     pixelgl.WindowConfig
	win        *pixelgl.Window
	asciiAtlas *text.Atlas

	pb        *data.Problem
	figures   []*FigureEntity
	miscEdges [][]*data.Edge

	draggedVtx *data.Vertex

	OnDrag       func(e *data.Edge, mousePos pixel.Vec)
	OnVertexDrag func(v *data.Vertex, mousePos pixel.Vec)
}

func NewVisualizer(cfg pixelgl.WindowConfig, pb *data.Problem) *Visualizer {
	return &Visualizer{
		winCfg:     cfg,
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

		builtMiscEdges := []*imdraw.IMDraw{}
		for i, e := range vis.miscEdges {
			builtMiscEdges = append(builtMiscEdges, BuildEdges(EdgeBuilder{
				edges:     e,
				baseColor: colors[i+2%len(colors)],
				thickness: 2,
			}))
		}

		grid := buildGrid(win.Bounds().Max.Scaled(1 / k))

		for !vis.win.Closed() && !vis.win.JustReleased(pixelgl.KeyEscape) {
			startTime := time.Now()

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

func (vis *Visualizer) PushEdges(edges ...[]*data.Edge) *Visualizer {

	vis.miscEdges = append(vis.miscEdges, edges...)
	return vis
}

func (vis *Visualizer) buildStaticObjects() {

}

// Dirty mouse selection
func (vis *Visualizer) updateInputs() {

	// return

	if !vis.win.MouseInsideWindow() {
		return
	}

	// if vis.win.JustPressed()

	mousePos := vis.win.MousePosition().ScaledXY(pixel.V(1, -1))
	mousePos.Y += vis.win.Bounds().H() + marginTop

	found := false
	for _, f := range vis.figures {
		f.selectedIdx = -1
		f.selectedVtx = nil

		if !f.selectable || found {
			continue
		}

		for _, e := range f.fig.Edges {
			// mousePos := vis.win.MousePosition()
			// l := pixel.L(e.A.PVec().Scaled(k), e.B.PVec().Scaled(k))

			// if l.IntersectCircle(pixel.C(mousePos, 5)) == pixel.ZV {
			// 	continue
			// }

			v1 := e.A.PVec().Scaled(k)
			v2 := e.B.PVec().Scaled(k)

			if pixel.C(v1, 5).Contains(mousePos) {
				f.selectedVtx = e.A
				found = true
				continue
			} else if pixel.C(v2, 5).Contains(mousePos) {
				f.selectedVtx = e.B
				found = true
				continue
			}

			// f.selectedIdx = i
			// found = true
			// break
		}

		if vis.OnVertexDrag != nil &&
			f.selectedVtx != nil &&
			vis.win.Pressed(pixelgl.MouseButton1) &&
			vis.win.MousePreviousPosition() != vis.win.MousePosition() {
			vis.OnVertexDrag(f.selectedVtx, mousePos.Scaled(1/k))
		}
	}
}
