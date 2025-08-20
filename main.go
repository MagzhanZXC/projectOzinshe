package main

import "fmt"

func main() {
	sum := sum(3, 5)
	fmt.Println("The sum is:", sum)
}

func sum(a int, b int) int {
	return a + b
}
