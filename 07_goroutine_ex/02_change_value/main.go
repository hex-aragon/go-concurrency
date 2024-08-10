package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutaion := "hello"
	wg.Add(1)
	//고루틴 클로저
	go func(){
		defer wg.Done()
		salutaion = "welcome" //고루틴 값 변경 
	}()
	wg.Wait()
	fmt.Println(salutaion)
}