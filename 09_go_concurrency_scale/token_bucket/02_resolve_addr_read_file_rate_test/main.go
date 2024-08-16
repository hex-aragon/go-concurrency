package main

import (
	"context" // 컨텍스트 패키지: 작업의 마감 시간, 취소 신호, 요청 범위를 다루기 위해 사용
	"log"     // 로그 패키지: 프로그램 실행 중 로그를 기록하는 데 사용
	"os"      // OS 패키지: 운영 체제 기능을 이용하기 위해 사용
	"sync"    // 동기화 패키지: 고루틴 간의 동기화를 위한 WaitGroup 등을 제공
	"time"

	"golang.org/x/time/rate"
)

//Limit는 어떤 이벤트의 최대 빈도를 정의한다.
//Limit는 초당 이벤트의 수를 나타낸다. 0의 Limit는 아무런 이벤트도 허용하지 않는다.
type Limit float64

type Limiter struct {
	r rate.Limit
	b int 
}

//NewLimiter는 새로운 r의 속도를 가짐 최대 b개의 토큰을 가지는 새로운 Limiter를 리턴한다.
func NewLimiter(r Limit, b int) *Limiter {
	return nil 
}
//Every는 Limit에 대한 이벤트 사이의 최소 시간 간격을 변환한다. 
func Every(interval time.Duration) Limit  {
	return 0.1 
}

//ex
//rate.Limit(events/timePeriod.Seconds())

func Per(eventCount int , duration time.Duration) rate.Limit {
	return rate.Every(duration/time.Duration(eventCount))
}

//Wait는 WaitN(ctx, 1)의 축약형이다. 
func(lim *Limiter)  Wait(ctx context.Context) {
}

//WaitN함수는 lim가 n개의 이벤트 발생을 허용할 때까지 대기한다.
//n이 Limiter의 버퍼 사이즈를 초과하면 error를 리턴하며, Context는 취소된다.
//그렇지 않은 경우에는 Context의 Deadline이 지날 때까지 대기한다. 
func (lim *Limiter) WaitN(ctx context.Context, n int) (err error) {
	return nil 
}

// APIConnection 구조체 정의
type APIConnection struct{
	rateLimiter *rate.Limiter
} 


// Open 함수: APIConnection 객체를 생성하여 반환
func Open() *APIConnection {
	return &APIConnection{
		rateLimiter: rate.NewLimiter(rate.Limit(1), 1), //1초당 1개 이벤트라는 속도제한을 검, 모든 API 연결에 대해 
	}
}

// ReadFile 메서드: 파일을 읽는 작업을 수행하는 메서드(현재는 동작하지 않음)
func (a *APIConnection) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil { //속도제한에 의해 요청을 완료하기에 충분한 접근토큰을 가질 때까지 대기한다. wait 함수 
		return err 
	}
	// 실제 파일 읽기 동작은 주석 처리됨
	return nil // 오류 없이 nil 반환
}

// ResolveAddress 메서드: 주소를 확인하는 작업을 수행하는 메서드(현재는 동작하지 않음)
func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err 
	}
	// 실제 주소 확인 동작은 주석 처리됨
	return nil // 오류 없이 nil 반환
}

func main() {
	defer log.Printf("Done.") // main 함수가 끝날 때 "Done." 로그 출력
	log.SetOutput(os.Stdout) // 로그 출력을 표준 출력으로 설정
	log.SetFlags(log.Ltime | log.LUTC) // 로그에 시간과 UTC 기준 시간을 표시하도록 설정

	apiConnection := Open() // APIConnection 객체를 생성
	var wg sync.WaitGroup    // 고루틴 동기화를 위한 WaitGroup 생성
	wg.Add(20)               // 총 20개의 작업을 기다리도록 설정

	// 10개의 ReadFile 고루틴 생성
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done() // 고루틴 작업이 끝나면 WaitGroup에 작업 완료를 알림
			err := apiConnection.ReadFile(context.Background()) // 파일 읽기 수행
			if err != nil {
				log.Printf("cannot ReadFile: %v", err) // 에러 발생 시 로그 출력
			}
			log.Printf("ReadFile") // 파일 읽기 성공 시 로그 출력
		}()
	}

	// 10개의 ResolveAddress 고루틴 생성
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done() // 고루틴 작업이 끝나면 WaitGroup에 작업 완료를 알림
			err := apiConnection.ResolveAddress(context.Background()) // 주소 확인 수행
			if err != nil {
				log.Printf("cannot ResolveAddress: %v", err) // 에러 발생 시 로그 출력
			}
			log.Printf("ResolveAddress") // 주소 확인 성공 시 로그 출력
		}()
	}

	wg.Wait() // 모든 고루틴이 끝날 때까지 대기
}
