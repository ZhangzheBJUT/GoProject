package encode

import (
	"testing"
)

func Benchmark_MD5New(b *testing.B) {

	b.StopTimer() //停止该函数压力测试的时间

	//进行一些其他的操作。。。。。。

	b.StartTimer() //开启该函数压力测试的时间

	for i := 0; i < b.N; i++ {
		md5 := new(MD5Encode)
		md5.GenerateAppKey("admin", "1234")
		b.Log(b.N)
	}

}
