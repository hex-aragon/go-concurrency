package main

import "fmt"

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!!!"
	}()

	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v\n", ok, salutation)


	intStream := make(chan int)
	close(intStream)
	integer, ok := <- intStream 
	fmt.Printf("(%v), %v\n", ok, integer)


	intStream2 := make(chan int)
	go func(){
		//defer close로 채널 안닫으면 데드락 걸림
		defer close(intStream2)
		for i := 1; i <= 5; i++ {
			intStream2 <- i 
		}
	}()

	for integer := range intStream2 {
		//fmt.Println(integer)
		fmt.Printf("%v \n", integer)
	}
}