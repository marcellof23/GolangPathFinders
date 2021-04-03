package main

import (
	"strconv"
)

type Graph struct {
	Nodes      map[int]*Vertex
	TotalNodes int
}

func CreateGraph() *Graph {
	return &Graph{
		Nodes: map[int]*Vertex{},
	}
}

type Vertex struct {
	Key   int
	Nodes map[int]*Vertex
}

func CreateVertex(key int) *Vertex {
	return &Vertex{
		Key:   key,
		Nodes: map[int]*Vertex{},
	}
}

func (vertex *Vertex) String() string {
	str := strconv.Itoa(vertex.Key) + ":"
	for _, adjacent := range vertex.Nodes {
		str = str + " " + strconv.Itoa(adjacent.Key)
	}
	return str
}

func (Graphs *Graph) AddVertex(v int) {
	vertex := CreateVertex(v)
	Graphs.Nodes[v] = vertex
}

func (Graphs *Graph) AddEdge(v1, v2 int) {
	u := Graphs.Nodes[v1]
	v := Graphs.Nodes[v2]

	u.Nodes[v.Key] = v
	v.Nodes[u.Key] = u

	Graphs.Nodes[u.Key] = u
	Graphs.Nodes[v.Key] = v
}

func (Graphs *Graph) String() string {
	str := ""
	i := 0
	for _, vertex := range Graphs.Nodes {
		if i > 0 {
			str = str + "\n"
		}
		i++
		str += vertex.String()
	}
	return str
}
