package main

//通道channel和线程routine练习
import (
	"fmt"
)

//定义无缓存通道
var c = make(chan int)
var a string

func read() {
	a = "read before"
	<-c
	a = "read after"
}

func main() {
	go read()
	fmt.Println(a)
	c <- 20 //写入
	fmt.Println(a)

}
