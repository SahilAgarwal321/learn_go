package main

import "fmt"

func basicSlices() {
	arrayExample := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println(arrayExample)

	var slicedArray []int = arrayExample[1:4]
	fmt.Println(slicedArray)

	var slicedArray2 []int = arrayExample[1:]
	fmt.Println(slicedArray2)

	var slicedArray3 []int = arrayExample[:4]
	fmt.Println(slicedArray3)

	var slicedArray4 []int = arrayExample[:]
	fmt.Println(slicedArray4)
}

// Slices are like references to arrays
// A slice does not store any data, it just describes a section of an underlying array.
// Changing the elements of a slice modifies the corresponding elements of its underlying array.
// Other slices that share the same underlying array will see those changes.
func slicesAsPointers() {

	arrayExample := [4]string{
		"1st",
		"2nd",
		"3rd",
		"4th",
	}
	fmt.Println(arrayExample)

	slicedArray := arrayExample[0:2]  // elements 0,1
	slicedArray2 := arrayExample[1:3] // elements 1,2
	fmt.Println(slicedArray, slicedArray2)

	slicedArray2[0] = "Changed!"
	fmt.Println(arrayExample)
	fmt.Println(slicedArray, slicedArray2)

}

func main() {
	basicSlices()
	fmt.Println("---------------------------------------------")
	slicesAsPointers()
}
