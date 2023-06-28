package main

import "fmt"

type Data struct {
	x int
}

func (self Data) ValueTest() {
	fmt.Printf("Value:%p\n", &self)
}

func (self *Data) PointerTest() {
	fmt.Printf("Pointer:%p\n", self)
}

func main() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)

	d.ValueTest()
	d.PointerTest()

	p.ValueTest()
	p.PointerTest()
}

/*
Data: 0xc00001c098
Value:0xc00001c0b8
Pointer:0xc00001c098
Value:0xc00001c0d0
Pointer:0xc00001c098
*/
