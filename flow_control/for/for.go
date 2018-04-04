package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	// i:=0 is the init part of the for loop
	// i < 10 is the condition part of the for loop
	// i++ is the post part of the for loop
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	for {
		fmt.Println("This is an infinite loop, but breaks due to break stmt!")
		break
	}
}
