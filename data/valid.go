package data

import "math"

const million = float64(1e6)

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

func GoodRatio(oldEdge, newEdge *Edge, ε int) bool {
	return LengthRatio(oldEdge, newEdge) <= float64(ε)/million
}

func ScaleToPreserveRatio(oldEdge, newEdge *Edge, ε int) float64 {
	l1 := oldEdge.SqLength()
	lr := newEdge.SqLength() / l1
	if lr == 1 {
		return 1
	}

	allowed := float64(ε)/million
	target := 1 + allowed
	if lr < 1 {
		target = 1 - allowed
	}

	targetLen := math.Sqrt(l1 * target)
	return targetLen / newEdge.PLine().Len()
}
