package data

import "math"

const million = float64(1000000)

// IsValid tells if the figure f is a valid figure regarding the original figure and ε
func (f Figure) IsValid(original Figure, ε int) bool {
	if len(f.Edges) != len(original.Edges) ||
		len(f.Vertices) != len(original.Vertices) {
		return false
	}

	for i, newEdge := range f.Edges {
		originalEdge := original.Edges[i]
		if LengthRatio(originalEdge, newEdge) > float64(ε)/million {
			return false
		}
	}

	return true
}

func LengthRatio(oldEdge, newEdge *Edge) float64 {
	oldSqLength := oldEdge.SqLength()
	if oldSqLength == 0 {
		return million
	}

	return math.Abs(newEdge.SqLength()/oldSqLength - 1)
}
