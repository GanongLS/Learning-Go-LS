package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switchConstruct()
	switchIntendToBreakLoop()
	switchBreakLoop()
	blankSwitch()
	switchDoAndDont()
	switchFizzBuss()
	// gotoRules()
	goodGoto()
}

func goodGoto() {
	fmt.Println("Good Goto")

	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)
	fmt.Println("")

}

// func gotoRules() {
// 	a := 10
// 	goto skip // tidak boleh jump karena ngelewatin variable declaration b
// 	b := 20
// skip:
// 	c := 30
// 	fmt.Println(a, b, c)
// 	if c > a {
// 		goto inner // tidak boleh jumpe ke inner scope
// 	}

// 	if a < b {
// 	inner:
// 		fmt.Println("a is less than b")
// 	}
// }

func switchFizzBuss() {
	fmt.Println("Switch Fizz Buss")

	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}

	fmt.Println("")

}

func switchDoAndDont() {

	fmt.Println("Do and Don't")

	// Blank switches are pretty cool, but don’t overdo them. If you find that you have written a blank switch where all your cases are equality comparisons against the same variable:
	a := 3
	// DON'T
	switch {
	case a == 2:
		fmt.Println("a is 2")
	case a == 3:
		fmt.Println("a is 3")
	case a == 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("a is ", a)
	}

	// you should replace it with an expression switch statement:
	// DO

	switch a {
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	case 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("a is ", a)
	}

	fmt.Println("")

}

func blankSwitch() {
	fmt.Println("Blank Switch")

	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}
	fmt.Println("")

}

func switchBreakLoop() {
	fmt.Println("Break the loop")

loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}
	fmt.Println("")

}

func switchIntendToBreakLoop() {
	fmt.Println("intend to break the loop but fail")

	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break
		default:
			fmt.Println(i, "is boring")
		}
	}
	fmt.Println("")

}

func switchConstruct() {
	fmt.Println("Switch Construct")
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"} // empty case "octopus" or "gopher" do nothing.

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	fmt.Println("")

}
