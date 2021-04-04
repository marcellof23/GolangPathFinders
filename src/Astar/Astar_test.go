package Astar

// go test -v ./...

import (
	"fmt"
	"math"
	"testing"
)

type Points struct {
	lat float64
	lng float64
}

const (
	EARTH_RADIUS = 6371
)

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func NewPoints(lat float64, lng float64) Points {
	return Points{lat: lat, lng: lng}
}
func (p Points) Lat() float64 {
	return p.lat
}

func (p Points) Lng() float64 {
	return p.lng
}

func (p Points) String() string {
	return fmt.Sprintf("(%f, %f)", p.Lat(), p.Lng())
}

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

func HeuristicHaversine(x, y Node) (km float64) {
	u := x.(Points)
	v := y.(Points)
	lat1 := degreesToRadians(u.Lat())
	lng1 := degreesToRadians(u.Lng())
	lat2 := degreesToRadians(v.Lat())
	lng2 := degreesToRadians(v.Lng())

	distLat := lat2 - lat1
	distLon := lng2 - lng1

	a := math.Pow(math.Sin(distLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(distLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	km = c * EARTH_RADIUS

	return km
}

func heuristicFunction(x, y Node) float64 {
	u := x.(Points)
	v := y.(Points)
	d_x := v.Lat() - u.Lat()
	d_y := v.Lng() - u.Lng()
	return 100 * math.Sqrt(float64(d_x*d_x+d_y*d_y))
}
func TestAstar(t *testing.T) {

	var arrNodes [4]Node
	arrNodes[0] = NewPoints(2.14, 3.15)
	arrNodes[1] = NewPoints(1.0, 7.0)
	arrNodes[2] = NewPoints(1.0, 6.0)
	arrNodes[3] = NewPoints(5.0, 6.0)
	g := CreateGraph().AddEdge(arrNodes[0], arrNodes[1]).AddEdge(arrNodes[1], arrNodes[2]).AddEdge(arrNodes[0], arrNodes[3])
	p := astar(g, arrNodes[2], arrNodes[3], HeuristicHaversine, HeuristicHaversine)
	fmt.Print("Path : ")
	for _, n := range p {
		fmt.Printf("%s ", n)
	}
	fmt.Println("")
}
