package main

import (
	"context"
	"log"
	"os"
	"sort"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

func MultiLimiter(limiters ...RateLimiter) *multiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}
	sort.Slice(limiters, byLimit)
	return &multiLimiter{limiters: limiters}
}

type multiLimiter struct {
	limiters []RateLimiter
}

func (l *multiLimiter) Wait(ctx context.Context) error {
	for _, l := range l.limiters {
		if err := l.Wait(ctx); err != nil {
			return err 
		}
	}
	return nil 
}

func (l *multiLimiter) Limit() rate.Limit {
	return l.limiters[0].Limit()
}


func Per(eventCount int , duration time.Duration) rate.Limit {
	return rate.Every(duration/time.Duration(eventCount))
}

type APIConnection struct{
	networkLimit,
	diskLimit,
	apiLimit RateLimiter
} 

func Open() *APIConnection {
	// secondLimit := rate.NewLimiter(Per(2, time.Second),1)
	// minuteLimit := rate.NewLimiter(Per(10, time.Minute),10)
	return &APIConnection{
		apiLimit : MultiLimiter(
			rate.NewLimiter(Per(2, time.Second),1),  
			rate.NewLimiter(Per(10, time.Minute),10)),
		diskLimit: MultiLimiter(
			rate.NewLimiter(rate.Limit(1),1), 
		),
		networkLimit : MultiLimiter(
			rate.NewLimiter(Per(3, time.Second),3),
		),

	}
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	if err := MultiLimiter(a.apiLimit, a.diskLimit).Wait(ctx); err != nil {
		return err 
	}
	//작업하는 척 
	return  nil 
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := MultiLimiter(a.apiLimit, a.networkLimit).Wait(ctx); err != nil {
		return err 
	}
	//작업하는 척 
	return  nil 
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
