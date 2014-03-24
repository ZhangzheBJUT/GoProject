/home/zha/goProj/src/PBService
package main

import (
	"fmt"
)

const ngorountine = 20

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {

	leftmost := make(chan int)
	var left, right chan int = nil, leftmost

	for i := 0; i < ngorountine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 6
	x := <-leftmost
	fmt.Println(x)

}
