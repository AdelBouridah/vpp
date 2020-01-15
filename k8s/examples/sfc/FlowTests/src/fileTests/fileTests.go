package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	fmt.Println("Running for loop…")
	for j := 10; j <= 12; j++ {

		go func(i int) {
				defer wg.Done()
				f, _ := os.Create("tmp/file"+strconv.Itoa(i)+".txt")
			        //Check(err)
			        defer f.Close()
				f.WriteString("eeeeeeeeeeee\n")
				f.Sync()
				//fmt.println("tests ********")
		}(j)
	}
	wg.Wait()
	fmt.Println("Finished for loop")

}
