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

	transform.Fold(pb.Figure, pb.Figure.Edges[30], transform.FoldLeft)
	transform.Matrix(pb.Figure, pb.Figure.Edges[22].B, pixel.IM.Moved(pixel.V(8, 2)), pb.Epsilon)
	transform.Matrix(pb.Figure, pb.Figure.Edges[5].B, pixel.IM.Moved(pixel.V(1, 3)), pb.Epsilon)
	transform.Matrix(pb.Figure, pb.Figure.Edges[0].B, pixel.IM.Moved(pixel.V(-3, 2)), pb.Epsilon)
	transform.Matrix(pb.Figure, pb.Figure.Edges[9].B, pixel.IM.Moved(pixel.V(1, 2)), pb.Epsilon)
	//transform.Matrix(pb.Figure, pb.Figure.Edges[1].B, pixel.IM.Moved(pixel.V(-1, 0)), pb.Epsilon)
	transform.Fold(pb.Figure, pb.Figure.Edges[23], transform.FoldRight, pb.Figure.Edges[27], pb.Figure.Edges[24], pb.Figure.Edges[17])

	//transform.Fold(pb.Figure, pb.Figure.Edges[3], transform.FoldRight)

	unfitEdges := cmd.Analyze(pb, original, false)

	cmd.WriteSolution(pb.Figure.Solution(), "2")

	gfx.DrawEdges(
		pixelgl.WindowConfig{
			Title:  "Ant",
			Bounds: pixel.R(0, 0, 1000, 800),
		},
		pb.Hole.Edges,
		original.Edges,
		pb.Figure.Edges,
		unfitEdges,
	)
}
