package main

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"log"
	"testPBService/arith"
)

func main() {

	stub, client, err := arith.DialArithService("tcp", "192.168.2.172:62222")
	if err != nil {
		log.Fatal(`arith.DialArithService("tcp", "192.168.2.172:62222"`, err)
	}
	defer client.Close()

	var args arith.ArithRequest
	var reply arith.ArithResponse

	args.A = proto.Int32(7)
	args.B = proto.Int32(8)

	if err = stub.Multiply(&args, &reply); err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.GetA(), args.GetB(), reply.GetVal())
}
