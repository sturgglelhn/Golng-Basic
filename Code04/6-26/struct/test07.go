package main

import (
	"fmt"
)

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

// Dog狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪~\n", d.name)
}

/*结构体的"继承"*/
func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}
	d1.wang()
	d1.move()
}

/*
	结构体字段的可见性
	结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）
*/
