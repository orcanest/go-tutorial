package main

import "fmt"

func main() {

	var (
		age         int     = 20
		name, movie string  = "jack", "Good fells"
		score       float64 = 7
	)

	fmt.Println(name, "is a good student")
	fmt.Println(name, "is", age, "years old")
	fmt.Println(movie, name, "favorite movie score is", score)
}
