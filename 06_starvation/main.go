package main

import (
	"fmt"
	"sync"
	"time"
)

//기아상태
//동시 프로세스가 작업을 수행하는데 필요한 모든 리소스 상태를 얻을 수 없는 상태


func main(){
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1*time.Second 
	
	greedyWorker := func(){
		defer wg.Done()

		var count int 
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3*time.Nanosecond)
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWorker := func(){
		defer wg.Done()

		var count int 
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()
			
			count++
		}
		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}