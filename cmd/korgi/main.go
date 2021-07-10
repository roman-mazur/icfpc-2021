package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/cmd"
	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/gfx"
	"github.com/roman-mazur/icfpc-2021/transform"
)

func main() {
	pb := data.ParseProblem("transform/testdata/korgi.problem")
	original := pb.Figure.Copy()

	transform.Fold(pb.Figure, pb.Figure.Edges[27], transform.FoldRight)
	transform.Fold(pb.Figure, pb.Figure.Edges[3], transform.FoldLeft)
	transform.Fold(pb.Figure, pb.Figure.Edges[30], transform.FoldLeft)

	unfitEdges := cmd.Analyze(pb, original, true)

	cmd.WriteSolution(pb.Figure.Solution(), 2)

	gfx.DrawEdges(
		pixelgl.WindowConfig{
			Title:  "Ant",
			Bounds: pixel.R(0, 0, 1000, 800),
		},
		pb.Hole.Edges,
		pb.Figure.Edges,
		unfitEdges,
	)
}
