package main

import "fmt"

type student2 struct {
	id   int
	name string
	age  int
}

func demo(ce []student2) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	ce = append(ce, student2{3, "xiaowang", 56})
	//return ce
}

func main() {
	/*var ce []student2
	ce = []student2{
		{1, "xiaoming", 22},
		{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)*/

	ce := make(map[int]student2)
	ce[1] = student2{1, "xiaolizi", 22}
	ce[2] = student2{2, "wamg", 23}
	fmt.Println(ce)
	delete(ce, 1)
	fmt.Println(ce)
}
