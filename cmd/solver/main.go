package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/cmd"
	"github.com/roman-mazur/icfpc-2021/cmd/solver/algorithm"
	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/gfx"
	"github.com/roman-mazur/icfpc-2021/profiling"
)

var (
	asService  = flag.Bool("as-service", false, "No UI")
	iterations = flag.Int("iterations", 1000, "Number of iterations")
	genSize    = flag.Int("gen-size", 1024, "Gen size")
)

func fatalUsage() {
	log.Fatalf("Usage: %s file.problem\n", os.Args[0])
}

func main() {
	flag.Parse()
	log.Println("Hello ICFP Contest!")
	problemPath := "problems/problem.1"
	if len(flag.Args()) >= 1 {
		problemPath = flag.Args()[0]
	}
	algorithm.GenerationSize = *genSize

	go profiling.Start()

	pb := data.ParseProblem(problemPath)
	original := pb.Figure.Copy()
	origPb := data.Problem{
		Hole:    pb.Hole,
		Figure:  &original,
		Epsilon: pb.Epsilon,
	}

	bestMatch := algorithm.Solve(*pb.Figure, *pb.Hole, pb.Epsilon, *iterations)
	pb.Figure = &bestMatch.Figure

	unfit := cmd.Analyze(pb, *origPb.Figure, *asService)
	if len(unfit) == 0 {
		score := int(-1.0 / bestMatch.Score)
		solutionName := fmt.Sprintf("%s-score-%f", strings.ReplaceAll(problemPath, "/", "_"), float64(score))
		if cmd.IsBetterSolution(solutionName, score) {
			cmd.WriteSolution(data.Solution{bestMatch.Figure.Vertices}, solutionName)
			fmt.Printf("Wrote %s\n", solutionName)
		} else {
			fmt.Printf("Didn't wrote %s: a better solution exists for score %d\n", solutionName, score)
		}

	}

	if !*asService {
		vis := gfx.NewVisualizer(pixelgl.WindowConfig{
			Title:  filepath.Base(problemPath),
			Bounds: pixel.R(0, 0, 1000, 800),
		}, &origPb)

		vis.PushFigure(pb.Figure, true, 2, true).PushEdges(unfit).Start()
	}
}
