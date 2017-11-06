package main

import (
	"fmt"
)

//处理异常
func cs(i int, j int) int {
	defer func() {

		if err := recover(); err != nil {
			fmt.Println("出错了：", err)
		}
	}()
	k := i / j
	fmt.Println("计算完成：")
	return k
}

func main() {
	fmt.Println(cs(2, 1))
	fmt.Println(cs(3, 0))
}
