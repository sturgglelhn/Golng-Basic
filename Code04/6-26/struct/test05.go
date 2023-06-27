package main

import "fmt"

// MyInt 将int定义为自定义MyInt类型
type MyInt2 int

// SayHello为MyInt添加一个SayHello的方法
func (m MyInt2) SayHello() {
	fmt.Println("Hello,我是一个int")
}

// Person结构体Person类型，结构体允许其成员字段在生命时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
type Person2 struct {
	string
	int
}

// Address 地址结构体
type Address struct {
	Province string
	City     string
}

// User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address //匿名结构体
}

func main() {
	/*任意添加方法*/
	var m1 MyInt2
	m1.SayHello()
	m1 = 100
	fmt.Printf("%#v %T\n", m1, m1)

	/*结构体的匿名字段*/
	p1 := Person2{
		"amo",
		20,
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.string, p1.int)

	fmt.Println("------------")

	/*嵌套结构体*/
	user1 := User{
		Name:   "Woman",
		Gender: "女",
		Address: Address{
			Province: "湖北",
			City:     "荆州",
		},
	}
	fmt.Printf("user1=%#v\n", user1)

	/*
		嵌套匿名结构体：当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。
	*/
	var user2 User
	user2.Name = "man"
	user2.Gender = "男"
	user2.Address.Province = "上海" //通过匿名结构体.字段名访问
	user2.City = "松江"             //直接访问匿名结构体的字段名
	fmt.Printf("user2=%#v\n", user2)
}
