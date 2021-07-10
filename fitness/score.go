package fitness

import (
	"math"

	"github.com/roman-mazur/icfpc-2021/data"
)

// FitScore returns a fitness score of the Figure inside the Hole.
// The lower the better.
func FitScore(f data.Figure, h data.Hole) (score float64) {
	unfits := ListUnfits(f, h)
	if len(unfits) == 0 {
		return 0
	}

	for _, u := range unfits {
		for _, v := range u.Vertices {
			vertScore := math.MaxFloat64

			for _, hv := range h.Vertices {
				dist := data.Edge{A: v, B: &hv}.SqLength()

				if dist < vertScore {
					vertScore = dist
				}
			}

			score += vertScore
		}
	}

	return score / float64(len(h.Vertices))
}
