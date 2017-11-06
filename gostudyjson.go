//go语言操作json
package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

//声明类型，json转换
type user struct {
	Name     string
	Password string
}

//声明一个包含users的数组
type users struct {
	Cuser []user
}

//说明，结构体只有大写字母开头的属性才会被json识别
func main() {
	var us users
	var str = `{"cuser":[{"name":"zhangsan","password":"12345"},{"name":"lisi","password":"12345"}]}`
	json.Unmarshal([]byte(str), &us)
	fmt.Println(len(us.Cuser))

	//对象转为json
	us.Cuser = append(us.Cuser, user{Name: "wangwu", Password: "332"})
	b, err := json.Marshal(us)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	//base64编码
	var s = "hello"
	bs := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(bs)
	//反编译
	s1, err := base64.StdEncoding.DecodeString(string(bs))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(s1))
}
