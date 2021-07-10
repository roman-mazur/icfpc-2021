package algorithm

import (
	"sort"

	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/fitness"
)

var GenerationSize = 64

type GenerationItem struct {
	Figure data.Figure
	Score  float64
}

type Generation []GenerationItem

func newGeneration(orig data.Figure, h data.Hole, ε, size int) Generation {
	gen := make(Generation, 0, size)

	for i := 0; i < size; i++ {
		candidate := orig.Copy()
		randomAlter(&candidate)
		if candidate.IsValid(orig, ε) {
			gen = append(gen, GenerationItem{
				Figure: candidate,
				Score:  fitness.FitScore(candidate, h),
			})
		}
	}

	sort.Slice(gen, func(i, j int) bool {
		return gen[i].Score < gen[j].Score
	})

	return gen
}

func Solve(f data.Figure, h data.Hole, ε, iter int) *data.Figure {
	generation := newGeneration(f, h, ε, GenerationSize)

	if len(generation) == 0 {
		return nil
	}
	return &generation[0].Figure
}
