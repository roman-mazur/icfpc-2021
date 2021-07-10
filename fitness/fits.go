package fitness

import "github.com/roman-mazur/icfpc-2021/data"

// Fit returns true if the data.Figure fits in the given data.Hole.
// Could be optimized (for instance, checking if any point of the data.Figure vertices are in the data.Hole, and check if no
// data.Figure's edges intersect with data.Hole edges.)
func Fit(f data.Figure, h data.Hole) bool {
	for _, v := range f.Vertices {
		if !h.Contains(v) {
			return false
		}
	}

	for _, e := range f.Edges {
		for _, he := range h.Edges {
			if e.Intersect(*he) {
				return false
			}
		}
	}

	return true
}

// Unfit describes an unfit occurrence, pointing to the unfitting Edge, and anfitting vertices (1 or 2)
type Unfit struct {
	Edge   *data.Edge
	Unfits []*data.Vertex
}

// ListUnfits returns a list of unfitting Unfit occurrences from the figure, inside the Hole.
// Output array length is null if figure fits in the Hole.
func ListUnfits(f data.Figure, h data.Hole) (list []Unfit) {
	list = make([]Unfit, 0, len(f.Edges))

	for _, e := range f.Edges {
		aFits := h.Contains(*e.A)
		bFits := h.Contains(*e.B)

		if !aFits || !bFits {
			unfits := make([]*data.Vertex, 0, 2)
			if !aFits {
				unfits = append(unfits, e.A)
			}
			if !bFits {
				unfits = append(unfits, e.B)
			}

			list = append(list, Unfit{e, unfits})
		}
	}

	return
}
