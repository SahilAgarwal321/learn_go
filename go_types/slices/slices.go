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

	sliceWithoutLength := []string{"Slice: with no", "length", "specified"}
	fmt.Println(sliceWithoutLength)

	sliceWithoutLengthUsingStruct := []struct {
		i int
		b bool
	}{
		{5, true},
		{3, false},
		{764, true},
	}

	fmt.Println(sliceWithoutLengthUsingStruct)

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

// Function showcasing len() and cap() on slices.
func sliceLengthAndCapacity() {
	arrayExample := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%s\t len=%d cap=%d\t %v\n", "arrayExample", len(arrayExample), cap(arrayExample), arrayExample)
	// printSlice("arrayExample", arrayExample) // gives error since in func def its (s []int) and not (s [6]int)

	var slicedArray []int = arrayExample[1:4]
	printSlice("slicedArray", slicedArray)

	var slicedArray2 []int = arrayExample[1:]
	printSlice("slicedArray2", slicedArray2)

	printSlice("slicedArray2", slicedArray2[:3])
	printSlice("arrayExample", arrayExample[:3])

	var defaultExample [6]int
	var nilExample []int

	fmt.Printf("%s\t len=%d cap=%d\t %v\n", "defaultExample", len(defaultExample), cap(defaultExample), defaultExample)
	printSlice("nilExample", nilExample)

}

// Prints slice name, it's length, capacity, values
func printSlice(s string, x []int) {
	fmt.Printf("%s\t len=%d cap=%d\t %v\n",
		s, len(x), cap(x), x)
}

func main() {
	basicSlices()
	fmt.Println("---------------------------------------------")
	slicesAsPointers()
	fmt.Println("---------------------------------------------")
	sliceLengthAndCapacity()
}
