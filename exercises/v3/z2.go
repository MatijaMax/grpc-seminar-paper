package v3

import(
	"sync"
	"fmt"
)

func StartLocals(arr []int) chan int{
	c := make(chan int)
	var wg sync.WaitGroup
	for i:=0; i<len(arr); i+=2{
		end := i + 2
		if end > len(arr) {
			end = len(arr)
		}
		wg.Add(1)
		go LocalMini(arr[i: end], c, &wg)		
	}

	go func(){
		wg.Wait()
		close(c)
	}()

	return c
}

func StartLocals2(arr []int, c chan int){ 
	for i:=0; i<len(arr); i+=2{
		end := i + 2
		if end > len(arr) {
			end = len(arr)
		}
		go LocalMini2(arr[i: end], c)		
	}

	

	globalMini := arr[0]
	for i:= 0; i < 4; i++{
		localMini := <-c
		if localMini<globalMini{
			globalMini = localMini
			fmt.Printf("%d", localMini)
		}
	}
	fmt.Printf("Global minimum: %d\n", globalMini)
}

func LocalMini(arr []int, c chan int, wg *sync.WaitGroup){
	defer wg.Done() 
	if len(arr) == 1 {
		c <- arr[0]
	} else if arr[0] < arr[1] {
		c <- arr[0]
	} else {
		c <- arr[1]
	}
}

func LocalMini2(arr []int, c chan int){
	if len(arr) == 1 {
		c <- arr[0]
	} else if arr[0] < arr[1] {
		c <- arr[0]
	} else {
		c <- arr[1]
	}
}


func GlobalMini(c chan int, pivot int) int{
	globalMini := pivot
	fmt.Println("Started global")
	for localMini := range c{
		if localMini<globalMini{
			globalMini = localMini
		}
	}
	return globalMini	
}