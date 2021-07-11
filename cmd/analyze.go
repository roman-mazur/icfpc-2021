package cmd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/fitness"
)

func IsBetterSolution(problemPath string, score int) bool {
	var rgx = regexp.MustCompile("problem.(\\d+)")
	var submitted = make(map[int]int)

	matches := rgx.FindStringSubmatch(problemPath)
	if len(matches) < 2 {
		return false
	}

	problemId, err := strconv.Atoi(matches[1])
	if err != nil {
		return false
	}

	submittedFp := filepath.Join("solutions", "submitted.json")
	if f, err := os.Open(submittedFp); err == nil {
		defer f.Close()
		if err := json.NewDecoder(f).Decode(&submitted); err != nil {
			return false
		}
	}

	bestScore, ok := submitted[problemId]
	if !ok {
		return true
	}
	log.Println("Problem", problemId, "best score is", bestScore)
	return bestScore > score
}

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
