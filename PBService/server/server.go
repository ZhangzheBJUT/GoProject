package server

import (
	"PBService/arith"
	"code.google.com/p/goprotobuf/proto"
	"errors"
)

type Arith int

func (t *Arith) Multiply(args *arith.ArithRequest, reply *arith.ArithResponse) error {
	reply.Val = proto.Int32(args.GetA() * args.GetB())
	return nil
}

func (t *Arith) Divide(args *arith.ArithRequest, reply *arith.ArithResponse) error {
	if args.GetB() == 0 {
		return errors.New("divide by zero")
	}
	reply.Quo = proto.Int32(args.GetA() / args.GetB())
	reply.Rem = proto.Int32(args.GetA() % args.GetB())
	return nil
}
