package main

import (
	server "DiffPBServicePart/rpcserver"
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"log"
	"manyPBServices/entity"
)

func main() {

	stub, client, err := server.DialArithServiceMultiply("tcp", "192.168.2.172:62228")
	//stub, client, err := server.DialArithServiceDivide("tcp", "192.168.2.172:62228")
	if err != nil {
		log.Fatal(`arith.DialArithService("tcp", "192.168.2.172:62228"`, err)
	}
	defer client.Close()

	var args entity.ArithRequest
	var reply entity.ArithResponse

	args.A = proto.Int32(7)
	args.B = proto.Int32(2)

	if err = stub.Multiply(&args, &reply); err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("Arith: %d/%d=%d\n", args.GetA(), args.GetB(), reply.GetVal())
}
