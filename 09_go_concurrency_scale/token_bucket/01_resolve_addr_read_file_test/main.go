package main

import (
	"context" // 컨텍스트 패키지: 작업의 마감 시간, 취소 신호, 요청 범위를 다루기 위해 사용
	"log"     // 로그 패키지: 프로그램 실행 중 로그를 기록하는 데 사용
	"os"      // OS 패키지: 운영 체제 기능을 이용하기 위해 사용
	"sync"    // 동기화 패키지: 고루틴 간의 동기화를 위한 WaitGroup 등을 제공
)

type APIConnection struct{} // APIConnection 구조체 정의

// Open 함수: APIConnection 객체를 생성하여 반환
func Open() *APIConnection {
	return &APIConnection{}
}

// ReadFile 메서드: 파일을 읽는 작업을 수행하는 메서드(현재는 동작하지 않음)
func (a *APIConnection) ReadFile(ctx context.Context) error {
	// 실제 파일 읽기 동작은 주석 처리됨
	return nil // 오류 없이 nil 반환
}

// ResolveAddress 메서드: 주소를 확인하는 작업을 수행하는 메서드(현재는 동작하지 않음)
func (a *APIConnection) ResolveAddress(ctx context.Context) error {
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
