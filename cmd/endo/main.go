package main

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/cmd"
	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/gfx"
	"github.com/roman-mazur/icfpc-2021/transform"
)

func solution(pb *data.Problem) {
	transform.Fold(pb.Figure, pb.Figure.Edges[8], transform.FoldLeft)
	transform.Fold(pb.Figure, pb.Figure.Edges[7], transform.FoldRight)
	transform.Rotate(pb.Figure, pb.Figure.Edges[38], math.Pi/8, pb.Epsilon)
	transform.Rotate(pb.Figure, pb.Figure.Edges[30], math.Pi/4, pb.Epsilon)
	transform.Rotate(pb.Figure, pb.Figure.Edges[54], math.Pi/8, pb.Epsilon)
	transform.Rotate(pb.Figure, pb.Figure.Edges[47], -math.Pi/4, pb.Epsilon)

	transform.Matrix(pb.Figure, pb.Figure.Edges[15].B, pixel.IM.Moved(pixel.V(0, -12)), pb.Epsilon)
}

func experiment(pb *data.Problem, second *data.Figure) {
	fmt.Println("experiment", pb.Figure.Edges[38])
	transform.RotateScale(pb.Figure, pb.Figure.Edges[38], math.Pi/8, pb.Epsilon, false)
	transform.RotateScale(second, second.Edges[38], math.Pi/8, pb.Epsilon, true)
}

func main() {
	pb := data.ParseProblem("transform/testdata/endo.problem")
	original := pb.Figure.Copy()

	//second := pb.Figure.Copy()
	//experiment(pb, &second)

	solution(pb)

	unfitEdges := cmd.Analyze(pb, original, false)

	cmd.WriteSolution(pb.Figure.Solution(), "5")

	gfx.DrawEdges(
		pixelgl.WindowConfig{
			Title:  "Endo",
			Bounds: pixel.R(0, 0, 1000, 800),
		},
		pb.Hole.Edges,
		original.Edges,
		//second.Edges,
		pb.Figure.Edges,
		unfitEdges,
	)
}
