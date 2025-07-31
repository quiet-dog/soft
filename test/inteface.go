package main

import "fmt"

type Test interface {
	A()
	B()
}

type Demo struct {
}

func (d *Demo) A() {
	fmt.Println("2")
}

func main() {
	var t Test   // 定义接口类型变量
	d := &Demo{} // 创建 Demo 的实例，注意是指针类型
	t = d        // 将 Demo 实例赋值给接口变量
	t.A()        // 调
}
