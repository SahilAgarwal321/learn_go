package main

import "fmt"

func main() {
	var a [2]string
	// An array's length is part of its type, so arrays cannot be resized.
	a[0] = "Hello"
	a[1] = "World"
	// a[2] = "Gives error when uncommented"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	exampleDefaultValues := [6]int{2, 3, 5, 7}
	fmt.Println(exampleDefaultValues)
}
