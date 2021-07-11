package algorithm

import (
	"log"
	"math/rand"
	"sort"
	"sync"

	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/fitness"
)

var GenerationSize = 1024

type GenerationItem struct {
	Id        int
	Figure    data.Figure
	Flattened data.Figure
	Score     float64
}

type Generation []GenerationItem

func newGeneration(parents []GenerationItem, h data.Hole, ε, size, iter int) Generation {
	gen := make(Generation, size+len(parents))
	wg := new(sync.WaitGroup)
	wg.Add(size)

	for i := 0; i < len(parents); i++ {
		gen[i] = parents[i]
	}

	for i := len(parents); i < size+len(parents); i++ {
		go (func(i int) {
			defer wg.Done()
			parent := parents[rand.Intn(len(parents))]
			candidate := parent.Figure.Copy()

			for isValid := false; !isValid; {
				applied := randomAlter(&candidate, &h, ε)
				isValid = true // Until stretches are implemented, `candidate.IsValid(parent.Figure, ε)` will always be true
				if !isValid {
					log.Println(iter, i, " INVALID, retrying ", applied)
					continue
				}

				flattened := candidate.FlattenToGrid()
				score := fitness.FitScore(flattened, h)
				gen[i] = GenerationItem{
					Id:        i,
					Figure:    candidate,
					Flattened: flattened,
					Score:     score,
				}
			}
		})(i)
	}

	wg.Wait()

	sort.Slice(gen, func(i, j int) bool {
		return gen[i].Score < gen[j].Score
	})

	return gen
}

func Solve(f data.Figure, h data.Hole, ε, iter int) (result GenerationItem) {
	selection := []GenerationItem{}
	parents := []GenerationItem{{Figure: f, Score: fitness.FitScore(f, h)}}
	bestScore := 0.0
	dislikes := 0

	for i := 0; i < iter; i++ {
		log.Println("New generation", i, "/", iter, "- gen size:", GenerationSize, "- dislikes:", dislikes, "best generation score:", bestScore)
		generation := newGeneration(append(selection, parents...), h, ε, GenerationSize, i)
		if len(generation) == 0 {
			break
		}
		parents = append(parents, selection...)
		selection = generation[0:max(GenerationSize/64, 1)]
		bestScore = selection[0].Score

		for _, res := range selection {
			if res.Score <= result.Score && res.Flattened.IsValid(f, ε) {
				result.Figure = res.Flattened
				result.Score = res.Score
				dislikes = int(-1.0 / result.Score)
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
