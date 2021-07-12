package search

import "github.com/roman-mazur/icfpc-2021/data"

type sitem struct {
	currV, parentV *data.Vertex
	currF          *data.Figure
	newParentV     *data.Vertex
}

type stack struct {
	items []*sitem
	size int

	visited map[*data.Vertex]struct{}
}

func (s *stack) push(v, p *data.Vertex, f *data.Figure, oldV *data.Vertex) {
	if s.visited == nil {
		s.visited = make(map[*data.Vertex]struct{})
	}
	if _, visited := s.visited[v]; visited {
		return
	}
	s.visited[v] = struct{}{}

	s.size++
	item := &sitem{currF: f, currV: v, newParentV: oldV, parentV: p}
	if s.size > len(s.items) {
		s.items = append(s.items, item)
	} else {
		s.items[s.size-1] = item
	}
}

func (s *stack) pop() (*data.Vertex, *data.Figure, *data.Vertex, *data.Vertex) {
	if s.size == 0 {
		panic("empty stack")
	}
	s.size--
	res := s.items[s.size]
	s.items[s.size] = nil
	return res.currV, res.currF, res.parentV, res.newParentV
}

func (s *stack) empty() bool {
	return s.size == 0
}

