package main

import "fmt"

func main() {

	/*空指针*/
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%s\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}

	/*new和make*/
	var a3 *int
	a3 = new(int)
	*a3 = 100
	fmt.Println(*a3)

	var b map[string]int
	b["测试"] = 100
	fmt.Println(b)
}
