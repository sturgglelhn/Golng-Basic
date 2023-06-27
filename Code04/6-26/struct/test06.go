package main

import "fmt"

// Address 地址结构体
type Address2 struct {
	Province   string
	City       string
	CreateTime string
}

// Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

// User 用户结构体
type User2 struct {
	Name   string
	Gender string
	Address2
	Email
}

/*
嵌套结构体的字段名冲突：当访问结构体成员时会现在结构体中查找该字段，找不到再去匿名结构体中查找。
*/
func main() {
	var user3 User2
	user3.Name = "pprof"
	user3.Gender = "女"

	user3.Address2.CreateTime = "2000"
	user3.Email.CreateTime = "2000"

	fmt.Printf("user3=%#v\n", user3)
}
