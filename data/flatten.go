package data

import "math"

func (f Figure) FlattenToGrid() Figure {
	f = f.Copy()

	for i, v := range f.Vertices {
		f.Vertices[i].X = math.Round(v.X)
		f.Vertices[i].Y = math.Round(v.Y)
		f.Vertices[i].Metadata.Reset()
	}

	return f
}
