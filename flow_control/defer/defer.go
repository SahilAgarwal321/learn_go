// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
package main

import "fmt"

func basicDeferExample() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// Deferred function calls are pushed onto a stack.
// When a function returns, its deferred calls are executed in last-in-first-out order.
func basicStackedDefers() {
	fmt.Println("0")
	defer fmt.Println("1")
	fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")
}

func loopStackedDefers() {

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	fmt.Println("basicDeferExample")
	basicDeferExample()
	fmt.Println("basicStackedDefers")
	basicStackedDefers()
	fmt.Println("loopStackedDefers")
	loopStackedDefers()
}
