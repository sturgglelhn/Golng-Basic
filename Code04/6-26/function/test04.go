package main

func test04() (int, int) {
	return 1, 2
}

func add2(x, y int) int {
	return x + y
}

func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}
	return x
}

func main() {
	println(add2(test04()))
	println(sum(test04()))
}
