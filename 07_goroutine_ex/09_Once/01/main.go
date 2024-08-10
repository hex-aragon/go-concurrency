package main

import (
	"fmt"
	"sync"
)

func main() {
    var count int // 증가시킬 변수

    // increment 함수: count를 1 증가시킴
    increment := func() {
        count++
    }

    var once sync.Once // 한 번만 실행을 보장하는 Once 객체

    var increments sync.WaitGroup
    increments.Add(100) // 100개의 고루틴을 기다리도록 설정

    // 100개의 고루틴 생성
    for i := 0; i < 100; i++ {
        go func() {
            defer increments.Done()
            once.Do(increment) // increment 함수를 한 번만 실행
        }()
    }

    increments.Wait() // 모든 고루틴이 완료될 때까지 대기
    fmt.Printf("Count is %d\n", count) // 최종 count 값 출력
}