package search

import (
	"fmt"
	"math"

	"github.com/roman-mazur/icfpc-2021/data"
	"github.com/roman-mazur/icfpc-2021/fitness"
)

const debug = false

func log(iter int, format string, args... interface{}) {
	if debug {
		for i := 0; i < iter; i++ {
			fmt.Print("  ")
		}
		fmt.Printf(format + "\n", args...)
	}
}

func Solve(f *data.Figure, hole *data.Hole, v *data.Vertex, eps int) *data.Figure {
	var s stack
	s.push(v, nil, f, nil)

	last := f

	for !s.empty() {
		v, f, p, newP := s.pop()
		log(s.size,"searching for %s", v)
		for _, candidate := range selectTargets(f, hole, v, p, newP, eps, s.size) {
			nextF := f.Copy()

			nextV := nextF.FindV(*v)
			if nextV == nil {
				continue
			}
			nextV.X = candidate.X
			nextV.Y = candidate.Y

			othersAffected := false
			for _, connected := range nextF.GetConnectedVertices(nextV) {
				oldEdge := &data.Edge{A: v, B: connected}
				newEdge := &data.Edge{A: nextV, B: connected}
				if !data.GoodRatio(oldEdge, newEdge, eps) {
					s.push(connected, v, &nextF, nextV)
					othersAffected = true
				}
			}
			if !othersAffected {
				log(s.size, "we have something with %s: %s", nextV, nextF)
				return &nextF
			}
			last = &nextF
		}
	}
	return last
}

func Solution(figure *data.Figure, hole *data.Hole, eps int) *data.Figure {
	original := figure.Copy()
	res := original

	unfits := fitness.ListUnfits(res, *hole)
	for _, u := range unfits {
		for _, v := range u.Vertices {
			 s := Solve(&res, hole, v, eps)
			 if s != nil && s.IsValid(original, eps) {
			 	res = *s
			 }
		}
	}
	if len(res.Edges) == 0 {
		panic("Wrong solution")
	}
	return &res
}

func selectTargets(f *data.Figure, hole *data.Hole, v, p, newP *data.Vertex, eps, iter int) []data.Vertex {
	connections := f.GetConnectedVertices(v)
	var res []data.Vertex

	// First, consider hole vertices.
	for _, hv := range hole.Vertices {
		good := true
		for _, c := range connections {
			k := data.LengthRatioRelative(&data.Edge{A: v, B: c}, &data.Edge{A: &hv, B: c}, eps)
			//log(iter,"%s / %s => %f", hv, c, k)
			if k > 4 {
				good = false
				break
			}
		}
		if good {
			//log(iter,"Consider %s as a target for %s", hv, v)
			res = append(res, hv)
		}
	}

	// Use p (previous vertex) to find an area that will satisfy the constraints.
	if newP != nil && p != nil {
		res = append(res, generateAreaCandidates(v, p, newP)...)
	}

	return res
}

// generateAreaCandidates find potential vertices around new parent trying to keep the same
func generateAreaCandidates(v, p, newP *data.Vertex) []data.Vertex {
	res := make(map[data.Vertex]struct{})

	radius := math.Sqrt(data.Edge{A: p, B: v}.SqLength())
	step := math.Pi / 8
	for a := 0.0; a < 2 * math.Pi; a += step {
		point := data.Vertex{X: newP.X + math.Cos(a)*radius, Y: newP.Y + math.Sin(a)*radius}
		for i := -1.0; i <= 1.0; i++ {
			for j := -1.0; j <= 1.0; j++ {
				res[data.Vertex{
					X:        math.Round(point.X + i),
					Y:        math.Round(point.Y + j),
				}] = struct{}{}
			}
		}
	}

	list := make([]data.Vertex, len(res))
	i := 0
	for v := range res {
		list[i] = v
		i++
	}
	return list
}
