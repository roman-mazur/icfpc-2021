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
	Id     int
	Figure data.Figure
	Score  float64
}

type Generation []GenerationItem

func newGeneration(parents []GenerationItem, h data.Hole, ε, size, iter int) Generation {
	gen := make(Generation, size)
	wg := new(sync.WaitGroup)
	wg.Add(size)

	for i := 0; i < size; i++ {
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

				score := fitness.FitScore(candidate, h)
				log.Println(iter, i, " valid ", score, applied)
				gen[i] = GenerationItem{
					Id:     i,
					Figure: candidate,
					Score:  score,
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
	result.Figure = f

	for i := 0; i < iter; i++ {
		generation := newGeneration(append(selection, result), h, ε, GenerationSize, i)
		if len(generation) == 0 {
			break
		}
		selection = generation[0 : GenerationSize/10]

		for _, res := range selection {
			flattened := res.Figure.FlattenToGrid()
			if flattened.IsValid(f, ε) && fitness.Fit(flattened, h) {
				score := fitness.FitScore(flattened, h)
				if score > result.Score {
					continue
				}

				result.Figure = flattened
				result.Score = score

				log.Println("Intermediary dislikes", -1.0/result.Score)
			}
		}
	}
	log.Println("Number of dislikes", -1.0/result.Score)
	return
}
