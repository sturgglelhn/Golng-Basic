package main

import "fmt"

/*
new

new是一个内置的函数，它的函数签名如下：
	func new(Type) *Type
其中：
	1、Type表示类型，new函数只接受一个参数，这个参数是一个类型。
	2、*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。

new与make的区别
1、二者都是用来做内存分配的
2、make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
3、而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/

func main() {

	/*new*/
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) //*int
	fmt.Printf("%T\n", b) //*int
	fmt.Println(*a)       //0
	fmt.Println(*b)       //false

	/*make*/
	var b2 map[string]int
	b2 = make(map[string]int, 20)
	b2["测试"] = 100
	fmt.Println(b2)

}
