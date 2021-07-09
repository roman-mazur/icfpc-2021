package transform

import (
	"github.com/roman-mazur/icfpc-2021/data"
)

type FoldDirection byte
const (
	FoldRight FoldDirection = iota
	FoldLeft
)

func Fold(figure data.Figure, edge *data.Edge, dir FoldDirection) data.Figure {
	// Line: y=k*x+b
	// TODO
	return figure
}

func Rotate(figure data.Figure, edge *data.Edge, angle float64) data.Figure {
	// TODO
	return figure
}
