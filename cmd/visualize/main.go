package main

import (
	"os"
	"path/filepath"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/cmd"
	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/gfx"
)

func main() {
	pb := data.ParseProblem(os.Args[1])
	pbOrig := data.ParseProblem(os.Args[1])
	solution := data.ParseSolution(os.Args[2])
	solutionFigure := pb.Figure.Copy()

	for i := range solution.Vertices {
		solutionFigure.Vertices[i].X = solution.Vertices[i].X
		solutionFigure.Vertices[i].Y = solution.Vertices[i].Y
	}

	pb.Figure = &solutionFigure
	unfitEdges := cmd.Analyze(pb, solutionFigure, false)

	vis := gfx.NewVisualizer(pixelgl.WindowConfig{
		Title:  filepath.Base(os.Args[1]),
		Bounds: pixel.R(0, 0, 1000, 800),
	}, pbOrig)

	vis.PushFigure(pb.Figure, true, 2, true).PushEdges(unfitEdges).Start()
}
