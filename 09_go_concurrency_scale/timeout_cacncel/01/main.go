package main

func main() {
	var value interface{}
	done := make(chan interface{})
	valueStream := make(chan interface{})
	select {
	case <-done:
		return 
	case value = <-valueStream:
	}

	result := reallyLongCalculation(value)

	select {
	case <-done:
		return 
	case resultStream<-result:
	}

	reallyLongCalculation := func(
		done <- chan interface{},
		value interface{},
	) interface{} {
		intermediateResult := reallyLongCalculation(value)
		select {
		case <-done:
			return nil 
		default:
		}

		return longCaluclation(intermediateResult)
	}

	reallyLongCalculation := func(
		done<- chan interface{},
		value interface{},
	) interface{} {
		intermediateResult := longCalculation(done, value)
		return longCaluclation(done, intermediateResult)
	}
	
}