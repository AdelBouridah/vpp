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
	//"encoding/json"
	"sync"
	//"bytes"
	//"time"
)
func iperfinstance(cnf string, port string, i int, wg *sync.WaitGroup){
	fmt.Println(" ************************************************************************************************")
	fmt.Println(" ******		Test IPERF Flux "+strconv.Itoa(i)+" Path "+cnf+" Port "+port+"     **********")
	fmt.Println(" *************************************************************************************************")
	cmd := exec.Command("kubectl", "exec", cnf , "--","iperf", "-c", "192.168.187.2", "-p", port, "-y", "C") // CSV Output  "-y", "C"
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}	
	fmt.Println(" "+string(out)+"test iperf ")
	defer wg.Done()
	f, _ := os.Create("tmp/file-path"+strconv.Itoa(i)+"CNF-"+cnf+"-port-"+port+".txt")
	defer f.Close()
	//fmt.Println("test iperf ********\n"+string(out)+"\n ***")
	f.WriteString(cnf+","+string(out))
	f.Sync()
}

func main() {
	fmt.Println(" *************************************************************************")
	fmt.Println(" ******			Get Inputs PAths CNFs 		     **********")
	fmt.Println(" *************************************************************************")
	// 1. get INput paths CNF - Pods names  (to execute later iperf on these cnfs)
	cmdget := exec.Command("kubectl", "get", "pods", )
	cmdgrep := exec.Command("grep", "-oE","rep-linux-cnf1\\S*")

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


	fmt.Println(" *************************************************************************")
	fmt.Println(" ******			Create the Ring for Ketama           **********")
	fmt.Println(" *************************************************************************")
	var buckets []ketama.Bucket

	for i := 0; i <= len(inputsCNFs)-2; i++ {
		b := &ketama.Bucket{Label: inputsCNFs[i], Weight: 1}
		buckets = append(buckets, *b)
		//fmt.Println(inputsCNFs[i])
	}

	k, _ := ketama.New(buckets)

	fmt.Println(" *************************************************************************")
	fmt.Println(" ******			TEST IPERF - BEGIN 		     **********")
	fmt.Println(" *************************************************************************")
	nbrFlux, _ := strconv.Atoi(os.Args[1])
	fmt.Println(" Flux Numbers : "+strconv.Itoa(nbrFlux))
	var wg sync.WaitGroup
	wg.Add(nbrFlux)
	fmt.Println("Running for loop…")
	mPorts := make(map[string]int)
	for i := 0; i < nbrFlux; i++ {
		s := k.Hash("iperf" + strconv.Itoa(i))
		mPorts[s]++
		if (mPorts[s]<=9){
			go iperfinstance(s, "500"+strconv.Itoa(mPorts[s]), i, &wg)
		}else{
			go iperfinstance(s, "50"+strconv.Itoa(mPorts[s]), i, &wg)
		}
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
