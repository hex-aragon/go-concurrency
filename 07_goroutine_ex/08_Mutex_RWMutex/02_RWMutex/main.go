package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
    // producer 함수: 잠금을 획득했다가 즉시 해제하는 작업을 5번 반복
    producer := func(wg *sync.WaitGroup, l sync.Locker) {
        defer wg.Done()
        for i := 5; i > 0; i-- {
            l.Lock()
            l.Unlock()
            time.Sleep(1)
        }
    }

    // observer 함수: 잠금을 획득하고 유지
    observer := func(wg *sync.WaitGroup, l sync.Locker) {
        defer wg.Done()
        l.Lock()
        defer l.Unlock()
    }

    // test 함수: producer와 여러 개의 observer를 동시에 실행하고 걸린 시간 측정
    test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
        var wg sync.WaitGroup
        wg.Add(count + 1)
        beginTestTime := time.Now()
        go producer(&wg, mutex)
        for i := count; i > 0; i-- {
            go observer(&wg, rwMutex)
        }

        wg.Wait()
        return time.Since(beginTestTime)
    }

    // 결과를 표 형식으로 출력하기 위한 tabwriter 설정
    tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
    defer tw.Flush()

    var m sync.RWMutex
    fmt.Fprintf(tw, "Readers\tRWMutext\tMutex\n")
    
    // 2의 거듭제곱으로 reader 수를 증가시키며 테스트 실행
    for i := 0; i < 20; i++ {
        count := int(math.Pow(2, float64(i)))
        fmt.Fprintf(
            tw,
            "%d\t%v\t%v\n",
            count,
            test(count, &m, m.RLocker()),  // RWMutex의 읽기 잠금 사용
            test(count, &m, &m),           // 일반 Mutex로 사용
        )
    }
}