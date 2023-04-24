package main

import (
	"fmt"
	"sync"
	"time"
)

func complexWork() {
	time.Sleep(10 * time.Second)
}

func simpleWork() {
	time.Sleep(10 * time.Millisecond)
}

func testComplex() {
	fmt.Println("Starting complex working.. ")
	complexWork()
	fmt.Println("Complex work done!")
}

func testSimple() {
	fmt.Println("Starting simple working..")
	simpleWork()
	fmt.Println("Simple work done!")
}
func main() {
	// Execution time without goroutines
	startTime := time.Now()

	testComplex()
	testSimple()
	testComplex()

	fmt.Println("Execution time without goroutines:", time.Since(startTime))

	// Execution time with goroutines
	var wg sync.WaitGroup
	wg.Add(1)

	startTime = time.Now()

	//create a channel to recive the result of selected goroutine

	result := make(chan int)

	go func(){
		defer wg.Done()
		
		select {
		case <-time.After(time.Second):
			testComplex()
			result <- 1
		case <-time.After(time.Second):
			testSimple()
			result <- 2
		case <-time.After(time.Second):
			testComplex()
			result <- 3	
		}
	}()

	//wait for the selected goroutine to complete
		<-result	
	
	wg.Wait()

	fmt.Println("Execution time with goroutines:", time.Since(startTime))
}
