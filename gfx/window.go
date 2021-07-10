package gfx

import (
	"github.com/faiface/pixel/pixelgl"
)

func NewWin(cfg pixelgl.WindowConfig) {
	pixelgl.Run(func() {
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		for !win.Closed() {
			win.Update()
		}
	})
}
