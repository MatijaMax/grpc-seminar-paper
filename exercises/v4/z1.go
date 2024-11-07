package v4

import (
	"sync"
	"fmt"
)

type Counter struct{
	I int
	Mutex sync.Mutex
}

func Grow(counter *Counter, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		counter.Mutex.Lock()
		if counter.I < 100{
			counter.I ++
			fmt.Printf("%v", counter.I)
			fmt.Println()
		}else{
			counter.Mutex.Unlock()
			return
		}		
		counter.Mutex.Unlock()
	}
}


