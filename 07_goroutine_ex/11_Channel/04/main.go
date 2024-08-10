package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
    var stdoutBuff bytes.Buffer
    defer stdoutBuff.WriteTo(os.Stdout) // 프로그램 종료 직전에 버퍼의 내용을 표준 출력으로 출력

    intStream := make(chan int, 4) // 버퍼 크기가 4인 정수 채널 생성

    // 생산자 고루틴
    go func() {
        defer close(intStream) // 채널 닫기
        defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
        for i := 0; i < 4; i++ {
            fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
            intStream <- i // 채널에 정수 보내기
        }
    }()

    // 소비자 (메인 고루틴)
    for integer := range intStream {
        fmt.Fprintf(&stdoutBuff, "Received %v. \n", integer)
    }
}