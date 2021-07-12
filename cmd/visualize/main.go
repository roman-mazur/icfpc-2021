package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/cmd"
	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/fitness"
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
		solutionFigure.Vertices[i].Metadata.Reset()
	}

	pb.Figure = &solutionFigure
	unfit := cmd.Analyze(pb, solutionFigure, false)

	var wasUpdated = true
	var adjustedScore = fitness.FitScore(*pb.Figure, *pb.Hole)

	for once := true; once || wasUpdated && len(unfit) > 0; once = false {
		vis := gfx.NewVisualizer(pixelgl.WindowConfig{
			Title:  filepath.Base(os.Args[1]),
			Bounds: pixel.R(0, 0, 1000, 800),
		}, pbOrig)

		wasUpdated = vis.PushFigure(pb.Figure, true, 2, true).PushEdges(unfit).Start()

		adjustedScore = fitness.FitScore(*pb.Figure, *pb.Hole)
		log.Println("New score: ", int(-1.0/adjustedScore))
		unfit = cmd.Analyze(pb, *pb.Figure, false)
	}

	if len(unfit) == 0 && pb.Figure.IsValid(*pb.Figure, pb.Epsilon) {
		score := int(-1.0 / adjustedScore)
		log.Println("New score:", score)

		solutionName := fmt.Sprintf("%s-score-%f", strings.ReplaceAll(os.Args[1], "/", "_"), float64(score))
		if cmd.IsBetterSolution(solutionName, score) {
			cmd.WriteSolution(data.Solution{pb.Figure.Vertices}, solutionName)
			fmt.Printf("Wrote %s\n", solutionName)
		} else {
			fmt.Printf("Didn't write %s: a better solution exists for score %d\n", solutionName, score)
		}
	}

}
