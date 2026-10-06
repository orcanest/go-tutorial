package main

import "fmt"

func main() {
	a, b := 5, 7
	a, b = b, a
	fmt.Println(a, b)
}
