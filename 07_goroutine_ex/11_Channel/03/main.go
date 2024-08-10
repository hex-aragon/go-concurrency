package main

import (
	"fmt"
	"sync"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			<-begin //고루틴들 대기 

			fmt.Printf("%v has began\n",i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	fmt.Println("close 1")
	close(begin)  //채널 닫으면 모든 고루틴이 대기상태에서 벗어남 
	fmt.Println("close 2")
	wg.Wait()
}