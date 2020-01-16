package main

import (
	"fmt"
	"ketama"
	"strconv"
	"os/exec"
	"strings"
	"log"
	"os"
	//"io"
	"sync"
	//"bytes"
	"time"
	"container/heap"
	"gonum.org/v1/gonum/stat/distuv"
	"golang.org/x/exp/rand"
	"math"
	"encoding/csv"
)
func iperfinstance(cnf string, port string, i int, wg *sync.WaitGroup, flowTime float64, flowRoughput float64, sfcID int){
	fmt.Println(" ************************************************************************************************")
	fmt.Println(" ******		Test IPERF Flux ID="+strconv.Itoa(i)+" Path "+cnf+" Port "+port+"           **********")
	fmt.Println(" *************************************************************************************************")
	//
	ftime:= int(math.RoundToEven(flowTime))
	if (ftime==0){
		ftime=1
	}
	froughput:=int(math.RoundToEven(flowRoughput))
	cmd := exec.Command("kubectl", "exec", cnf , "--","iperf", "-c", "192.168.187.2", "-p", port, "-y", "C", "-u","-b", strconv.Itoa(froughput)+"m" , "-t", strconv.Itoa(ftime) ) // CSV Output  "-y", "C"
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("********************** iperf error ********************")
		log.Fatal(err)
	}
	fmt.Println(" "+string(out)+"test iperf ")
	defer wg.Done()
	f, _ := os.Create("tmp/file-flow"+strconv.Itoa(i)+"Path-"+cnf+"-port-"+port+".txt")
	defer f.Close()
	//fmt.Println("test iperf ********\n"+string(out)+"\n ***")
	f.WriteString("FlowID-"+strconv.Itoa(i)+","+"SFC-ID"+strconv.Itoa(sfcID)+","+"Path-id-"+cnf+","+string(out))
	f.Sync()
}

func main() {
	/*
	fmt.Println(" *******************************************************************************************")
	fmt.Println(" ******			Programm inputs are in this order                                    **********")
	fmt.Println(" ***********Flows Number, server Ports limit, lambda1, lambda2, SFC Number******************")
	fmt.Println(" *******************************************************************************************")
	*/
   // Before We call inputpaths here

	// Here Ketam Ring (before)

	fmt.Println(" *************************************************************************")
	fmt.Println(" ******			TEST IPERF - BEGIN 		                             **********")
	fmt.Println(" *************************************************************************")
	/*
	fmt.Println(" *******************************************************************************************")
	fmt.Println(" ******			Programm inputs are in this order                                    **********")
	fmt.Println(" ***********Flows Number, server Ports limit, lambda1, lambda2, SFC Number******************")
	fmt.Println(" *******************************************************************************************")
	*/
	nbrFlux, _ := strconv.Atoi(os.Args[1])
	limitPorts, _ := strconv.Atoi(os.Args[2])
	lambda, _:= strconv.ParseFloat(os.Args[3], 64)
	lambda2, _:= strconv.ParseFloat(os.Args[4], 64)
	sfcNumber, _ := strconv.Atoi(os.Args[5])

	fmt.Println(" *****************************************************************************")
	fmt.Println(" ****** BEGIN:	GENERATING FLOWS, Tables and priority queue for FLOWS *********")
	fmt.Println(" *****************************************************************************")

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
	// Put all the Flows in stratTime priority Queue & in inpuFile (trace)
	flowArray:= make([]flow, 0)
	pqFlows := make(PriorityQueue, 0)


	for i:=0;i<nbrFlux;i++{
		 startTime=startTime+(r.ExpFloat64()/lambda)
		 duration:= rDurationFlow.ExpFloat64()/lambda2
		 sfc:= rSFC.Intn(sfcNumber-1) + 1  //sfc num is between 1 and sfcNumber
		 roughtput:=rRoughputFlow.Rand()
		 oneFlow:=flow{
				id: i,
				startTime: startTime,   // Possionien(lambda)
				sfcID: sfc,						//
				durationFlow: duration,  //expo
				roughputFlow: roughtput,  //random-uniform
		 }
		 //flowArray[i]=oneFlow
		 flowArray=append(flowArray, oneFlow)
		 heap.Push(&pqFlows, &oneFlow)

	}
	///////////////////////////////////////////////////////////////////////////
	///               prepare file of inputs                                ///
	///////////////////////////////////////////////////////////////////////////
	//
	fileInputs, errFileInpouts := os.Create("inputParams.txt")
	checkError("Cannot create file", errFileInpouts)
	defer fileInputs.Close()
	lineFile:=make([]string, 5)
	for i:=0;i<nbrFlux;i++{
		// Write results in File of inputs fmt.Sprintf("%f",   )
		//lineFile[0]=strconv.Itoa(i)
		lineFile[0]="FlowID-"+strconv.Itoa(flowArray[i].id)
		lineFile[1]=strconv.Itoa(flowArray[i].sfcID)
		lineFile[2]=fmt.Sprintf("%f",flowArray[i].roughputFlow)
		lineFile[3]=fmt.Sprintf("%f",flowArray[i].durationFlow)
		lineFile[4]=fmt.Sprintf("%f",flowArray[i].startTime)
		writerFile := csv.NewWriter(fileInputs)
		writerFile.Comma=' '
		defer writerFile.Flush()
		errWrite:= writerFile.Write(lineFile)
		checkError("Cannot write to file", errWrite)
	}
	fmt.Println(" *****************************************************************************")
	fmt.Println(" ******	END: GENERATING FLOWS, Tables and priority queue for FLOWS     ******")
	fmt.Println(" *****************************************************************************")
	//fmt.Println(" Flux Numbers : "+strconv.Itoa(nbrFlux))
	var wg sync.WaitGroup
	//wg.Add(nbrFlux)
	fmt.Println(" *************************************************************************")
	fmt.Println(" ******			Create SFCs Rings for Ketama                       **********")
	fmt.Println(" ******			Each SFC Ring has its cnf inputs Paths             **********")
	fmt.Println(" *************************************************************************")
	// For each SFC
	allInputPaths:=make([][]string, sfcNumber)
	ringsSFCs:=make([] *ketama.Continuum, sfcNumber)
	//var ringsSFCs [] *ketama.Continuum
	for sfcID:=1; sfcID<=sfcNumber; sfcID++{
		allInputPaths[sfcID-1]= getInputCnfsSFC(sfcID)
		ringsElem, errRings:=createKetamaRingSFC(allInputPaths[sfcID-1])
		checkError("Cant create Ring", errRings)
		ringsSFCs[sfcID-1]= ringsElem
	}

	fmt.Println(" *************************************************************************")
	fmt.Println(" ******			End- creating SFCs Rings for Ketama                **********")
	fmt.Println(" *************************************************************************")




	mPorts := make(map[string]int)
	fmt.Println("Running for loop…")
	OldStartTime:=0.0
	for pqFlows.Len() > 0 {
		// Get the next flow item to be send
		item := heap.Pop(&pqFlows).(*flow)

		// Select a path (CNF Input) with ketam in the concerned Ring -Related to the SFC ID
		//s := k.Hash("iperf" + strconv.Itoa(item.id))
		s:= ringsSFCs[item.sfcID-1].Hash("iperf" + strconv.Itoa(item.id))
		mPorts[s]++  // use another port if the path has always iperf in progress


		if (mPorts[s]>limitPorts){  //Not usefull usecase, whereas added here to avoid crashing i.e. listening ports must be sufficient
			// try later to re-initialize mPorts[s]=0
			mPorts[s]=1
			fmt.Println("limit port depassed re-use from port 1 for path of CNf:",s)
			//defer wg.Done()
		}//else{

				var vport string
				if (mPorts[s]<=9){
					vport="500"+strconv.Itoa(mPorts[s])

				}else{
					vport="50"+strconv.Itoa(mPorts[s])
				}
				fmt.Println(" *****************************************************************************")
				fmt.Println(" ******	Excutethe iperf with the FLOWS                                 ******", vport)
				fmt.Println(" *****************************************************************************")
				fmt.Printf("\nStart time %v for FLOWID=:  %v \n", item.startTime, item.id)
				duration := time.Duration( (item.startTime-OldStartTime) * float64(time.Second))
				OldStartTime=item.startTime
				timer1 := time.NewTimer(duration)
				<-timer1.C
        //fmt.Println("Timer 1 fired", duration)
				wg.Add(1)
				go iperfinstance(s, vport, item.id, &wg, item.durationFlow, item.roughputFlow, item.sfcID)
		//}

		/*go iperfinstance(s, "500"+strconv.Itoa(mPorts[s]), i, &wg)
		//fmt.Println(i, s)
		for j := 2; j <= 3; j++ {
			//wg.Add(1)
			//time.Sleep(0,01 * time.Second)
			go iperfinstance(s, "500"+strconv.Itoa(j), i, &wg)
			//defer wg.Done()
		}*/

	}
	fmt.Println(" Before wait   ")
	wg.Wait()
	fmt.Println("Finished for loop")
	/*cmdFileRest := exec.Command("cat", "*", ">", "result.xls")
	_, errFileRST := cmdFileRest.CombinedOutput()
	if errFileRST != nil {
		log.Fatal(errFileRST)
	}*/
	fmt.Println(" *************************************************************************")
	fmt.Println(" ******			TEST IPERF - END 		     **********")
	fmt.Println(" *************************************************************************")
}
// chech if there is an error and dispaly a meesage in the case
func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}
//  create a Ketama ring for a given sfc
func createKetamaRingSFC(sfcInpuPaths []string)(*ketama.Continuum, error){
	var buckets []ketama.Bucket

	for i := 0; i <= len(sfcInpuPaths)-2; i++ {
		b := &ketama.Bucket{Label: sfcInpuPaths[i], Weight: 1}
		buckets = append(buckets, *b)
		//fmt.Println(inputsCNFs[i])
	}

	k, err := ketama.New(buckets)
	return k, err
}
// rerurn input cnfs (paths) fior a gievn sfc
func getInputCnfsSFC(sfcID int)[]string{
	fmt.Println(" *************************************************************************")
	fmt.Println(" ******			Get Inputs Paths CNFs 		     ******************************")
	fmt.Println(" *************************************************************************")
	// 1. get INput paths CNF - Pods names  (to execute later iperf on these cnfs)
	cmdget := exec.Command("kubectl", "get", "pods", )
	cmdgrep := exec.Command("grep", "-oE","cnf1-sfc"+strconv.Itoa(sfcID)+"\\S*")

	// Get cmdget's stdout and attach it to grep's stdin.
	pipe, _ := cmdget.StdoutPipe()
	defer pipe.Close()

	cmdgrep.Stdin = pipe

	// Run cmdget first.
	cmdget.Start()

	// Run and get the output of grep.
	res, _ := cmdgrep.Output()
	a:=string(res)
	inputsCNFs := strings.Split(a,"\n")
	//inputsCNFs contain all the inputs CNFs, whereas the last element is a whitespace one because of the '\n'
	//inputsCNFs=inputsCNFs[0:len(inputsCNFs)-2]
	return inputsCNFs
}
