package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//live lock : 동시에 프로그램이 연산을 수행하지만 실제 프로그램 상태는 진행상황이 없는 상태

func main(){

	cadence := sync.NewCond(&sync.Mutex{})
	go func(){
		for range time.Tick(1*time.Millisecond){
			cadence.Broadcast()
		}
	}()

	//뮤텍스 잠금, 잠금해제 
	takeStep := func(){
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	//방향 시도 
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1 )
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true 
		}

		takeStep()
		atomic.AddInt32(dir, -1)
		return false 
	}

	var left, right int32 
	tryLeft := func(out *bytes.Buffer) bool {return tryDir("left",&left, out)}
	tryRight := func(out *bytes.Buffer) bool {return tryDir("right", &right, out)}
	
	fmt.Println("tryLeft", tryLeft)
	fmt.Println("tryRight", tryRight)


	walk := func(walking *sync.WaitGroup, name string ){
		var out bytes.Buffer
		defer func() {fmt.Println(out.String())}()
		defer walking.Done()

		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0 ; i < 5; i++{
			if tryLeft(&out) || tryRight(&out) {
				return 
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)

	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()

}