package main

import (
	"fmt"
	"sync"
)


func ConcurrencyTest(){
	var memoryAccess sync.Mutex  //동기화 변수 설정 
	var value int 

	go func() {
		memoryAccess.Lock() //잠금
		value++ 
		memoryAccess.Unlock() //잠금 해제 
	}()

	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is 0\n")
	} else {
		fmt.Printf("the value is %v. \n",value)
	}
	memoryAccess.Unlock()
}

func main(){
	ConcurrencyTest()
}
