package rpcserver

import (
	"PBService/arith"
	protorpc "code.google.com/p/protorpc"
	"log"
	"net"
	"net/rpc"
)

// ListenAndServeArithService listen announces on the local network address laddr
// and serves the given ArithService implementation.
func ListenAndServeArithService(network, addr string, x arith.ArithService) error {

	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	srv := rpc.NewServer()
	if err := srv.RegisterName("ArithService", x); err != nil {
		return err
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("lis.Accept(): %v\n", err)
		}
		go srv.ServeCodec(protorpc.NewServerCodec(conn))
	}
}
