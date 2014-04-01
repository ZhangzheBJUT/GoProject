package main

import (
	"fmt"
	"time"
)

/*
   生产者生成“max”个 int64 的数字，并且将其放入 channel “c” 中。需要注意的是
   这里用 defer 在函数退出的时候关闭了 channel。
*/
func producer(c chan int64, max int) {
	defer close(c)
	for i := 0; i < max; i++ {
		c <- time.Now().Unix()
	}
}

func consumer(c chan int64) {

	var v int64
	ok := true
	for ok {
		if v, ok = <-c; ok {
			fmt.Println(v)
		}
	}
}

func main() {
	c := make(chan int64)
	go producer(c, 10)
	consumer(c)
}
