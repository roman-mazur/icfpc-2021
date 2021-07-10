package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/gfx"
	"github.com/roman-mazur/icfpc-2021/transform"
)

func main() {
	pb := data.ParseProblem("transform/testdata/spider.problem")
	transform.Fold(pb.Figure, pb.Figure.Edges[37], transform.FoldRight)
	firstFold := pb.Figure.Copy()
	transform.Fold(pb.Figure, pb.Figure.Edges[4], transform.FoldLeft)
	secondFold := pb.Figure.Copy()
	gfx.DrawEdges(
		pixelgl.WindowConfig{
			Title:  "Hello ICFP Contest!",
			Bounds: pixel.R(0, 0, 1000, 800),
		},
		pb.Hole.Edges,
		firstFold.Edges,
		secondFold.Edges,
		[]*data.Edge{pb.Figure.Edges[37]},
		[]*data.Edge{pb.Figure.Edges[4]},
	)
}
