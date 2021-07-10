package cmd

import (
	"log"

	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/fitness"
)

func Analyze(problem *data.Problem, original data.Figure, failFast bool) (unfitEdges []*data.Edge) {
	if !problem.Figure.IsValid(original, problem.Epsilon) {
		if failFast {
			log.Fatal("incorrect figure")
		}
		log.Println("incorrect figure")
	}
	unfits := fitness.ListUnfits(*problem.Figure, *problem.Hole)
	unfitEdges = make([]*data.Edge, len(unfits))
	for i, unfit := range unfits {
		unfitEdges[i] = unfit.Edge
	}
	return
}