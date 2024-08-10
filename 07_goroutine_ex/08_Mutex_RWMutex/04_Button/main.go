package main

import (
	"fmt"
	"sync"
)

// Button 구조체: 클릭 이벤트를 나타내는 Condition Variable을 포함
type Button struct {
    Clicked *sync.Cond
}

func main() {
    // 버튼 객체 생성 및 Condition Variable 초기화
    button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

    // subscribe 함수: Condition Variable에 대한 리스너(고루틴) 등록
    subscribe := func(c *sync.Cond, fn func()) {
        var goroutineRunning sync.WaitGroup
        goroutineRunning.Add(1)
        go func() {
            goroutineRunning.Done()
            c.L.Lock()
            defer c.L.Unlock()
            c.Wait() // 신호를 기다림
            fn()     // 신호 수신 시 실행할 함수
        }()
        goroutineRunning.Wait()
    }

    var clickRegistered sync.WaitGroup
    clickRegistered.Add(3) // 3개의 리스너 등록 예정

    // 첫 번째 리스너 등록
    subscribe(button.Clicked, func() {
        fmt.Println("Maximizing window.")
        clickRegistered.Done()
    })

    // 두 번째 리스너 등록
    subscribe(button.Clicked, func() {
        fmt.Println("Displaying annoying dialog box!")
        clickRegistered.Done()
    })

    // 세 번째 리스너 등록
    subscribe(button.Clicked, func() {
        fmt.Println("Mouse clicked.")
        clickRegistered.Done()
    })

    // 모든 리스너에게 신호 전송
	// 한번 호출하면 세 개의 핸들러가 모두 실행됨, Cond 타입을 사용하는 주된 이유
    button.Clicked.Broadcast()

    // 모든 리스너의 실행 완료를 기다림
    clickRegistered.Wait()
}