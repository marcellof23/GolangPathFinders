package algorithm

import (
	"fmt"
	"strconv"
)

func DFS(G *Graph, source *Vertex, path func(int), vis []bool) {

	if source == nil {
		return
	}
	vis[source.Key] = true
	path(source.Key)

	for _, v := range source.Nodes {
		if !vis[v.Key] {
			DFS(G, v, path, vis)
		}

	}
}

func initialize(N int, g *Graph) {
	for i := 1; i <= N; i++ {
		g.AddVertex(i)
	}
}
func TestDfs() {
	N := 10
	g := CreateGraph()

	initialize(N, g)

	g.AddEdge(1, 9)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(3, 4)
	g.AddEdge(5, 6)
	g.AddEdge(5, 8)
	g.AddEdge(6, 7)
	g.AddEdge(9, 10)
	g.AddEdge(2, 3)

	visit := []int{}
	path := func(i int) {
		visit = append(visit, i)
	}
	vis := make([]bool, N+1)
	DFS(g, g.Nodes[5], path, vis)

	fmt.Println(visit)
}

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
