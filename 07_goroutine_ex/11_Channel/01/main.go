package main

import "fmt"

func main() {
	var receiveChan <-chan interface{}
	var sendChan chan<- interface{}
	intStream := make(chan int)
	dataStream := make(chan interface{})

	receiveChan = dataStream
	sendChan = dataStream
	fmt.Println(receiveChan)
	fmt.Println(sendChan)
	fmt.Println(intStream)

	fmt.Printf("%T\n",receiveChan)
	fmt.Printf("%T\n",sendChan)
	fmt.Printf("%T\n",intStream)

	stringStream := make(chan string)
	go func(){
		stringStream <- "Hello Channels !!"
	}()

	fmt.Println("<-stringStream", <-stringStream)
	
}