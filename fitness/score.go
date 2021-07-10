package fitness

import "github.com/roman-mazur/icfpc-2021/data"

// FitScore returns a fitness score of the Figure inside the Hole.
// The lower the better.
func FitScore(f data.Figure, h data.Hole) float64 {
	unfits := ListUnfits(f, h)
	if len(unfits) == 0 {
		return 0
	}

	return 1
}
