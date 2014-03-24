//
//  main.go
//  golang
//
//  Created by zhangzhe on 2014-3-25
//

package main

import (
	"PBService/rpcserver"
	"PBService/server"
	"fmt"
)

func main() {
	fmt.Println("MyServer start.....")
	rpcserver.ListenAndServeArithService("tcp", "192.168.2.172:62222", new(server.Arith))
}
