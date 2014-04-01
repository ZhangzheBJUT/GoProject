package main

import (
	"fmt"
	"log"
	"time"
)

func doWork() {
	time.Sleep(time.Second * 5)
	fmt.Println("do workd......")
	panic("exception ocurred....")
}
func startServer() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)

		}

	}()
	doWork()

}
func main() {

	startServer()
}
