package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printNumbers(i, &wg)
	}
	wg.Wait()
}


func printNumbers(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println(i)
}