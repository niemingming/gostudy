package main

import (
	"fmt"
)

//go语言学习，变量声明，常量声明以及数组
var a int
var b = 6 //常量声明
const c = "const"
const (
	d = iota
	e
	f
	g
)

//字符串练习，引入strings包，用于字符串的处理
var h = "小妞"
var a1 = [5]int{1, 2, 3, 4, 5}

func main() {
	//数组声明
	k := [5]string{"张三", "李四", "王五"}
	k[3] = "赵二"
	k[4] = "小六"
	//key-value数组
	l := [3]string{1: "张三", 2: "李四", 0: "王五"}

	//切片
	var s1 = l[1:3]
	var s2 = []int{1, 2, 3}
	//创建非空map
	var m1 = make(map[string]string)
	m1["name"] = "张三"
	m2 := map[string]string{
		"name": "李四",
		"age":  "18",
	}
	fmt.Println("变量声明a=", a, ";b=", b)
	fmt.Println("常量声明c=", c, ";d=", d, ";e=", e, ";f=", f, ";g=", g)
	fmt.Println("字符串：h[0]=", h[0])
	for i := 0; i < len(h); i++ {
		fmt.Println("h[", i, "]=", h[i])
	}
	fmt.Println("数组：", k)
	for m, n := range l {
		fmt.Println(m, "=", n)
	}
	fmt.Println("切片：", s2)
	fmt.Println(s1)
	fmt.Println("map:", m1, "m2=", m2)

	fmt.Println("pause")
}
