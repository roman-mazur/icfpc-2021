package algorithm

import (
	"log"
	"sort"

	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/fitness"
)

var GenerationSize = 4

type GenerationItem struct {
	Figure data.Figure
	Score  float64
}

type Generation []GenerationItem

func newGeneration(orig data.Figure, h data.Hole, ε, size, iter int) Generation {
	gen := make(Generation, 0, size)

	for i := 0; i < size; i++ {
		candidate := orig.Copy()
		applied := randomAlter(&candidate, ε)
		if candidate.IsValid(orig, ε) {
			log.Println(iter, " valid ", applied)
			gen = append(gen, GenerationItem{
				Figure: candidate,
				Score:  fitness.FitScore(candidate, h),
			})
		} else {
			log.Println(iter, " INVALID ", applied)
		}
	}

	sort.Slice(gen, func(i, j int) bool {
		return gen[i].Score < gen[j].Score
	})

	return gen
}

func Solve(f data.Figure, h data.Hole, ε, iter int) *data.Figure {
	base := f
	for i := 0; i < iter; i++ {
		generation := newGeneration(f, h, ε, GenerationSize, i)
		if len(generation) == 0 {
			return nil
		}
		base = generation[0].Figure
	}
	return &base
}
