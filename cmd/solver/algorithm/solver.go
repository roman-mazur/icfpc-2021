package algorithm

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"

	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/fitness"
)

var GenerationSize = 1024
var NbParallel = 1

type GenerationItem struct {
	Id        int
	Figure    data.Figure
	Flattened data.Figure
	Score     float64
}

type Generation []GenerationItem

func newGeneration(wg *sync.WaitGroup, parents []GenerationItem, h data.Hole, ε, size, iter int) Generation {
	gen := make(Generation, size+len(parents))
	wg.Add(size)

	for i := len(parents); i < size+len(parents); i++ {
		go (func(i int) {
			defer wg.Done()
			parent := parents[rand.Intn(len(parents))]
			candidate := parent.Figure.Copy()

			for isValid, attempt := false, 0; !isValid && attempt < 10; attempt++ {
				if len(candidate.Edges) == 0 {
					panic(fmt.Errorf("bad candidate data %d", i))
				}
				applied := randomAlter(&candidate, &h, ε)
				isValid = candidate.IsValid(parent.Figure, ε)
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

	for i := 0; i < len(parents); i++ {
		gen[i] = parents[i]
	}

	return gen
}

func Solve(f data.Figure, h data.Hole, ε, iter int) (result GenerationItem) {
	var selection = []GenerationItem{}
	result.Id = -1
	allParents := []GenerationItem{{Figure: f, Score: fitness.FitScore(f, h)}}
	bestScore := 0.0
	dislikes := 0
	worstGen := 0
	worstScore := 0.0
	noChangeSince := 0

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		return
	}()

	for i := 0; i < iter && noChangeSince < max(iter/5, 300); i++ {
		log.Println("New generation", i, "/", iter, "- gen size:", GenerationSize, "- lastChange:", noChangeSince, "- dislikes:", dislikes, "best / worst generation score:", bestScore, "/", worstScore)
		wg := new(sync.WaitGroup)

		var gens = make([]Generation, NbParallel)
		for j := 0; j < NbParallel; j++ {
			ancestors := append(selection, allParents...)

			if worstGen == j && shouldEvictGen(i, iter) {
				ancestors = []GenerationItem{{Figure: f, Score: fitness.FitScore(f, h)}}
			}
			gens[j] = newGeneration(wg, ancestors, h, ε, GenerationSize, i)
			if len(gens[j]) == 0 {
				return
			}
		}
		allParents = append(allParents, selection...)

		wg.Wait()
		worstScore = gens[0][0].Score
		for j := 0; j < NbParallel; j++ {
			sort.Slice(gens[j], func(l, m int) bool {
				return gens[j][l].Score < gens[j][m].Score
			})
			if gens[j][0].Score > worstScore {
				worstScore = gens[j][0].Score
				worstGen = j
			}
		}

		selection = make([]GenerationItem, 0, NbParallel*GenerationSize/64)
		for j := 0; j < NbParallel; j++ {
			selection = append(selection, gens[j][0:max(GenerationSize/64, 1)]...)
		}
		sort.Slice(selection, func(i, j int) bool {
			return selection[i].Score < selection[j].Score
		})
		if selection[0].Score < bestScore {
			bestScore = selection[0].Score
		}

		if dislikes > 0 {
			noChangeSince++
		}
		for _, res := range selection {
			if result.Id == -1 || (res.Score <= result.Score && res.Flattened.IsValid(f, ε)) {
				noChangeSince = 0
				result.Figure = res.Flattened
				result.Score = res.Score
				fmt.Println("New result")

				if result.Id == -1 && res.Flattened.IsValid(f, ε) || result.Id == 0 {
					result.Id = 0 // Always set something as a result.
					fmt.Println("New valid result")

					dislikes = int(-1.0 / result.Score)
				}
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

func shouldEvictGen(i, iter int) bool {
	return false
}
