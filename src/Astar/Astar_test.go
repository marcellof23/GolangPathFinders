package Astar

import (
	"fmt"
	"image"
	"math"
	"testing"
)

const (
	earthRaidusKm = 6371 // radius of the earth in kilometers.
)

type Graphs map[Node][]Node

func CreateGraph() Graphs {
	return make(map[Node][]Node)
}

func (Graf Graphs) adj(NumNodes Node) []Node {
	return Graf[NumNodes]
}

func (Graf Graphs) AddEdge(a, b Node) Graphs {
	Graf[a] = append(Graf[a], b)
	Graf[b] = append(Graf[b], a)
	return Graf
}

func heuristicFunction(x, y Node) float64 {
	u := x.(image.Point)
	v := y.(image.Point)
	d_x := v.X - u.X
	d_y := v.Y - u.Y
	return 100 * math.Sqrt(float64(d_x*d_x+d_y*d_y))
}
func TestAstar(t *testing.T) {

	var arrNodes [4]Node
	arrNodes[0] = image.Pt(2, 3)
	arrNodes[1] = image.Pt(1, 7)
	arrNodes[2] = image.Pt(1, 6)
	arrNodes[3] = image.Pt(5, 6)
	g := CreateGraph().AddEdge(arrNodes[0], arrNodes[1]).AddEdge(arrNodes[1], arrNodes[2]).AddEdge(arrNodes[0], arrNodes[3])
	p := astar(g, arrNodes[2], arrNodes[3], heuristicFunction, heuristicFunction)
	fmt.Print("Path : ")
	for _, n := range p {
		fmt.Printf("%s ", n)
	}
	fmt.Println("")
}
