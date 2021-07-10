package main

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/fitness"
	"github.com/roman-mazur/icfpc-2021/gfx"
	"github.com/roman-mazur/icfpc-2021/transform"
)

func main() {
	pb := data.ParseProblem("transform/testdata/ant.problem")
	original := pb.Figure.Copy()

	transform.Fold(pb.Figure, pb.Figure.Edges[37], transform.FoldRight)
	transform.Fold(pb.Figure, pb.Figure.Edges[4], transform.FoldLeft)
	//	transform.Rotate(pb.Figure.Edges[8], -math.Pi/8)

	if !pb.Figure.IsValid(original, pb.Epsilon) {
		log.Fatal("incorrect figure")
	}

	unfits := fitness.ListUnfits(*pb.Figure, *pb.Hole)
	unfitEdges := make([]*data.Edge, len(unfits))
	for i, unfit := range unfits {
		unfitEdges[i] = unfit.Edge
	}

	gfx.DrawEdges(
		pixelgl.WindowConfig{
			Title:  "Spider",
			Bounds: pixel.R(0, 0, 1000, 800),
		},
		pb.Hole.Edges,
		pb.Figure.Edges,
		unfitEdges,
		//		[]*data.Edge{original.Edges[9]},
		//		[]*data.Edge{original.Edges[10]},
	)
}
