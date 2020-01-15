package main

 import (
         "fmt"
         "golang.org/x/exp/rand"
				 "os"
				 "strconv"
				 "gonum.org/v1/gonum/stat/distuv"
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
 func main() {

	 			 flowArray:= make([]flow, 0)
				 //var oneFlow flow
	 			 //var t[1000] float64
         lambda, _:= strconv.ParseFloat(os.Args[1], 64)
				 lambda2, _:= strconv.ParseFloat(os.Args[2], 64)
				 sfcNumber, _ := strconv.Atoi(os.Args[3])
				 //fmt.Println("lambda",lambda)
         r := rand.New(rand.NewSource(100))
				 rDurationFlow:= rand.New(rand.NewSource(20))
				 rSFC:=rand.New(rand.NewSource(20))
				 rRoughputFlow:=distuv.Uniform{
						Min:10.0,
						Max:20.0,
						Src: rand.NewSource(100) ,
				 }
				 startTime:=0.0
				 //startTime:=r.ExpFloat64()/lambda
         // Put all the Flows in stratTime priority Queue
         pqFlows := make(PriorityQueue, 0)
				 for i:=0;i<100;i++{
					  startTime=startTime+(r.ExpFloat64()/lambda)
						duration:= rDurationFlow.ExpFloat64()/lambda2
						sfc:= rSFC.Intn(sfcNumber)
						roughtput:=rRoughputFlow.Rand()
					  oneFlow:=flow{
							 id: i,
				 			 startTime: startTime,   // Possionien(lambda)
				 			 sfcID: sfc,
				 			 durationFlow: duration,  //expo
				 			 roughputFlow: roughtput,  //random-uniform
						}
						//flowArray[i]=oneFlow
						flowArray=append(flowArray, oneFlow)
            heap.Push(&pqFlows, &oneFlow)
				 }
				 /*fmt.Println("les oneFlow sont")
				 for i:=0;i<100;i++{
					 	fmt.Println("oneFlow",i, "  est %v", flowArray[i].sfcID)
				 }*/
         ///////// test priorityQueue
         fmt.Println(" ***************Test priority queue**********")
         // Take the items out; they arrive in decreasing priority order.
         for pqFlows.Len() > 0 {
           item := heap.Pop(&pqFlows).(*flow)
           fmt.Println("\nFlow- \n", item.id, item.startTime,  item.durationFlow, item.roughputFlow)
         }
         /*
         for i:=0;i<100;i++{
            fmt.Println("oneFlow",i, "  est %v", flowArray[i].sfcID)
         }*/
 }
