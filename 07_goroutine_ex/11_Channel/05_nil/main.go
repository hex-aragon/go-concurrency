package main

func main() {
	var dataStream chan interface{}
	//<-dataStream
	//dataStream <- struct{}{}
	close(dataStream)
}