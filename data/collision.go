package data

func (v Vertex) Equal(with Vertex) bool {
	return v.X == with.X && v.Y == with.Y
}

// Intersect returns true if given Edge intersects with w.
// It implements a linear equation resolution algorithm.
// More details: https://izziswift.com/how-can-i-check-if-two-segments-intersect/
func (e Edge) Intersect(w Edge) bool {
	dx0 := e.B.X - e.A.X
	dx1 := w.B.X - w.A.X
	dy0 := e.B.Y - e.A.Y
	dy1 := w.B.Y - w.A.Y
	p0 := dy1*(w.B.X-e.A.X) - dx1*(w.B.Y-e.A.Y)
	p1 := dy1*(w.B.X-e.B.X) - dx1*(w.B.Y-e.B.Y)
	p2 := dy0*(e.B.X-w.A.X) - dx0*(e.B.Y-w.A.Y)
	p3 := dy0*(e.B.X-w.B.X) - dx0*(e.B.Y-w.B.Y)
	return (p0*p1 <= 0) && (p2*p3 <= 0)
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
