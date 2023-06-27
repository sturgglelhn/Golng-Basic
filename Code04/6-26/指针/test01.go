package main

import "fmt"

/*
区别于C/C++中的指针，GO语言中的指针不能进行便宜和运算，是安全指针。
要搞明白GO语言中的指针需要先知道3个概念：指针地址、指针类型和指针取值。
只需要记住两个符号：&（取地址）和*（根据地址取值）
取变量指针的语法如下：

	ptr := &v	//v的类型为T

其中：

	v:代表取地址的变量，类型为T
	ptr:用于接收地址的变量，ptr的类型就位*T，称做T的指针类型。*代表指针。

总结：

	取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址去除地址指向的值。
	变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：\
		1、对变量进行取地址（&）操作，可以获得这个变量的指针变量。
		2、指针变量的值是指针地址。
		3、对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
*/
func main() {
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中

	/*指针地址和指针类型*/
	fmt.Printf("a:%d ptr:%p\n", a, &a) //a:10 ptr:0xc00001c098
	fmt.Printf("b:%p type:%T\n", b, b) //b:0xc00001c098 type:*int
	fmt.Println(&b)                    //0xc00000a028

	/*指针取值*/
	c := *b                         // 指针取值（根据指针取内存取值）
	fmt.Printf("type of c:%T\n", c) //type of c:int
	fmt.Printf("type of c:%v\n", c) //type of c:10

	/*指针传值示例：*/
	a2 := 10
	modify1(a2)
	fmt.Println(a2) //10
	modify2(&a2)
	fmt.Println(a2) //100
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}
