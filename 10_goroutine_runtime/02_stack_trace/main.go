package main

func main() {
	//go 1.6 이전에는 패닉 에러 발생시 모든 고루틴에서 스택트레이스 에러 출력
	//go 1.6 이후에는 패닉 에러가 발생한 고루틴에서만 스택 트레이스 에러 출력
	waitForever := make(chan interface{})
	go func() {
		panic("test panic")
	}()
	<-waitForever
}