package Astar

// go test -v ./...

import (
	"fmt"
	"math"
)

type Points struct {
	lat float64
	lng float64
	idx int
}

const (
	EARTH_RADIUS = 6371
)

func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func NewPoints(lat float64, lng float64, idx int) Points {
	return Points{lat: lat, lng: lng, idx: idx}
}
func (p Points) Lat() float64 {
	return p.lat
}

func (p Points) Lng() float64 {
	return p.lng
}
func (p Points) Idx() int {
	return p.idx
}

func (p Points) String() string {
	return fmt.Sprintf("(%f, %f, %d)", p.Lat(), p.Lng(), p.Idx())
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

// func TestAstar(t *testing.T) {

// 	var arrNodes [4]Node
// 	arrNodes[0] = NewPoints(2.14, 3.15)
// 	arrNodes[1] = NewPoints(1.0, 7.0)
// 	arrNodes[2] = NewPoints(1.0, 6.0)
// 	arrNodes[3] = NewPoints(5.0, 6.0)
// 	g := CreateGraph().AddEdge(arrNodes[0], arrNodes[1]).AddEdge(arrNodes[1], arrNodes[2]).AddEdge(arrNodes[0], arrNodes[3])
// 	p := astar(g, arrNodes[2], arrNodes[3], HeuristicHaversine, HeuristicHaversine)
// 	fmt.Print("Path : ")
// 	for _, n := range p {
// 		fmt.Printf("%s ", n)
// 	}
// 	fmt.Println("")
// }

func PrintAstar() {

	var arrNodes [8]Node
	arrNodes[0] = NewPoints(-6.892616, 107.610423, 0)
	arrNodes[1] = NewPoints(-6.890464174304161, 107.60724717865062, 1)
	arrNodes[2] = NewPoints(-6.887407, 107.611507, 2)
	arrNodes[3] = NewPoints(-6.886107754151466, 107.60814840083982, 3)
	arrNodes[4] = NewPoints(-6.886011891139484, 107.61040145641657, 4)
	arrNodes[5] = NewPoints(-6.884371565786224, 107.61305147865777, 5)
	arrNodes[6] = NewPoints(-6.893819336137485, 107.61446768497895, 6)
	arrNodes[7] = NewPoints(-6.890591990423812, 107.60992938762139, 7)
	g := CreateGraph()
	g.AddEdge(arrNodes[0], arrNodes[1])
	g.AddEdge(arrNodes[0], arrNodes[6])
	g.AddEdge(arrNodes[0], arrNodes[7])
	g.AddEdge(arrNodes[1], arrNodes[2])
	g.AddEdge(arrNodes[1], arrNodes[3])
	g.AddEdge(arrNodes[2], arrNodes[3])
	g.AddEdge(arrNodes[2], arrNodes[4])
	g.AddEdge(arrNodes[2], arrNodes[7])
	g.AddEdge(arrNodes[3], arrNodes[4])
	g.AddEdge(arrNodes[4], arrNodes[5])
	g.AddEdge(arrNodes[5], arrNodes[6])
	p := astar(g, arrNodes[3], arrNodes[5], HeuristicHaversine, HeuristicHaversine)
	fmt.Print("Path : ")
	for _, n := range p {
		fmt.Printf("%s ", n)
	}
	fmt.Println("")
}

func StringAstar() string {

	var arrNodes [4]Node
	arrNodes[0] = NewPoints(2.14, 3.15, 0)
	arrNodes[1] = NewPoints(1.0, 7.0, 1)
	arrNodes[2] = NewPoints(1.0, 6.0, 2)
	arrNodes[3] = NewPoints(5.0, 6.0, 3)
	g := CreateGraph().AddEdge(arrNodes[0], arrNodes[1]).AddEdge(arrNodes[1], arrNodes[2]).AddEdge(arrNodes[0], arrNodes[3])
	p := astar(g, arrNodes[2], arrNodes[3], HeuristicHaversine, HeuristicHaversine)
	fmt.Print("Path : ")
	str := ""
	for _, n := range p {
		str += n.String()
	}
	return str
}
