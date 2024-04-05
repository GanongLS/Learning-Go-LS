package main

import (
	"fmt"
)

func main() {
	standardFor()
	infiniteFor()
	forRange()
	iteratingOverString()
	forRangeValueIsCopied()
	labelingYourFor()
	whatToChoose()
}

func whatToChoose() {
	fmt.Println("What to Choose")

	// First the for-range loop:
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		if i == 0 {
			continue
		}
		if i == len(evenVals)-1 {
			break
		}
		fmt.Println(i, v)
	}
	// And here’s the same code, with a standard for loop:
	// evenVals := []int{2, 4, 6, 8, 10}
	for i := 1; i < len(evenVals)-1; i++ { // standard loop best for accessing value from slice while not all the value being accesed.
		fmt.Println(i, evenVals[i])
	}
	fmt.Println("")

}

func labelingYourFor() {
	fmt.Println("Labeling your for")

	samples := []string{"hello", "apple_π!"}

	for _, sample := range samples {
		for i, r := range sample {
			if r == 'l' {
				continue
			}
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}

	fmt.Println("")

}

func forRangeValueIsCopied() {
	fmt.Println("For Range value is Copied")

	// Modifying the value doesn’t modify the source
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)
	fmt.Println("")

}

func iteratingOverString() {
	fmt.Println("Iterating Over String")

	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample { // iterating over string return the runes as value.
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
	fmt.Println("")

}

func forRange() {
	fmt.Println("For Range")

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	for _, v := range evenVals { // ignoring the index
		fmt.Println(v)
	}

	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames { // ignoring the values
		fmt.Println(k)
	}

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println("Map", m)

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
	fmt.Println("")

}

func infiniteFor() {
	fmt.Println("Initiate For")
	x := 0
	for {
		fmt.Println("Hello")
		x++
		// and break
		if x == 100 {
			break
		}
	}

	// nested if
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// and continue
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	// my version
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("")
}

func standardFor() {
	fmt.Println("Standard For Construct")
	// var def; condition; increment
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Go allows you to leave out one or more of the three parts of the for statement.
	i := 0
	for ; i < 10; i++ { // menskip definition i,
		fmt.Println(i)
	}

	// or you’ll leave off the increment because you have a more complicated increment rule inside the loop:
	for i := 0; i < 10; {
		fmt.Println(i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}

	// The Condition-Only for Statement
	// When you leave off both the initialization and the increment in a for statement, do not include the semicolons.

	j := 1
	for j < 100 {
		fmt.Println(j)
		j = j * 2
	}

	fmt.Println("")

}
