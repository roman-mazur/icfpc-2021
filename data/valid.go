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

		ratio := math.Abs(newEdge.SqLength()/originalEdge.SqLength() - 1)

		if ratio > float64(epsilon)/float64(million) {
			return false
		}
	}

	return true
}
