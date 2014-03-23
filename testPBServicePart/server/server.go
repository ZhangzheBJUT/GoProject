package server

import (
	"code.google.com/p/goprotobuf/proto"
	"errors"
	"testPBServicePart/entity"
)

type ArithMutiply int

func (this *ArithMutiply) Multiply(args *entity.ArithRequest, reply *entity.ArithResponse) error {
	reply.Val = proto.Int32(args.GetA() * args.GetB())
	return nil
}

type ArithDivide int

func (this *ArithDivide) Divide(args *entity.ArithRequest, reply *entity.ArithResponse) error {

	if args.GetB() == 0 {
		return errors.New("divide by zero")
	}
	reply.Quo = proto.Int32(args.GetA() / args.GetB())
	reply.Rem = proto.Int32(args.GetA() % args.GetB())
	return nil
}
