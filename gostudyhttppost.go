package main

//携带参数的复杂post请求
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//发送请求
func main() {
	//构建请求体
	v := url.Values{}
	v.Set("mobile", "18111111111")
	//构造请求内容。
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //编译后的内容作为请求体
	//创建客户端
	client := &http.Client{}
	//构建请求
	request, err1 := http.NewRequest("POST", "http://passport.baidu.com?password=111", body)
	if err1 != nil {
		fmt.Println("请求出错了1")
	}
	request.Header.Set("Content-type", "application/x-wwww-form-urlencoded;param=value")
	//发送请求
	response, err := client.Do(request)

	defer response.Body.Close()
	if err != nil {
		fmt.Println("请求出错了1")
	}

}
