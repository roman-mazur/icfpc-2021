package main

import (
	"encoding/json"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/fitness"
	"github.com/roman-mazur/icfpc-2021/gfx"
	"github.com/roman-mazur/icfpc-2021/transform"
)

func main() {
	pb := data.ParseProblem("transform/testdata/ant.problem")
	original := pb.Figure.Copy()

	transform.Fold(pb.Figure, pb.Figure.Edges[37], transform.FoldRight)
	transform.Fold(pb.Figure, pb.Figure.Edges[4], transform.FoldLeft)
	transform.Rotate(pb.Figure.Edges[20], -math.Pi/4)
	transform.Rotate(pb.Figure.Edges[26], math.Pi*0.75)
	transform.Rotate(pb.Figure.Edges[10], math.Pi/2)
	transform.Rotate(pb.Figure.Edges[34], math.Pi*0.9)
	transform.Rotate(pb.Figure.Edges[36], math.Pi)
	transform.Rotate(pb.Figure.Edges[24], math.Pi*0.4)

	if !pb.Figure.IsValid(original, pb.Epsilon) {
		log.Fatal("incorrect figure")
	}

	writeSolution(pb.Figure.Solution())

	unfits := fitness.ListUnfits(*pb.Figure, *pb.Hole)
	unfitEdges := make([]*data.Edge, len(unfits))
	for i, unfit := range unfits {
		unfitEdges[i] = unfit.Edge
	}

	gfx.DrawEdges(
		pixelgl.WindowConfig{
			Title:  "Spider",
			Bounds: pixel.R(0, 0, 1000, 800),
		},
		pb.Hole.Edges,
		pb.Figure.Edges,
		unfitEdges,
		//		[]*data.Edge{original.Edges[9]},
		//		[]*data.Edge{original.Edges[10]},
	)
}

func writeSolution(sol data.Solution) {
	if solFile, err := os.Create(filepath.Join("solutions", "3.json")); err == nil {
		defer solFile.Close()
		if err := json.NewEncoder(solFile).Encode(sol); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}
