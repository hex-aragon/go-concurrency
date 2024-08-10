package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys 
	}

	//수신 전용 채널 
	var c <- chan interface{}
	var wg sync.WaitGroup
	noop := func() {wg.Done(); <-c}

	const numGoroutines = 1e4 
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i >0; i-- {
		//fmt.Println("i",i)
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%3.fkb\n", float64(after-before)/numGoroutines/1000)
}