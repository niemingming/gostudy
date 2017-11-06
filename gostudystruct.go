package main

//练习struct声明和使用
//声明学生、老师、班级类型

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

type teacher struct {
	name string
	kc   string
}
type class struct {
	student
	tea   teacher
	index int
}

func main() {
	var s1 student
	s1.name = "张三"
	s1.age = 12
	t1 := &teacher{"杰伦", "音乐"}
	c1 := class{student: student{"李四", 112}, tea: teacher{"那英", "历史"}, index: 12}

	fmt.Println(s1)
	fmt.Println(t1)
	fmt.Println(c1)

}
