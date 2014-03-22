#Go开发工具包

* 配置文件管理工具包

* 加密工具包
	
	hash文件夹中存放的是用于生存哈希值的go文件，提供了MD5和SHA1两种算法,appkey.go是接口文件

    	type IAppkey interface {
	    GenerateAppKey(...string) string
	    }


    该方法参数是一个可变字符串类型，用于表示待哈希的字符串，返回值为哈希值
    md5.go和sh1.go实现了该接口。
	
    


* MongoDB链接管理工具包