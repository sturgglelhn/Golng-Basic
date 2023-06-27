package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}

func main() {
	/*
		结构体实例化
	*/
	var p1 person
	p1.name = "pprof.cn"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)

	/*
		匿名结构体
	*/
	var user struct {
		Name string
		Age  int
	}

	user.Name = "Gandolf"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	/*
		创建指针类型结构体，我们可以通过是new关键字对结构体进行实例化，得到的是结构体的地址：
						p2=&main.person{name:"测试", city:"北京", age:18}
		需要注意的是在go语言中支持对结构体指针直接使用.来访问结构体的成员。
	*/
	var p2 = new(person)
	p2.name = "测试"
	p2.age = 18
	p2.city = "北京"
	fmt.Printf("p2=%#v\n", p2)

	/*
		取结构体的地址实例化，使用&对结构体去地址操作相对于该结构体类型进行了一次new实例化操作。
		输出结果：
			*main.person
			p3=&main.person{name:"", city:"", age:0}
			p3=&main.person{name:"博客", city:"成都", age:30}
		p3.name="博客"其实在底层是(*p3).name="博客"，这是go语言帮我们实现的语法糖。
	*/
	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3=%#v\n", p3)
	p3.name = "博客"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3)

	/*
		结构体初始化
		输出结果：
			p4=main.person{name:"", city:"", age:0}
	*/
	var p4 person
	fmt.Printf("p4=%#v\n", p4)

	/*
		使用键值对初始化，在使用健值对对结构体初始化时，键对应结构体的字段，值对应字段的初始值。

	*/
	p5 := person{
		name: "lihaonan",
		city: "上海",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"lihaonan", city:"上海", age:18}

	// 也可以对结构体指针进行健值对初始化，例如：
	p6 := &person{
		name: "mao",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"mao", city:"北京", age:18}

	//当某些字段没有初始值的时候，该字段也可以不写。此时，没有置顶初始值的字段的值就是该字段类型的零值。
	p7 := &person{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}

	/*
		使用值的列表初始化
		初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值；
		使用这种格式初始化时，需要注意：
			1、必须初始化结构的所有字段
			2、初始值的填充顺序必须与字段的结构体中的声明顺序一直
			3、该方式不能和键值初始化方式混用
	*/
	p8 := &person{
		"zi",
		"深圳",
		18,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"zi", city:"深圳", age:18}

	/*
		构造函数
	*/
	p9 := newPerson("li", "测试", 88)
	fmt.Printf("p9=%#v\n", p9)

	//结构体中的顺序是连续的
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
	/*
		n.a 0xc0000a6178
		n.b 0xc0000a6179
		n.c 0xc0000a617a
		n.d 0xc0000a617b
	*/

}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

/*
我们可以通过.来访问结构体的字段（成员变量），例如p1.name和p1.age等

输出结果：
p1={pprof.cn 北京 18}
p1=main.person{name:"pprof.cn", city:"北京", age:18}
*/

/*
输出结果：
struct { Name string; Age int }{Name:"Gandolf", Age:18}

*/
