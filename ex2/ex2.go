package main

import (
	"fmt"
)

var data int 
	
func main(){
	

	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("the value is 0\n")
	} else {
		fmt.Printf("the value is %v. \n",data)
	}
}