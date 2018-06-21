// Exploring variables, how to define them and their scopes
package main

import "fmt"

func returnInputPlusFive(x int) int {
	var abc = x + 5
	return abc
}

var defaultBoolVariableExample bool
var defaultIntVariableExample int
var defaultStringVariableExample string
var python, java bool = true, true

func main() {
	var intVariableExample, stringVariableExample = 24, "okay"
	shortVariableExample := 221
	fmt.Println("Default examples -")
	fmt.Println(defaultBoolVariableExample, defaultStringVariableExample, defaultIntVariableExample)
	fmt.Println("Other examples -")
	fmt.Println(python, java)
	fmt.Println(intVariableExample, stringVariableExample, shortVariableExample)
	// fmt.Printlin(abc) //gives error
}
