package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)


func main() {
	log.SetFlags(log.Ltime | log.LUTC)
	log.SetOutput(os.Stdout)

	go func(){
		//기본 제공 프로파일 함수
		goroutines := pprof.Lookup("goroutine")
		for range time.Tick(1*time.Second) {
			log.Printf("goroutine count : %d\n", goroutines.Count())
		}
	}()

	//절대 종료되지 않는 고루틴을 몇 개 생성한다.
	var blockForever chan struct{}
	for i := 0; i < 10; i++{
		go func() {
			<-blockForever
		}()
		time.Sleep(500*time.Millisecond)
	}


	//사용자 제공 프로파일 함수 
	newP := func (name string) * pprof.Profile{
		prof := pprof.Lookup(name)
		if prof == nil {
			prof = pprof.NewProfile(name)
		}
		return prof 
	}
	prof := newP("my_package_namespace")
	fmt.Println("prof",prof)
}

