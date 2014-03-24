//
//  main.go
//  golang
//
//  Created by zhangzhe on 2013-8-25
//
package main

import (
	f "configureUtil/configure"
	"fmt"
)

const (
	CONFIGUREFILEPATH = "./conf/system.conf"
)

type Integer string

func main() {

	//第一种使用方式
	var path f.ConfFilePath = CONFIGUREFILEPATH
	f.SetConfigurePath(path)
	fmt.Println(f.Value("datebase", "IP"))

	//第二种使用方式
	temp := f.Conf(path, "NodeIP")
	fmt.Println(temp.Name("general"))
	fmt.Println(temp.Name("system"))
	fmt.Println(temp.Name("user"))

}
