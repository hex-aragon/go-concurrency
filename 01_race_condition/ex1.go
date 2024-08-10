package main

import (
	"fmt"
	"time"
)


func main(){
	//레이스 컨디션
	//둘 이상의 작업이 올바른 순서로 실행되어야 하지만 프로그램이 그렇게 작성되지 않아서 이 순서가 유지되는 것이 보장되지 않을 때 발생 
	var data int 
	
	go func() {
		data++
	}()
	
	
	time.Sleep(1*time.Second)
	fmt.Println("data",data)
	if data == 0 {
		fmt.Printf("the value is %v. \n",data)
	}
}