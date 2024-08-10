package main

import (
	"fmt"
)

var data int 
	
func main(){
	//메모리 접근 동기화 
	//동시에 실행되는 두 프로세스가 동일한 메모리 영역에 접하려고 시도하고 있음
	//각 프로세스가 메모리에 접근하는 방식은 원자적이지 않는다고 가정
	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("the value is 0\n")
	} else {
		fmt.Printf("the value is %v. \n",data)
	}
}