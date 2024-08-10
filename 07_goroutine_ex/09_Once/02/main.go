package main

import (
	"fmt"
	"sync"
)

func main() {
    var count int // 증가시킬 변수

    increment := func() {
        count++
    }

    decrement := func() {
        count--
    }

    var once sync.Once // 한 번만 실행을 보장하는 Once 객체
    once.Do(increment)
    once.Do(decrement)
    
    fmt.Printf("Count is %d\n", count) // 최종 count 값 출력
}