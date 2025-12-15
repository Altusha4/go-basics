package main

import "fmt"

func multiply(a int, b int) int {
	return a * b
}

func main() {
	result := multiply(3, 4)
	fmt.Println(result)
}
