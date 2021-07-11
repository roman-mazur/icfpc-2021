package main

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/cmd"
	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/gfx"
	"github.com/roman-mazur/icfpc-2021/transform"
)

func main() {
	pb := data.ParseProblem("transform/testdata/endo.problem")
	original := pb.Figure.Copy()

	transform.Fold(pb.Figure, pb.Figure.Edges[8], transform.FoldLeft)
	transform.Fold(pb.Figure, pb.Figure.Edges[7], transform.FoldRight)
	transform.Rotate(pb.Figure, pb.Figure.Edges[38], math.Pi/8, pb.Epsilon)
	transform.Rotate(pb.Figure, pb.Figure.Edges[30], math.Pi/4, pb.Epsilon)
	transform.Rotate(pb.Figure, pb.Figure.Edges[54], math.Pi/8, pb.Epsilon)
	transform.Rotate(pb.Figure, pb.Figure.Edges[47], -math.Pi/4, pb.Epsilon)

	transform.Matrix(pb.Figure, pb.Figure.Edges[15].B, pixel.IM.Moved(pixel.V(0, -12)), pb.Epsilon)

	unfitEdges := cmd.Analyze(pb, original, true)

	cmd.WriteSolution(pb.Figure.Solution(), "5")

	gfx.DrawEdges(
		pixelgl.WindowConfig{
			Title:  "Endo",
			Bounds: pixel.R(0, 0, 1000, 800),
		},
		pb.Hole.Edges,
		original.Edges,
		pb.Figure.Edges,
		unfitEdges,
	)
}
