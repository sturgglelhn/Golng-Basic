package main

import "fmt"

/*
1、对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然。
2、对于方法（如struct），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以。
*/

// 1.普通函数
// 接收值类型参数的函数
func valueIntTest(a int) int {
	return a + 10
}

// 接收指之好着呢类型参数的函数
func pointerIntTest(a *int) int {
	return *a + 10
}

func structTestValue() {
	a := 2
	fmt.Println("valueIntTest:", valueIntTest(a))
	//函数的参数为值类型，则不能直接将指针作为参数传递
	//fmt.Println

	b := 5
	fmt.Println("pointerIntTest:", pointerIntTest(&b))
}

// 2.方法
type PersonD struct {
	id   int
	name string
}

// 接收者为值类型
func (p PersonD) valueShowName() {
	fmt.Println(p.name)
}

// 接收者为指针类型
func (p *PersonD) pointShowName() {
	fmt.Println(p.name)
}

func structTestFunc() {
	// 值类型调用 方法
	personValue := PersonD{101, "hello world"}
	personValue.valueShowName()
	personValue.pointShowName()

	//指针类型调用方法
	personPointer := &PersonD{102, "hello golang"}
	personPointer.valueShowName()
	personPointer.pointShowName()

	//与普通函数不同，接收者为指针类型和值类型的方法，指针类型和值类型的变量均可相互调用
}

func main() {
	structTestValue()
	structTestFunc()
}
