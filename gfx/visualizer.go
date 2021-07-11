package gfx

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/roman-mazur/icfpc-2021/data"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type Visualizer struct {
	winCfg pixelgl.WindowConfig
	win    *pixelgl.Window

	OnDrag func(e *data.Edge, mousePos pixel.Vec)

	meshes     []*EdgeMesh
	asciiAtlas *text.Atlas
}

func NewVisualizer(cfg pixelgl.WindowConfig) *Visualizer {
	return &Visualizer{
		winCfg:     cfg,
		meshes:     make([]*EdgeMesh, 0),
		asciiAtlas: text.NewAtlas(basicfont.Face7x13, text.ASCII),
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

		for !vis.win.Closed() && !vis.win.JustReleased(pixelgl.KeyEscape) {
			startTime := time.Now()

			vis.win.Clear(colornames.Gray)
			vis.updateInputs()

			for _, m := range vis.meshes {
				m.Build().Draw(vis.win)

				if m.showIndexes {
					for _, t := range m.BuildLabels(vis.asciiAtlas) {
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

func (vis *Visualizer) PushEdges(edges []*data.Edge, selectable bool, thickness float64, showIndexes bool) {
	vis.meshes = append(vis.meshes, &EdgeMesh{
		edges:       edges,
		selectable:  selectable,
		selected:    nil,
		thickness:   thickness,
		color:       colors[len(vis.meshes)],
		showIndexes: showIndexes,
	})
}

// Dirty mouse selection
func (vis *Visualizer) updateInputs() {
	if vis.win.Pressed(pixelgl.MouseButton1) && vis.win.MousePreviousPosition() != vis.win.MousePosition() {
		vis.OnDrag(vis.meshes[1].edges[5], vis.win.MousePosition())
	}

	// if !vis.win.JustReleased(pixelgl.MouseButton1) {
	// 	return
	// }

	// found := false
	// for _, m := range vis.meshes {
	// 	m.selected = nil

	// 	if !m.selectable || found {
	// 		continue
	// 	}

	// 	for _, e := range m.edges {
	// 		mousePos := vis.win.MousePosition().ScaledXY(pixel.V(1, -1))
	// 		mousePos.Y += vis.win.Bounds().H() + marginTop
	// 		l := pixel.L(e.A.PVec().Scaled(k), e.B.PVec().Scaled(k))

	// 		if l.IntersectRect(pixel.R(mousePos.X-10, mousePos.Y-10, mousePos.X+10, mousePos.Y+10)) == pixel.ZV {
	// 			continue
	// 		}

	// 		m.selected = e
	// 		found = true
	// 		break
	// 	}
	// }
}
