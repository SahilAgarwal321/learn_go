// Go only runs selected switch case. Break stmt is built in and thus, not needed.
package main

import (
	"fmt"
	"runtime"
	"time"
)

func basicSwitchCaseExample() {
	var x = 12
	switch x {
	case 12:
		fmt.Println("12")
	case 31:
		fmt.Println("31")
	}
}

func basicSwitchWithDefaultCaseExample() {
	var x = 123
	switch x {
	case 12:
		fmt.Println("12")
	case 31:
		fmt.Println("31")
	default:
		fmt.Println("Something else.")
	}
}

func switchCaseWithInitPart() {
	fmt.Print("Go runs on ")
	// os := runtime.GOOS is the init part of the switch stmt
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	// fmt.Printf("%s.", os)
	// Uncommenting abv stmt gives error. Scope of `os` is inside `switch` stmt.
}

// Above examples only had equivalence check condition for cases.
// This one also is equivalence check, but checks time.Saturday == today + 0 / 1 / 2
func basicSwitchCaseExample2() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

//  This construct can be a clean way to write long if-then-else chains.
func switchWithNoConditionPart() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	basicSwitchCaseExample()
	basicSwitchWithDefaultCaseExample()
	switchCaseWithInitPart()
	basicSwitchCaseExample2()
	switchWithNoConditionPart()
}
