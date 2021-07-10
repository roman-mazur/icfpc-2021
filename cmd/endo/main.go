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

	//transform.Fold(pb.Figure, pb.Figure.Edges[18], transform.FoldRight)
	transform.Matrix(pb.Figure, pb.Figure.Edges[38].B, pixel.IM.Rotated(pb.Figure.Edges[38].A.PVec(), math.Pi/16), pb.Epsilon)
	transform.Matrix(pb.Figure, pb.Figure.Edges[30].B, pixel.IM.Rotated(pb.Figure.Edges[30].A.PVec(), math.Pi/2), pb.Epsilon)

	unfitEdges := cmd.Analyze(pb, original, true)

	cmd.WriteSolution(pb.Figure.Solution(), 5)

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
