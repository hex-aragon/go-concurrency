package main

import (
	"fmt"
	"sync"
)

func main() {
    var numCalcsCreated int // 생성된 계산기(메모리 블록) 수를 추적

    // sync.Pool 객체 생성 및 초기화
    calcPool := &sync.Pool{
        New: func() interface{} {
            numCalcsCreated += 1
            mem := make([]byte, 1024) // 1KB 크기의 메모리 블록 생성
            return &mem
        },
    }

    // 초기에 4개의 객체를 Pool에 추가
    calcPool.Put(calcPool.New())
    calcPool.Put(calcPool.New())
    calcPool.Put(calcPool.New())
    calcPool.Put(calcPool.New())

    const numWorkers = 1024 * 1024 // 동시에 실행할 작업자(고루틴) 수
    var wg sync.WaitGroup
    fmt.Println("numWorkers", numWorkers)
    wg.Add(numWorkers)

    // 많은 수의 고루틴을 생성하여 Pool 사용
    for i := numWorkers; i > 0; i-- {
        go func() {
            defer wg.Done()
            mem := calcPool.Get().(*[]byte) // Pool에서 메모리 블록 가져오기
            defer calcPool.Put(mem) // 사용 후 Pool에 반환
            // 여기서 mem을 사용한 작업을 수행할 수 있음
        }()
    }

    wg.Wait() // 모든 고루틴이 완료될 때까지 대기
    fmt.Printf("%d calculators were created.\n", numCalcsCreated)
}


// 이 코드의 동작 방식을 논리적으로 설명하면 다음과 같습니다:

// sync.Pool 객체를 생성합니다. 이 Pool은 1KB 크기의 메모리 블록([]byte)을 관리합니다.
// Pool의 New 함수는 새로운 메모리 블록이 필요할 때 호출됩니다. 이 함수는 메모리 블록을 생성하고 numCalcsCreated를 증가시킵니다.
// 초기에 4개의 메모리 블록을 생성하여 Pool에 추가합니다.
// 1,048,576개(1024 * 1024)의 고루틴을 생성합니다. 각 고루틴은:

// Pool에서 메모리 블록을 가져옵니다(Get()).
// 메모리 블록을 사용합니다 (이 예제에서는 실제 사용은 생략되어 있습니다).
// 사용이 끝나면 메모리 블록을 Pool에 반환합니다(Put()).


// 모든 고루틴이 완료될 때까지 기다립니다.
// 최종적으로 생성된 메모리 블록(계산기)의 수를 출력합니다.

// 이 코드의 주요 포인트는 sync.Pool의 사용입니다. sync.Pool은 임시 객체의 집합을 관리하는 데 사용됩니다. 이는 다음과 같은 이점을 제공합니다:

// 메모리 할당 및 가비지 컬렉션의 부하를 줄입니다.
// 동시성 환경에서 객체를 효율적으로 재사용할 수 있게 합니다.

// 이 예제에서는 백만 개 이상의 고루틴이 메모리 블록을 사용하지만, 실제로 생성되는 메모리 블록의 수는 그보다 훨씬 적을 것입니다. 이는 sync.Pool이 객체를 효과적으로 재사용하기 때문입니다.
// 이러한 패턴은 많은 수의 임시 객체를 사용하는 고성능 서버 애플리케이션에서 특히 유용합니다. 메모리 사용을 최적화하고 가비지 컬렉션의 부하를 줄이는 데 도움이 됩니다.