package main

import (
	"fmt"
	//"runtime"
	"sync"
)

var counter int = 0

func Count(lock *sync.Mutex) {

	lock.Lock()
	counter++
	fmt.Print(counter)
	lock.Unlock()
}

func ChannelTest() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:

		case ch <- 1:
		}

		i := <-ch
		fmt.Println("value received:", i)
	}
}
func main() {

	ChannelTest()

	//lock := &sync.Mutex{}

	//for i := 0; i < 10; i++ {
	//	go Count(lock)
	//}

	//for {
	//	lock.Lock()
	//	c := counter
	//	lock.Unlock()
	//	runtime.Gosched()
	//	if c >= 10 {
	//		break
	//	}
	//}
}
