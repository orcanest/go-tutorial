package main

import "fmt"

func main() {
	// TODO: Calculate and print the weight on Mars
	var earthWeight int = 80
	const gravityRatio = 0.3783
	marsWeight := float32(earthWeight) * gravityRatio
	fmt.Println(marsWeight)
}