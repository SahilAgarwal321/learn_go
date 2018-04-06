// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
package main

import (
	"fmt"
	"math"
)

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
	defer fmt.Println("") // To get to a newline after loop in defer executes
	for i := 0; i < 10; i++ {
		defer fmt.Print(i, " ")
	}
	fmt.Println("done")
}

func experiment1WithDefer(x int) int {
	if x == 1 {
		fmt.Println("1")
		return 1
	}
	defer fmt.Println("defer stmt here")
	if x == 2 {
		fmt.Println("2")
		return 2
	}
	return 0
}

// TODO: Need to do this with a list.
// TODO: Make var x a list.
// TODO: Replace math.Min(x, 2) with a function that edits the list (remove/add an element or clear the list)
func experiment2WithDefer() {
	var x float64 = 5
	fmt.Println("1.  x =", x)
	defer fmt.Println("2.  x =", x)
	defer math.Min(x, 2)
	fmt.Println("3.  x =", x)
}

func main() {
	var x float64 = 5
	math.Min(x, 2)
	fmt.Println("------------ basicDeferExample -------------")
	basicDeferExample()
	fmt.Println("------------ basicStackedDefers -------------")
	basicStackedDefers()
	fmt.Println("------------ loopStackedDefers -------------")
	loopStackedDefers()
	fmt.Println("------------ experiment1WithDefer with input 1 -------------")
	experiment1WithDefer(1)
	fmt.Println("------------ experiment1WithDefer with input 2 -------------")
	experiment1WithDefer(2)
	fmt.Println("------------ experiment1WithDefer with input 3 -------------")
	experiment1WithDefer(3)
	fmt.Println("------------ experiment2WithDefer -------------")
	experiment2WithDefer()
}
