package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		defer wg.Done()
		fmt.Println("1 goroutine sleeping")
		time.Sleep(1)
	}()
	
	wg.Add(1)
	go func(){
		defer wg.Done()
		fmt.Println("2 goroutine sleeping")
		time.Sleep(1)
	
	}()

	wg.Wait()
	fmt.Println("all goroutine sleeping")
}