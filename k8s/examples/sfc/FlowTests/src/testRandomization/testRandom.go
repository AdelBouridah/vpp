package main

 import (
         "fmt"
         "golang.org/x/exp/rand"
				 "os"
				 "strconv"
				 "gonum.org/v1/gonum/stat/distuv"
 )
 type flow struct {
     //name string
     id  int
		 startTime float64   // Possionien(lambda)
		 sfcID int
		 durationFlow float64  //expo
		 roughputFlow float64  //random-uniform
 }
 const(
	 nbMaxFlow=100
)
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
				 }
				 fmt.Println("les oneFlow sont")
				 for i:=0;i<100;i++{
					 	fmt.Println("oneFlow",i, "  est %v", flowArray[i].sfcID)
				 }
 }
