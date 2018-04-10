package main

import "fmt"

type vertex struct {
	X int
	Y int
	Z string
}

func basicStructExampleWithoutDeclaration() {
	fmt.Println(vertex{0, 0, "Zeroth"})
}

func basicStructExample() {
	var structExample vertex = vertex{1, 1, "First"}
	fmt.Println(structExample)
}

func structExampleChangingValues() {
	var structExample = vertex{2, 2, "Second"}
	fmt.Println(structExample)

	structExample.Z = "Changed!"
	fmt.Println("Values of structExample are:")
	fmt.Print("X: ", structExample.X, "\tY: ", structExample.Y, "\tZ: ", structExample.Z, "\n")
}

func structExampleUsingPointers() {
	structExample := vertex{3, 3, "Third"}
	fmt.Println(structExample)

	// Now using pointers with structs -

	pointerToStructExample1 := &structExample

	// *pointerToStructExample1.Z = "Changed using a pointer"
	// Above gives error
	pointerToStructExample1.Z = "Changed using a pointer"
	fmt.Println(structExample)
	(*pointerToStructExample1).Z = "Changed using a pointer differently"
	fmt.Println(structExample)

}

func main() {
	basicStructExampleWithoutDeclaration()
	fmt.Println("----------------------------------------------------------")
	basicStructExample()
	fmt.Println("----------------------------------------------------------")
	structExampleChangingValues()
	fmt.Println("----------------------------------------------------------")
	structExampleUsingPointers()

}
