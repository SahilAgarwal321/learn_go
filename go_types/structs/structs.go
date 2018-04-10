package main

import "fmt"

type vertex struct {
	X int
	Y int
	Z string
}

func basicStructUseWithoutDeclaration() {
	fmt.Println(vertex{0, 0, "Zeroth"})
}

func structDeclarationExamples() {
	var structExample1 vertex = vertex{1, 1, "First"} // has type vertex
	fmt.Println("structExample1", structExample1)

	var structExample2 = vertex{2, 2, "Second"} // has type vertex
	fmt.Println("structExample2", structExample2)

	structExample3 := vertex{3, 3, "Third"} // has type vertex
	fmt.Println("structExample3", structExample3)

	var structExample4 = vertex{X: 4, Z: "implicit"} // Y:0 is implicit
	fmt.Println("structExample4", structExample4)

	var structExample5 = vertex{} // X:0, Y:0, Z:''
	fmt.Println("structExample5", structExample5)

}

func structExampleChangingValues() {
	var structExample = vertex{2, 2, "Second"}

	structExample.Z = "Changed!"
	fmt.Println("Values of structExample are:")
	fmt.Print("X: ", structExample.X, "\tY: ", structExample.Y, "\tZ: ", structExample.Z, "\n")
}

func structExampleUsingPointers() {
	structExample := vertex{3, 3, "Third"}
	pointerToStructExample := &structExample // has type *vertex

	// *pointerToStructExample.Z = "Changed using a pointer"
	// Above gives error
	pointerToStructExample.Z = "Changed using a pointer"
	fmt.Println(structExample)
	(*pointerToStructExample).Z = "Changed using a pointer differently"
	fmt.Println(structExample)

	var p = &vertex{1, 2, "pointers"}
	fmt.Println(p)
	fmt.Println(*p)
}

func main() {
	basicStructUseWithoutDeclaration()
	fmt.Println("----------------------------------------------------------")
	structDeclarationExamples()
	fmt.Println("----------------------------------------------------------")
	structExampleChangingValues()
	fmt.Println("----------------------------------------------------------")
	structExampleUsingPointers()
}
