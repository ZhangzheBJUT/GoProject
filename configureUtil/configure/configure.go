//
//  main.go
//  golang
//
//  Created by zhangzhe on 2013-8-25
//
package configure

import (
	"encoding/json"
	"fmt"
	"os"
)

//配置文件路径
type ConfFilePath string

var filePath ConfFilePath

//一个配置项的组成 Key : Value
type Property struct {
	Name  string `json:"Key"`
	Value string
}

//Configure代表一个具体的配置项
type Configure struct {
	Item      string `json:"Item"`
	Configure []Property
}

//Configures代表一个配置项的集合
type Configures struct {
	Configures []Configure `json:"Configures"`
}

/*******************************************
函数描述:
    从配置文件返回配置项为confName的配置项

参数说明:
    配置项对象
********************************************/
func (this *Configures) configure(confName string) *Configure {

	length := len(this.Configures)
	for i := 0; i != length; i++ {
		configure := this.Configures[i]
		if configure.Item == confName {
			return &configure
		}
	}

	return nil
}

/*******************************************
函数描述:
    加载配置文件

参数说明:
    返回配置文件对象
********************************************/
func configure() (result *Configures) {

	file, err := os.Open(string(filePath))
	defer file.Close()

	if err != nil {
		fmt.Println("open file error!")
	}

	buf := make([]byte, 2048)

	n, err := file.Read(buf)
	newBuf := buf[:n]

	if n == 0 {
		fmt.Println("empty!")
		return
	}
	if err != nil {
		fmt.Println("open file error!")
	}

	var conf Configures
	err = json.Unmarshal(newBuf, &conf)
	if err != nil {
		fmt.Println("unmarsh file error!", err)
	}

	return &conf
}

/*******************************************
函数描述:
    返回一个confName的配置项

参数说明:
    path      配置文件路径
	confName  配置项的名称

	result   返回confName指定的配置对象
********************************************/
func Conf(path ConfFilePath, confName string) (result *Configure) {

	filePath = path
	configures := configure()
	result = configures.configure(confName)
	return result
}

/*******************************************
函数描述:
    返回对象中key为confName的值得

参数说明:
    confName为Key的值
********************************************/
func (this *Configure) Name(confName string) string {

	length := len(this.Configure)
	for i := 0; i != length; i++ {
		configure := this.Configure[i]
		if configure.Name == confName {
			return configure.Value
		}
	}

	return ""
}

func SetConfigurePath(path ConfFilePath) {
	filePath = path
}

/*******************************************
函数描述:
    返回一个Value的配置项

参数说明:
   confName  配置对象的名称
   key       配置对象中key对应的值
*******************b*************************/
func Value(confName string, key string) string {
	conf := Conf(filePath, confName)
	return conf.Name(key)
}
