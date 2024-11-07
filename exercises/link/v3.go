package main

import (
	"priprema/v3"
	//"fmt"
)

func main(){

	/*
	p := v3.Pravougaonik{
		A: 10,
		B: 5,
	}

	t := v3.Trougao{
		A: 3,
		B: 4,
		C: 5,
	}

	fmt.Printf("%d", p.Obim())
	fmt.Println()
	fmt.Printf("%d", t.Obim())
	*/

	//version 1
	/*
	array := []int{111, 51, 32, 3, 12, 33, 65, 4}
	c := v3.StartLocals(array)
	minimum := v3.GlobalMini(c, array[0])
	fmt.Printf("%v", minimum)
	*/
	//version 2

	array := []int{-111, 1, 32, 3, -5, 33, 65, 4}
	start := make(chan int, 4)
	v3.StartLocals2(array, start)	

}