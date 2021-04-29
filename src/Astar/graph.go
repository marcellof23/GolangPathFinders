package Astar

import (
	"strconv"
)

type Graf struct {
	NodeX      map[int]*Vertex
	TotalNodes int
}

func CreateGraf() *Graf {
	return &Graf{
		NodeX: map[int]*Vertex{},
	}
}

type Vertex struct {
	Key   int
	NodeX map[int]*Vertex
}

func CreateVertex(key int) *Vertex {
	return &Vertex{
		Key:   key,
		NodeX: map[int]*Vertex{},
	}
}

func (vertex *Vertex) String() string {
	str := strconv.Itoa(vertex.Key) + ":"
	for _, adjacent := range vertex.NodeX {
		str = str + " " + strconv.Itoa(adjacent.Key)
	}
	return str
}

func (Graphs *Graf) AddVertex(v int) {
	vertex := CreateVertex(v)
	Graphs.NodeX[v] = vertex
}

func (Graphs *Graf) AddEdge(v1, v2 int) {
	u := Graphs.NodeX[v1]
	v := Graphs.NodeX[v2]

	u.NodeX[v.Key] = v
	v.NodeX[u.Key] = u

	Graphs.NodeX[u.Key] = u
	Graphs.NodeX[v.Key] = v
}

func (Graphs *Graf) String() string {
	str := ""
	i := 0
	for _, vertex := range Graphs.NodeX {
		if i > 0 {
			str = str + "\n"
		}
		i++
		str += vertex.String()
	}
	return str
}
