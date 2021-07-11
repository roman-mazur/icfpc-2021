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
	pb := data.ParseProblem("transform/testdata/ant.problem")
	original := pb.Figure.Copy()

	transform.Fold(pb.Figure, pb.Figure.Edges[37], transform.FoldRight)
	transform.Fold(pb.Figure, pb.Figure.Edges[4], transform.FoldLeft)
	transform.Rotate(pb.Figure, pb.Figure.Edges[20], -math.Pi/4, pb.Epsilon)
	transform.Rotate(pb.Figure, pb.Figure.Edges[26], math.Pi*0.75, pb.Epsilon)
	transform.Rotate(pb.Figure, pb.Figure.Edges[10], math.Pi/2, pb.Epsilon)
	transform.Rotate(pb.Figure, pb.Figure.Edges[34], math.Pi*0.9, pb.Epsilon)
	transform.Rotate(pb.Figure, pb.Figure.Edges[36], math.Pi, pb.Epsilon)
	transform.Rotate(pb.Figure, pb.Figure.Edges[24], math.Pi*0.4, pb.Epsilon)

	unfitEdges := cmd.Analyze(pb, original, true)

	cmd.WriteSolution(pb.Figure.Solution(), "3")

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
