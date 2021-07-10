package algorithm

import (
	"sort"
	"sync"

	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/fitness"
)

var GenerationSize = 4

type Generation []struct {
	Figure data.Figure
	Score  float64
}

func newGeneration(orig data.Figure, h data.Hole, ε, size int) Generation {
	gen := make(Generation, size)

	wg := sync.WaitGroup{}
	wg.Add(size)

	for i := 0; i < size; i++ {
		go (func(i int) {
			defer wg.Done()
			gen[i].Figure = orig.Copy()

			for valid := false; !valid; valid = gen[i].Figure.IsValid(orig, ε) {
				randomAlter(&gen[i].Figure)
			}
			gen[i].Score = fitness.FitScore(gen[i].Figure, h)
		})(i)
	}

	wg.Wait()

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
