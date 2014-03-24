package encode

import (
	"fmt"
	"testing"
)

//************************************
/*
/*  MD5 和  SHA1 使用测试用例
/*
/************************************/
func Test_MD5(t *testing.T) {

	md5 := new(MD5)
	result := md5.GenerateAppKey("admin", "1234")
	t.Log(result)
}

func Test_SHA1(t *testing.T) {

	sha1 := new(SHA1)
	result := sha1.GenerateAppKey("admin", "1234")
	t.Log(result)
}

type MyTest int

func (this *MyTest) GenerateAppKey(args ...string) string {
	fmt.Println("Generate App key!")
	return "Hello"
}

func TestMyTest(t *testing.T) {

	temp := new(MyTest)
	temp.GenerateAppKey("Hello")
}

func Benchmark_MD5(b *testing.B) {

	b.StopTimer() //停止该函数压力测试的时间

	//进行一些其他的操作。。。。。。

	b.StartTimer() //开启该函数压力测试的时间

	for i := 0; i < b.N; i++ {
		md5 := new(MD5)
		md5.GenerateAppKey("admin", "1234")
	}

}
