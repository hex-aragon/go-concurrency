package main

func main() {
	stringStream := make(chan string)
	done := make(chan string)
	for _, s := range []string {"a", "b", "c",} {
		select {
		case <- done:
			return 
		case stringStream <- s:
		}
	}
}