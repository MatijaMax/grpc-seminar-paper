package v4

import (
	//"sync"
	"fmt"
)


func GrowC(counter *Counter, done chan bool){
	
	for {
		counter.Mutex.Lock()
		if counter.I < 100{
			counter.I ++
			fmt.Printf("%v", counter.I)
			fmt.Println()
		}else{
			counter.Mutex.Unlock()
			done <- true
			return
		}		
		counter.Mutex.Unlock()
	}
}


