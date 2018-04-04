package main

import "fmt"

func forBasicExample() {
	// i:=0 is the init part of the for loop
	// i < 10 is the condition part of the for loop
	// i++ is the post part of the for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forWithoutPostPartTry() {
	// i:=0 is the init part of the for loop
	// i < 10 is the condition part of the for loop
	sum := 0
	for i := 0; i < 10; {
		sum += i
		i = i + 1
	}
	fmt.Println(sum)
}

func forWhileExample() {
	// Removing init and post parts of the for loop
	// sum < 1000 is the condition part of the for loop
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func forInfiniteLoopExample() {
	// Removing init, condition, post. Creates an infinite loop.
	for {
		fmt.Println("This is an infinite loop, but breaks due to break stmt!")
		break
	}
}

func main() {
	fmt.Print("forBasicExample: ")
	forBasicExample()
	fmt.Print("forWhileExample: ")
	forWhileExample()
	fmt.Print("forInfiniteLoopExample: ")
	forInfiniteLoopExample()
	fmt.Print("forWithoutPostPartTry: ")
	forWithoutPostPartTry()
}
