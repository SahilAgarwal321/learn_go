package main

import (
	"fmt"

	"github.com/SahilAgarwal321/learn_go/basics/stringutil"
)

func main() {
	fmt.Printf("Hello, World.")
	fmt.Printf("\n")
	fmt.Printf(stringutil.Reverse("\n!oG ,olleH"))

	fmt.Printf("\nThe Second")
	fmt.Println("\nChance") // Does not add \n to the start. Only adds it to end.
	fmt.Printf("End")
}
