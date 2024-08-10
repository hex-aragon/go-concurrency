package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
}