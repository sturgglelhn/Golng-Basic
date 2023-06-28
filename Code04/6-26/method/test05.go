package main

import "fmt"

/*Golang匿名字段：可以像字段成员那样访问匿名字段方法*/
type User2 struct {
	id   int
	name string
}

type Manager struct {
	User2
	title string
}

func (self *User2) ToString() string {
	return fmt.Sprintf("User2: %p,%v", self, self)
}

func (self *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p,%v", self, self)
}

func main() {
	m := Manager{User2{1, "Tom"}, "administrator"}
	//fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
	fmt.Println(m.User2.ToString())
}

/*
	通过匿名字段，可获得和继承类似的复用能力。
*/
