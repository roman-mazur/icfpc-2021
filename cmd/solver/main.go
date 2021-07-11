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
	"github.com/roman-mazur/icfpc-2021/fitness"
	"github.com/roman-mazur/icfpc-2021/gfx"
)

var (
	asService    = flag.Bool("as-service", false, "No UI")
	iterations   = flag.Int("iterations", 1000, "Number of iterations")
	genSize      = flag.Int("gen-size", 256, "Gen size")
	parallelGens = flag.Int("gen-parallel", 3, "Number of parallel generations")
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
	algorithm.NbParallel = *parallelGens

	pb := data.ParseProblem(problemPath)
	original := pb.Figure.Copy()
	origPb := data.Problem{
		Hole:    pb.Hole,
		Figure:  &original,
		Epsilon: pb.Epsilon,
	}

	algorithm.GenerationSize = 16 * len(pb.Figure.Vertices)

	bestMatch := algorithm.Solve(*pb.Figure, *pb.Hole, pb.Epsilon, *iterations)
	pb.Figure = &bestMatch.Figure

	unfit := cmd.Analyze(pb, *origPb.Figure, *asService)
	if len(unfit) == 0 {
		score := int(-1.0 / bestMatch.Score)
		log.Println("Score:", score)
		solutionName := fmt.Sprintf("%s-score-%f", strings.ReplaceAll(problemPath, "/", "_"), float64(score))
		if cmd.IsBetterSolution(solutionName, score) {
			cmd.WriteSolution(data.Solution{bestMatch.Figure.Vertices}, solutionName)
			fmt.Printf("Wrote %s\n", solutionName)
		} else {
			fmt.Printf("Didn't write %s: a better solution exists for score %d\n", solutionName, score)
		}

	}

	if !*asService {
		wasValid := len(unfit) == 0
		isValid := false

		for once := true; once || wasValid && !isValid; once = false {
			vis := gfx.NewVisualizer(pixelgl.WindowConfig{
				Title:  filepath.Base(problemPath),
				Bounds: pixel.R(0, 0, 1000, 800),
			}, &origPb)

			vis.PushFigure(pb.Figure, true, 2, true).PushEdges(unfit).Start()

			unfit = cmd.Analyze(pb, *pb.Figure, false)
			isValid = len(unfit) == 0
		}

		if len(unfit) == 0 {
			score := int(-1.0 / fitness.FitScore(*pb.Figure, *pb.Hole))
			log.Println("New score:", score)

			solutionName := fmt.Sprintf("%s-score-%f", strings.ReplaceAll(problemPath, "/", "_"), float64(score))
			if cmd.IsBetterSolution(solutionName, score) {
				cmd.WriteSolution(data.Solution{bestMatch.Figure.Vertices}, solutionName)
				fmt.Printf("Wrote %s\n", solutionName)
			} else {
				fmt.Printf("Didn't write %s: a better solution exists for score %d\n", solutionName, score)
			}
		}
	}
}
