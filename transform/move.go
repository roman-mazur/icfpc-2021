package transform

import "github.com/roman-mazur/icfpc-2021/data"

func Move(figure *data.Figure, vector data.Vertex) {
	for i := range figure.Vertices {
		figure.Vertices[i].X += vector.X
		figure.Vertices[i].Y += vector.Y
		figure.Vertices[i].Metadata.Reset()
	}
}
