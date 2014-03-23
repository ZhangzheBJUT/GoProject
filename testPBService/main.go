package main

import (
	//	"testPBService/arith"
	"fmt"
	"testPBService/rpcserver"
	"testPBService/server"
)

func main() {
	fmt.Println("MyServer start.....")
	rpcserver.ListenAndServeArithService("tcp", "192.168.2.172:62222", new(server.Arith))
}
