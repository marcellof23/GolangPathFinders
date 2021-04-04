package Astar

import (
	pqheap "container/heap"
)

type Point struct {
	x, y float64
}

type Node interface {
	String() string
}
type Vertices []Node // vertices -> array of Node

type Graph interface {
	adj(n Node) []Node
}

// inisialisasi vertice baru
func CreateVertices(source Node) Vertices {
	return []Node{source}
}

type PairCost func(a, b Node) float64

func (v Vertices) TotalCost(d PairCost) (totalCost float64) {
	for i := 1; i < len(v); i++ {
		totalCost = totalCost + d(v[i-1], v[i])
	}
	return totalCost
}

func (v Vertices) addVertices(n Node) Vertices {
	newVertices := make([]Node, len(v), len(v)+1)
	copy(newVertices, v)
	newVertices = append(newVertices, n)
	return newVertices
}

// func (n Node) String() string {
//  str := strconv.Itoa(vertex.Key) + ":"
//  for _, adjacent := range vertex.Nodes {
//      str = str + " " + strconv.Itoa(adjacent.Key)
//  }
//  return str
// }

func astar(Graf Graph, start, dest Node, d, h PairCost) Vertices {
	closed_list := make(map[Node]bool)

	PQ := &priorityQueue{}
	pqheap.Init(PQ)
	pqheap.Push(PQ, &list_PQ{val: CreateVertices(start)})

	for PQ.Len() > 0 {
		p := pqheap.Pop(PQ).(*list_PQ).val.(Vertices)
		n := p[len(p)-1]
		if n == dest {
			return p
		}
		if closed_list[n] {
			continue
		}

		closed_list[n] = true

		for _, adjacent := range Graf.adj(n) {
			newVertices := p.addVertices(adjacent)
			pqheap.Push(PQ, &list_PQ{
				val:  newVertices,
				prio: -1 * (newVertices.TotalCost(d) + h(adjacent, dest)),
			})
		}
	}

	return nil
}
