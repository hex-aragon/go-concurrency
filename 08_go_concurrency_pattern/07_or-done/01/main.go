package main

import (
	"fmt"
	"time"
)

func main() {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	done := make(chan interface{})
	myChan := make(chan interface{})
	defer close(myChan)

	// 데이터를 보내는 고루틴 추가
	go func() {
		defer close(done)
		for i := 0; i < 5; i++ {
			myChan <- i
			time.Sleep(1 * time.Second) // 1초마다 값 전송
		}
	}()

	for val := range orDone(done, myChan) {
		fmt.Println("val", val)
	}
}
