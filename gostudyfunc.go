package main

//函数练习
import (
	"fmt"
)

//声明类型
type Circle struct {
	radius float64 //定义圆，半径属性
}

//接口练习，声明person，实现吃、喝、跑
type person struct {
	name string
	age  int
}

//实现方法
func (p person) eat() {
	fmt.Println("吃饭！")
}
func (p person) drink() {
	fmt.Println("喝水！")
}
func (p person) run() {
	fmt.Println("跑步！")
}

//声明接口
type eatinf interface {
	eat()
}
type drinkinf interface {
	drink()
}
type move interface {
	run()
	eatinf
	drinkinf
}

//主函数
func main() {
	//函数测试
	test(55)
	c := Circle{10}
	fmt.Println("面积：", c.getArea())
	//接口实现
	var m1 move
	m1 = person{"张三", 12}
	m1.eat()
	m1.drink()

	var m2 eatinf
	m2 = person{"李四", 13}
	m2.eat()
}

//普通函数
func test(i int) {
	fmt.Println("println", i)
}

//定义circle函数
func (c *Circle) getArea() float64 {
	return c.radius * c.radius * 3.14
}
