package main

import (
	"fmt"
	"ketama"
	"strconv"
	"os/exec"
	"strings"
	"log"
)

func main() {
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
	for i := 0; i < 100; i++ {
		s := k.Hash("iperf" + strconv.Itoa(i))
		//fmt.Println(i, s)
		cmd := exec.Command("kubectl", "exec", s , "--","iperf", "-c", "192.168.187.2")
		//cmd := exec.Command("kubectl", "get", "pods", )
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		fmt.Printf("combined out:\n%s\n", string(out))
		fmt.Println("iperf test for Path with  %s", s, " CNF as input")

	}
	fmt.Println(" *************************************************************************")
	fmt.Println(" ******			TEST IPERF - END 		     **********")
	fmt.Println(" *************************************************************************")
}
