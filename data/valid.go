package data

import "math"

const million = float64(1000000)

// IsValid tells if the figure f is a valid figure regarding the original figure and epsilon
func (f Figure) IsValid(original Figure, epsilon int) bool {
	if len(f.Edges) != len(original.Edges) ||
		len(f.Vertices) != len(original.Vertices) {
		return false
	}

	for i, newEdge := range f.Edges {
		originalEdge := original.Edges[i]
		if !GoodRatio(originalEdge, newEdge, epsilon) {
			return false
		}
	}

	return true
}

func GoodRatio(oldEdge, newEdge *Edge, epsilon int) bool {
	ratio := math.Abs(newEdge.SqLength()/oldEdge.SqLength() - 1)
	return ratio <= float64(epsilon)/million
}
