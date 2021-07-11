package gfx

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

func buildGrid(max pixel.Vec) *imdraw.IMDraw {
	imd := imdraw.New(nil)
	imd.Color = colornames.Black

	for x := 0; x < int(max.X); x += 1 {
		for y := 0; y < int(max.Y); y += 1 {
			imd.Push(pixel.V(float64(x), float64(y)).Scaled(k))
		}
	}

	imd.Circle(1, 0)
	imd.Reset()
	return imd
}
