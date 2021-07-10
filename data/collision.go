package data

func (v Vertex) Equal(with Vertex) bool {
	return v.X == with.X && v.Y == with.Y
}

// Intersect returns true if with intersects with the given Edge.
// It implements a linear equation resolution algorithm.
// More details: https://stackoverflow.com/questions/4977491/determining-if-two-line-segments-intersect/4977569#4977569
func (e Edge) Intersect(with Edge) bool {
	// Converts edge (origin, destination Vertex) into vectors (origin, vector Vertex)
	vecE := Vertex{e.B.X - e.A.X, e.B.Y - e.A.Y}
	vecW := Vertex{with.B.X - with.A.X, with.B.Y - with.A.Y}

	// Check if first vertices collides
	if e.A.Equal(*with.A) {
		return true
	}

	// Compute matrix determinent
	det := float64(vecW.X*vecE.Y - vecE.X*vecW.Y)
	if det == 0 {
		// Lines are parallel, no possible collision
		return false
	}

	// Compute lines intersection
	s := 1 / det * float64((e.A.X-with.A.X)*vecE.Y-(e.A.Y-with.A.Y)*vecE.X)
	t := 1 / det * float64(((0 - (e.A.X-with.A.X)*vecW.Y) - (e.A.Y-with.A.Y)*vecW.X))

	// Check if collision happens inside segments bounds
	if t >= 0 &&
		t <= 1 &&
		s >= 0 &&
		s <= 1 {
		return true
	}

	return false
}

// Contains returns true if v is contained in the given Hole.
// It computes the number of collisions with hole's edges, left and right from the given vertex. If it's odd, the
// vertex is inside. If it's even, the vertex is outside.
// It is inspired by the concave polygon collision described here: http://www.alienryderflex.com/polygon/
func (h Hole) Contains(v Vertex) bool {
	// MAX_BOUND defines an approximatively infinite ray coming from the vertex
	const MAX_BOUND = 1000000000

	leftLine := Edge{
		A: &Vertex{X: 0 - MAX_BOUND, Y: v.Y},
		B: &v,
	}
	rightLine := Edge{
		A: &v,
		B: &Vertex{X: MAX_BOUND, Y: v.Y},
	}

	leftCollisionCount := 0
	rightCollisionCount := 0
	for _, e := range h.Edges {
		if leftLine.Intersect(*e) {
			leftCollisionCount++
		}
		if rightLine.Intersect(*e) {
			rightCollisionCount++
		}
	}

	// If one of the collision count is odd, the vertex belongs to the Hole.
	// (Remember the vertex can belong to a Hole segment, in which case left and right counts will not be both odds.)
	if leftCollisionCount > 0 && rightCollisionCount > 0 && // At least one collision from both sides is required
		(leftCollisionCount%2 == 1 || rightCollisionCount%2 == 1) {
		return true
	}

	return false
}
