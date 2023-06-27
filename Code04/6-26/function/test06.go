package main

import "fmt"

func main() {
	var wahtever [5]struct{}

	for i := range wahtever {
		defer fmt.Println(i)
	}
}
