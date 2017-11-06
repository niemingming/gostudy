package main

import (
	"fmt"
	"html/template" //模板页面
	"log"
	"net/http"
)

//go实现服务端表单处理。

//服务端监听方法一，监听根目录
func sayHello(w http.ResponseWriter, r *http.Request) {
	//初始化form
	r.ParseForm() //序列化表单
	//显示port、地址和schema
	fmt.Println("服务器信息")
	fmt.Println("url:", r.URL.Path, ";schema:", r.URL.Scheme)
	fmt.Println(r.Form)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Print("key:", k)
		fmt.Println("value:", v)
	}
	fmt.Fprintf(w, "hello box")
}

//登录方法
func login(w http.ResponseWriter, r *http.Request) {
	//判断方法
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html") //模板文件加载模板
		t.Execute(w, nil)
	} else {
		//读取数据
		fmt.Println("username:", r.PostFormValue("username"))
		fmt.Println("password", r.PostFormValue("password"))
	}
}

//执行方法
func main() {
	//注册监听方法
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	//启动服务
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("启动异常", err)
	}
}
