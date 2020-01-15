package main

 import (
         /*"fmt"
         "golang.org/x/exp/rand"
				 "os"
				 "strconv"
				 "gonum.org/v1/gonum/stat/distuv"*/
         //"priorityQueue"
         "container/heap"
 )
 type flow struct {
     //name string
     id  int
		 startTime float64   // Possionien(lambda)
		 sfcID int
		 durationFlow float64  //expo
		 roughputFlow float64  //random-uniform
     index  int
 }
 const(
	 nbMaxFlow=100
)
//*********** Priority Queue of Flows based on the heap of golang
// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*flow

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority (i.e. priority=start time) so we use smaller than here.
	return pq[i].startTime < pq[j].startTime
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*flow)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *flow, 		 startTime float64, sfcID int, durationFlow float64, roughputFlow float64) {
  item.startTime=startTime
  item.sfcID=sfcID
  item.durationFlow=durationFlow
  item.roughputFlow=roughputFlow
	heap.Fix(pq, item.index)
}
//*********** END For Priority Queue codes
