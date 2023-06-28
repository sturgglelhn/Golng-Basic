package main

import "fmt"

type User3 struct {
	id   int
	name string
}

func (self *User3) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

func main() {
	u := User3{1, "Tom"}
	u.Test()

	mValue := u.Test
	mValue() // 隐式传递receiver

	mExpression := (*User3).Test
	mExpression(&u) //显示传递 receiver
}
