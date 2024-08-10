package main

import (
	"fmt"
	"sync"
)

//golang fork - join 모델
func main() {
	var wg sync.WaitGroup
	sayHello := func(){
		defer wg.Done()
		fmt.Println("Hello")
	}
	wg.Add(2)
	go sayHello() // 자체 고루틴 실행 
	go sayHello() // 자체 고루틴 실행 
	wg.Wait() //메인 고루틴에서 분기됐던 고루틴이 다시 합류하는 지점 
}