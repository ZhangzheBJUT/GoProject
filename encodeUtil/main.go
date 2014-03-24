//
//  main.go
//  golang
//
//  Created by zhangzhe on 2013-10-12
//
package main

import (
	"encodeUtil/appkey"
	"encodeUtil/encode"
	"fmt"
)

func MyAppkey(tool appkey.IAppKey) string {

	return tool.GenerateAppKey("admin", "123")
}

func main() {
	md5 := new(encode.MD5Encode)
	fmt.Println(MyAppkey(md5))

	sha1 := new(encode.SHA1Encode)
	fmt.Println(MyAppkey(sha1))
}
