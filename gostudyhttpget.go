package main

import (
	"fmt"
	"io/ioutil" //到处io工具包
	"net/http"  //导入http请求包
)

//发送http请求

func main() {
	//发送请求
	res, err := http.Get("http://wwww.baidu.com")
	if err != nil {
		fmt.Println("请求出错了！")
	}
	//必须关闭流
	if res != nil && res.Body != nil {
		defer res.Body.Close()
		//读取内容
		body, error := ioutil.ReadAll(res.Body)
		if error != nil {
			fmt.Println(error)
		}
		fmt.Println(string(body))
	}

}
