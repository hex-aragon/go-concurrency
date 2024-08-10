package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	fmt.Println("1")
	myPool.Get()
	fmt.Println("2")
	instance := myPool.Get()
	fmt.Println("3")
	myPool.Put(instance)
	fmt.Println("4")
	myPool.Get()
	fmt.Println("5")
}