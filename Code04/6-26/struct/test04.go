package main

import "fmt"

/*
	方法和接受者
	GO语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型的变量叫做接受者（Receiver）。
	接受者的概念就类似于其他语言中的this或者self。

	方法的定义格式如下：
		func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
			函数体
		}
	1、接收者变量：接受者中的参数变量在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。例如，Person类型的接受者变量应该命名为p，Connector类型的接收者变量应该命名为c等。
	2、接收者类型：接收者类型和参数类似，可以时指针类型和非指针类型。
	3、方法名、参数列表、返回参数：具体格式和函数定义相同。
*/

// Person结构体
type Person struct {
	name string
	age  int8
}

// NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好GO语言！%d\n", p.name, p.age)
}

func main() {
	p1 := NewPerson("测试", 25)
	p1.Dream()

	/*
		指针类型的接收者
	*/
	p2 := NewPerson("测试2", 25)
	fmt.Println(p2.age)
	p2.SetAge(30)
	fmt.Println(p2.age)

	/*
		值类型的接收者：当方法作用于值类型接收者时，GO语言会在代码运行时将接收者的值复制一份。
		在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
	*/
	p3 := NewPerson("测试3", 25)
	p3.Dream()
	fmt.Println(p3.age)
	p3.SegAge2(30)
	fmt.Println(p3.age)
}

// SegAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SegAge2 设置p的年龄
// 使用值接收者
func (p Person) SegAge2(newAge int8) {
	p.age = newAge
}

/*
什么时候应该使用指针类型接收者？
1、需要修改接收者中的值
2、接收者是拷贝代价比较大的大对象
3、保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/
