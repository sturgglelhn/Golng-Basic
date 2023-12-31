package main

/*
方法定义
结构定义：
    func (recevier type) methodName(参数列表)(返回值列表){}
    参数和返回值可以省略
·只能为当前包内命名类型定义方法。
·参数receiver可任意命名。如方法中未曾使用，可省略参数名。
·参数receiver类型可以是T或*T。基类型T不能是接口或指针。
·不支持方法重载，receiver只是参数签名组成部分。
·可用实例value或pointer调用全部方法，编译器自动转换
一个方法就是一个包含了接受者的函数，接受者可以是明明类型或者结构体类型的一个值或者是一个指针。
*/

type Test struct {
}

// 无参数、无返回值
func (t Test) method0() {
}

// 单参数、无返回值
func (t Test) method1(i int) {}

// 多参数、无返回值
func (t Test) method2(x, y int) {}

// 无参数、单返回值
func (t Test) method3() (i int) {
	return
}

// 多参数、多返回值
func (t Test) method4(x, y int) (z int, err error) {
	return
}

// 无参数、无返回值
func (t *Test) method5() {

}

// 单参数、无返回值
func (t *Test) method6(i int) {

}

// 多参数、无返回值
func (t *Test) method7(x, y int) {

}

// 无参数、单返回值
func (t *Test) method8() (i int) {
	return
}

// 多参数、多返回值
func (t *Test) method9(x, y int) (z int, err error) {
	return
}

func main() {}
