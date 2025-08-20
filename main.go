package main

import "fmt"

func main() {
	sum := sum(3, 5)
	fmt.Println("The sum is:", sum, multiply(4, 6), divide(10, 2))
}

func sum(a int, b int) int {
	return a + b
}
func multiply(a int, b int) int {
	return a * b
}
func divide(a int, b int) float64 {
	if b == 0 {
		return 0 // Avoid division by zero
	}
	return float64(a) / float64(b)
}
