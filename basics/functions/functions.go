// Demonstrates different syntaxes in which functions can be defined.
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

// Showcasing multiple results/returns from one function
func swap(x, y int) (int, int) {
	return y, x
}

// Showcasing multiple results/returns from one function
func swapStrings(x, y string) (string, string) {
	return y, x
}

// Showcasing named results/returns thingy.
// Splits a float number into two and returns both the values, as defined in the return (x, y float 32)
func splitFloat(sum float32) (x, y float32) {
	x = sum * 1 / 2
	y = sum - x
	return
}

// Showcasing named results/returns thingy.
// Splits an int number into two and returns both the values x, y as named in return
// Note that abc is not returned
func splitInt(sum int) (x, y int) {
	x = sum * 1 / 2
	var abc = 45
	abc = abc
	y = sum - x
	return
}

func main() {
	fmt.Println("Function ``add()`` adds numbers 42, 13: ", add(42, 13))
	fmt.Println(add2(42, 13))
	var x, y int = 23, 54
	x, y = swap(x, y)
	fmt.Println("swap(23, 54) returns: ", x, y)
	fmt.Println(swapStrings("How are you?", "Hi"))
}
