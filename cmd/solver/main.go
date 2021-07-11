package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/cmd"
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

	iteration := 1000

	pb := data.ParseProblem(os.Args[1])
	if len(os.Args) > 2 {
		iteration, _ = strconv.Atoi(os.Args[2])
	}
	if len(os.Args) > 3 {
		algorithm.GenerationSize, _ = strconv.Atoi(os.Args[3])
	}
	original := pb.Figure.Copy()
	bestMatch := algorithm.Solve(*pb.Figure, *pb.Hole, pb.Epsilon, iteration)
	pb.Figure = &bestMatch.Figure

	unfit := cmd.Analyze(pb, original, false)
	if len(unfit) == 0 {
		solutionName := fmt.Sprintf("%s-score-%f", strings.ReplaceAll(os.Args[1], "/", "_"), -1.0/bestMatch.Score)
		cmd.WriteSolution(data.Solution{bestMatch.Figure.Vertices}, solutionName)
	}

	os.Exit(0)
	gfx.DrawEdges(
		pixelgl.WindowConfig{
			Title:  filepath.Base(os.Args[1]),
			Bounds: pixel.R(0, 0, 1000, 800),
		},
		pb.Hole.Edges,
		original.Edges,
		pb.Figure.Edges,
		unfit,
	)
}
