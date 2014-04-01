package main

import (
	"fmt"
	"time"
)

func ProduceData(c chan int64) {

	time.Sleep(time.Second * 6)
	//time.Sleep(time.Second * 3)
	c <- 32
}
func main() {

	//首先实现一个匿名的超时等待函数
	timeout := make(chan bool, 1)
	defer close(timeout)

	go func() {
		time.Sleep(time.Second * 4)
		timeout <- true
	}()

	c := make(chan int64, 2)

	go ProduceData(c)

	select {
	case data := <-c:
		fmt.Println("Read a data", data)
	case <-timeout:
		fmt.Println("timeout......")
	}

}
