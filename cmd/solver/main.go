package main

import (
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/cmd/solver/algorithm"
	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/gfx"
)

func fatalUsage() {
	log.Fatalf("Usage: %s file.problem\n", os.Args[0])
}

func main() {
	log.Println("Hello ICFP Contest!")
	if len(os.Args) < 2 {
		fatalUsage()
	}

	pb := data.ParseProblem(os.Args[1])
	original := pb.Figure.Copy()
	pb.Figure = algorithm.Solve(*pb.Figure, *pb.Hole, pb.Epsilon, 10)
	gfx.DrawEdges(
		pixelgl.WindowConfig{
			Title:  "Hello ICFP Contest!",
			Bounds: pixel.R(0, 0, 1000, 800),
		},
		pb.Hole.Edges,
		original.Edges,
		pb.Figure.Edges,
	)
}
