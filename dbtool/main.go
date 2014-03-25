//
//  main.go
//  golang
//
//  Created by zhangzhe on 2014-1-25
//

package dbtool

import (
	"dbtool/mongo"
	"fmt"
	"os"
)

const (
	IP       = "127.0.0.1:21017"
	User     = "Admin"
	Password = "1234"
)

type Person struct {
	NAME  string
	PHONE string
}

func main() {

	handler, err := mongo.NewHandler(IP, User, Password)
	defer handler.DisConnect()

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	collection := handler.Collection("ZZ", "test")

	temp := &Person{
		PHONE: "123456",
		NAME:  "zhangzhe",
	}

	//一次可以插入多个对象 插入两个Person对象
	err = collection.Insert(temp)
	if err != nil {
		panic(err)
	}
}
