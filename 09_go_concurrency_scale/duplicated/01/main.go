package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{},
		id int,
		wg *sync.WaitGroup, 
		result chan<- int ,
	) {
		started := time.Now()
		defer wg.Done()


		//랜덤한 작업 부하 시뮬레이션 
		simulatedLoadTime := time.Duration(1+rand.Intn(5)) * time.Second
		select {
		case <-done:
		case <-time.After(simulatedLoadTime):
		}

		select {
		case <-done:
		case result <- id:
		}

		took := time.Since(started)
		//핸들러가 얼마나 오래 걸리는지 표시
		if took < simulatedLoadTime {
			took = simulatedLoadTime
		}
		fmt.Printf("%v took %v\n", id, took)
	}

	done := make(chan interface{})
	result := make(chan int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {  //핸들러 10개 시작 요청을 처리할 핸들러,
		go doWork(done, i, &wg, result)
	}

	firstReturned := <-result //여러 개의 핸들러 중에서 처음으로 리턴된 값을 받는다. 
	close(done)	//나머지 핸들러를 모두 취소
	wg.Wait()

	fmt.Printf("Received an answer from #%v\n", firstReturned)
}