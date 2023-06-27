package main

import "fmt"

/*
函数参数：
值传递：是指将实际参数复制一份传递到函数中，对参数修改将不会影响到实际参数。
引用传递：是指在调用函数是将实际参数的地址传递到函数中，对参数修改会影响到实际参数。
*/

/*定义相互交换值的函数*/
func swap(x, y *int) {
	var temp int

	temp = *x // 保存x的值
	*x = *y   // 将y值赋值x
	*y = temp //将temp值赋给y
}

func test3(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}

	return fmt.Sprintf(s, x)
}

func main() {
	var a, b int = 1, 2
	/*
		调用swap()函数
		&a 指向a指针，a变量的地址
		&b 指向b指针，b变量的地址
	*/
	swap(&a, &b)
	fmt.Println(a, b)

	println(test3("sum:%d", 1, 2, 3))
}
