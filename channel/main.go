package main

import (
	"fmt"
	"time"
)

var c = make(chan int, 10)

func Producer() {

	var counter int = 0

	for {
		c <- counter
		counter++
		time.Sleep(time.Second * 2)
	}

}

func Consumer() {

	for {
		counter := <-c
		time.Sleep(time.Second * 3)
		fmt.Println("counter:", counter)
	}

}

func main() {
	go Producer()
	go Consumer()

	for {
		time.Sleep(time.Second * 1000)
	}
}
