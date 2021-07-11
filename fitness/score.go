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
		return -1.0 / Dislikes(f, h)
	}

	nbVertices := len(h.Vertices)

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
		score += float64(nbVertices)
	}

	return score / float64(nbVertices)
}

func Dislikes(f data.Figure, h data.Hole) (score float64) {
	for _, hv := range h.Vertices {
		vertScore := math.MaxFloat64

		for _, v := range f.Vertices {

			dist := data.Edge{A: &v, B: &hv}.SqLength()

			if dist < vertScore {
				vertScore = dist
			}
		}

		score += vertScore
	}

	return
}
