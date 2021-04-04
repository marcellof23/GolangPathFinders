package Astar

type list_PQ struct {
	val interface {
	}
	prio float64
}

// ini priority queue buat custom priority

type priorityQueue []*list_PQ

func (PQ priorityQueue) Less(i, j int) bool {
	if PQ[i].prio <= PQ[j].prio {
		return false
	} else {
		return true
	}
}

func (PQ priorityQueue) Swap(i, j int) {
	PQ[i], PQ[j] = PQ[j], PQ[i]
}

func (PQ *priorityQueue) Push(items interface{}) {
	*PQ = append(*PQ, items.(*list_PQ))
}

func (PQ *priorityQueue) Pop() interface{} {
	PQ2 := *PQ
	n := len(PQ2)
	item := PQ2[n-1]
	*PQ = PQ2[0 : n-1]
	return item
}

func (PQ priorityQueue) Len() int {
	return len(PQ)
}
