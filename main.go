package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var sum int

func main() {
	wg.Add(1)
	go testGo()
	wg.Wait()
	fmt.Println("합: ", sum)
}
func testGo() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		go func() {
			sum += 1
		}()
	}
	fmt.Println("바깥쪽")
}
