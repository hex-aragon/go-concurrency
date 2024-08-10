package main

import (
	"fmt"
)

func main() {
	// TODO: Add your code here
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings{
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)
	fmt.Println("Done.")
}