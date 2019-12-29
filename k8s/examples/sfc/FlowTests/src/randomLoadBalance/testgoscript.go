package main

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
// to run:
// go run 01-simple-exec-v2.go

import (
	"log"
	"os/exec"
	"fmt"
	//"runtime"
	"strings"
	"math/rand"
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
	temp := strings.Split(a,"\n")
	//

	for i := 0; i < 10; i++ {
		/*s := k.Hash("foo" + strconv.Itoa(i))
		fmt.Println(i, s)
		m[s]++*/
		j:=rand.Intn(3)
		cmd := exec.Command("kubectl", "exec", temp[j], "--","iperf", "-c", "192.168.187.2")
		//cmd := exec.Command("kubectl", "get", "pods", )
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		fmt.Printf("combined out:\n%s\n", string(out))
		fmt.Println("iperf Path \n%s\n", j, " Direction 1 - test iperf")
	}
    
}

