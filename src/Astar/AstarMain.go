package Astar

// go test -v ./...

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/marcellof23/GolangPathFinders/src/Models"
)

var distance string

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
	return fmt.Sprintf("%d", p.Idx())
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

// ini kalo mau pake euclidian distance
func heuristicFunction(x, y Node) float64 {
	u := x.(Points)
	v := y.(Points)
	d_x := v.Lat() - u.Lat()
	d_y := v.Lng() - u.Lng()
	return 100 * math.Sqrt(float64(d_x*d_x+d_y*d_y))
}

func PrintAstar() {
	N := 8
	arrNodes := make([]Node, N)
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
	p, Dist := astar(g, arrNodes[3], arrNodes[5], HeuristicHaversine, HeuristicHaversine)
	fmt.Print("Path : ")
	for _, n := range p {
		fmt.Printf("%s ", n)
	}
	fmt.Printf("%f ", Dist)
	fmt.Println("")
}

// func StringAstar() string {

// 	var arrNodes [4]Node
// 	arrNodes[0] = NewPoints(2.14, 3.15, 0)
// 	arrNodes[1] = NewPoints(1.0, 7.0, 1)
// 	arrNodes[2] = NewPoints(1.0, 6.0, 2)
// 	arrNodes[3] = NewPoints(5.0, 6.0, 3)
// 	g := CreateGraph().AddEdge(arrNodes[0], arrNodes[1]).AddEdge(arrNodes[1], arrNodes[2]).AddEdge(arrNodes[0], arrNodes[3])
// 	p, Dist := astar(g, arrNodes[2], arrNodes[3], HeuristicHaversine, HeuristicHaversine)
// 	fmt.Print("Path : ")
// 	str := ""
// 	for _, n := range p {
// 		str += n.String()
// 	}
// 	return str
// }

func StringAstars(astardata Models.AstarData) (string, []float64) {
	var BobotEdge []float64
	arrNodes := make([]Node, astardata.NumNodes)
	stringSlice := strings.Split(astardata.Coordinates, ",")
	for i, u := range stringSlice {
		v := strings.Fields(u)
		p := make([]float64, 2)
		for j, str := range v {
			var errs error
			p[j], errs = strconv.ParseFloat(str, 64)
			if errs != nil {

			}
		}
		arrNodes[i] = NewPoints(p[0], p[1], i)
	}

	g := CreateGraph()
	strMatrix := strings.Split(astardata.MatrixRelation, ",")
	cnt := 0
	for i := 0; i < len(strMatrix); i++ {
		strMatrixIn := strings.Fields(strMatrix[i])
		for j := 0; j < len(strMatrix); j++ {
			if strMatrixIn[j] == "1" {
				cnt++
				g.AddEdge(arrNodes[i], arrNodes[j])
			}
		}
	}
	x, Dist := astar(g, arrNodes[astardata.SourceNode], arrNodes[astardata.DestNode], HeuristicHaversine, HeuristicHaversine)
	fmt.Print("Path : ")
	str := ""
	for i, n := range x {
		if i == 0 {
		} else if i == (len(x) - 1) {
			str += n.String()
			continue
		}
		u, err1 := strconv.Atoi(x[i].String())
		v, err2 := strconv.Atoi(x[i+1].String())
		BobotEdge = append(BobotEdge, HeuristicHaversine(arrNodes[u], arrNodes[v]))
		if err1 != nil && err2 != nil {

		}
		str += n.String() + ","
	}
	dist := Dist * -1
	distance = strconv.FormatFloat(dist, 'f', 6, 64)
	return str, BobotEdge
}

func HeuristicHaversineJson() string {
	return distance
}

func EdgeJson(BobotEdge []float64) string {
	str := ""
	for i, _ := range BobotEdge {
		if i == 0 {
			str += "("
		}
		if i == (len(BobotEdge) - 1) {
			str += strconv.FormatFloat(BobotEdge[i], 'f', 6, 64) + ")"
			continue
		}
		str += strconv.FormatFloat(BobotEdge[i], 'f', 6, 64) + ","

	}
	return str
}
