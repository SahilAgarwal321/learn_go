package main

import "fmt"

func ifBasicExample() {
	var x = 21
	if 5 < x {
		fmt.Println("5 is less than", x)
	} else {
		fmt.Println("5 is not less than", x)
	}
}

// Showcasing if, else-if, else stmts. Vary value of `x` and run.
func ifElseExample() {
	var x = 21
	if i := 21; i < x {
		fmt.Println(i, "is less than", x)
	} else if i == x {
		fmt.Println(i, "is equal to", x)
	} else {
		fmt.Println(i, "is greater than", x)
	}
	// fmt.Print(i)
	// Uncommenting abv stmt gives error. Scope of `i` is inside `if` and it's `else` stmts.
}

func main() {
	ifBasicExample()
	ifElseExample()
}
