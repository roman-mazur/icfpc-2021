package main

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/gfx"
	"github.com/roman-mazur/icfpc-2021/transform"
)

func main() {
	pb := data.ParseProblem("transform/testdata/spider.problem")

//	original := pb.Figure.Copy()
//	transform.Fold(pb.Figure, pb.Figure.Edges[37], transform.FoldRight)
//	firstFold := pb.Figure.Copy()
//	transform.Fold(pb.Figure, pb.Figure.Edges[4], transform.FoldLeft)
	transform.Rotate(pb.Figure.Edges[8], -math.Pi/8)
	unfits := pb.Figure.ListUnfits(pb.Hole)
	unfitEdges := make([]*data.Edge, len(unfits))
	for i, unfit := range unfits {
		unfitEdges[i] = unfit.Edge
	}

	gfx.DrawEdges(
		pixelgl.WindowConfig{
			Title:  "Hello ICFP Contest!",
			Bounds: pixel.R(0, 0, 1000, 800),
		},
		pb.Hole.Edges,
		pb.Figure.Edges,
		[]*data.Edge{pb.Figure.Edges[8]},
		unfitEdges,
//		[]*data.Edge{original.Edges[9]},
//		[]*data.Edge{original.Edges[10]},
	)
}
