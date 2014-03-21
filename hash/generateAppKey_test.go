package appkey

import (
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
