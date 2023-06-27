package main

import "fmt"

// 类型定义
type NewInt int

// 类型别名
type MyInt = int

func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)
}
