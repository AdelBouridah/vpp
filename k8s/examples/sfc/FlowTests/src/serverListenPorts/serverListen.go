package main

import (
	"fmt"
	//"ketama"
	"strconv"
	"os/exec"
	"strings"
	"log"
	"os"
	//"io"
	//"encoding/json"
	//"sync"
	//"bytes"
)
func main() {
	fmt.Println(" *************************************************************************")
	fmt.Println(" ******			Get Output Paths CNFs		     **********")
	fmt.Println(" *************************************************************************")
	// 2. get Output paths CNF - Pods names  (to execute later iperf on these cnfs)
	cmd2get := exec.Command("kubectl", "get", "pods", )
	cmd2grep := exec.Command("grep", "-oE","rep-linux-cnf5\\S*")

	// Get cmdget's stdout and attach it to grep's stdin.
	pipe2, _ := cmd2get.StdoutPipe()
	defer pipe2.Close()

	cmd2grep.Stdin = pipe2

	// Run cmdget first.
	cmd2get.Start()

	// Run and get the output of grep.
	res2, _ := cmd2grep.Output()
	a2:=string(res2)
	outputsCNFs := strings.Split(a2,"\n")
	//outputsCNFs contain all the inputs CNFs, whereas the last element is a whitespace one because of the '\n'
	//tests
	for i := 0; i <= len(outputsCNFs)-2; i++ {
		fmt.Println("cnf outputs	", outputsCNFs[i])
	}
	fmt.Println(" **************************************************************************************************")
	fmt.Println(" ******   Make output iperf servers listen on sevral Ports -For parallel Flows test      **********")
	fmt.Println(" **************************************************************************************************")
	//var int endPorts
	beginPortsRead, _ := strconv.Atoi(os.Args[1])
	endPortsRead, _ := strconv.Atoi(os.Args[2])
	//fmt.Println(" end ports : "+strconv.Itoa(endPorts))
	for i := 0; i <= len(outputsCNFs)-2; i++ {
		// Re-initialize begin and end port for the replicat 
		beginPorts:=beginPortsRead
		endPorts:=endPortsRead
		if (endPorts<=9){
			for j := beginPorts; j <= endPorts; j++ {
				//fmt.Println("cnf outputs	", outputsCNFs[i])
				cmd := exec.Command("kubectl", "exec", outputsCNFs[i] , "--","iperf", "-s", "-u", "-p", "500"+strconv.Itoa(j))
				//cmd := exec.Command("kubectl", "get", "pods", )
				err := cmd.Start()
				if err != nil {
					log.Fatalf("cmd.Run() failed with %s\n", err)
				}
				//fmt.Printf("combined out:\n%s\n", string(out))
				fmt.Println("CNF listen on  UDP port  ", outputsCNFs[i] , " on port 500"+strconv.Itoa(j))
			}
		}else{
			if (beginPorts<9){
					for j := beginPorts; j <= 9; j++ {
						//fmt.Println("cnf outputs	", outputsCNFs[i])
						cmd := exec.Command("kubectl", "exec", outputsCNFs[i] , "--","iperf", "-s","-u",  "-p", "500"+strconv.Itoa(j))
						//cmd := exec.Command("kubectl", "get", "pods", )
						err := cmd.Start()
						if err != nil {
							log.Fatalf("cmd.Run() failed with %s\n", err)
						}
						//fmt.Printf("combined out:\n%s\n", string(out))
						fmt.Println("CNF listen on UDP  port   ", outputsCNFs[i] , " on port 500"+strconv.Itoa(j))
					}
					beginPorts=10; //to continu from 10 ports in the next spet
			}
			for j := beginPorts; j <= endPorts; j++ {
				//fmt.Println("cnf outputs	", outputsCNFs[i])
				cmd := exec.Command("kubectl", "exec", outputsCNFs[i] , "--","iperf", "-s","-u",  "-p", "50"+strconv.Itoa(j), "-u")
				//cmd := exec.Command("kubectl", "get", "pods", )
				err := cmd.Start()
				if err != nil {
					log.Fatalf("cmd.Run() failed with %s\n", err)
				}
				//fmt.Printf("combined out:\n%s\n", string(out))
				fmt.Println("CNF listen on UDP port  ", outputsCNFs[i] , " on port 50"+strconv.Itoa(j))
			}

		}

	}

}
