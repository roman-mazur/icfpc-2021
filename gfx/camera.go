package gfx

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera struct {
	pos    pixel.Vec
	matrix pixel.Matrix
}

func (c *Camera) Update(win *pixelgl.Window) {
	c.matrix = pixel.IM.ScaledXY(
		pixel.ZV,
		pixel.V(1, -1),
	).Moved(pixel.V(0, win.Bounds().Max.Y)).Moved(win.Bounds().Center().Sub(c.pos))

	win.SetMatrix(c.matrix)
}
