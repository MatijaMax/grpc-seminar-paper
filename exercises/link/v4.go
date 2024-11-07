package main

import (
	"priprema/v4"
	//"fmt"
    //"sync"
)

func main(){

    /*
    var wg sync.WaitGroup
    brojac := v4.Counter{
        I: 0,
    }
    wg.Add(1)
    go v4.Grow(&brojac, &wg)
    wg.Add(1)
    go v4.Grow(&brojac, &wg)
    wg.Add(1)
    go v4.Grow(&brojac, &wg)
    wg.Add(1)
    go v4.Grow(&brojac, &wg)
    wg.Wait()
    */

    brojac := v4.Counter{
        I: 0,
    }

    done := make(chan bool)

    go v4.GrowC(&brojac, done)
    go v4.GrowC(&brojac, done)
    go v4.GrowC(&brojac, done)
    go v4.GrowC(&brojac, done)

    <- done


}