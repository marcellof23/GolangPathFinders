package hihi

// pake ini buat test go test -v ./...
import (
	"fmt"
	"testing"
)

func initialize(N int, g *Graph) {
	for i := 1; i <= N; i++ {
		g.AddVertex(i)
	}
}
func TestDfs(t *testing.T) {
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
