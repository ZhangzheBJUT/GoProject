//
//  main.go
//  golang
//
//  Created by zhangzhe on 2014-3-25
//
package main

import (
	"fmt"
	"testPBServicePart/rpcserver"
	"testPBServicePart/server"
)

const (
	Divide   = "ArithServiceDivide"
	Multiply = "ArithServiceMultiply"
)

func main() {

	fmt.Println("MyServer start.....")

	s := rpcserver.NewServer("tcp", "192.168.2.172:62228")

	s.AddService(Divide, new(server.ArithDivide))
	s.AddService(Multiply, new(server.ArithMutiply))

	err := s.StartServer()
	fmt.Println(err)
}
