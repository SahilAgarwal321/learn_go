package main

import "fmt"

func forBasicExample() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	// i:=0 is the init part of the for loop
	// i < 10 is the condition part of the for loop
	// i++ is the post part of the for loop
	fmt.Println(sum)
}

func forWithoutPostPartTry() {
	sum := 0
	for i := 0; i < 10; {
		sum += i
		i = i + 1
	}
	// i:=0 is the init part of the for loop
	// i < 10 is the condition part of the for loop
	// i++ is the post part of the for loop
	fmt.Println(sum)
}

func forWhileExample() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func forInfiniteLoopExample() {
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
