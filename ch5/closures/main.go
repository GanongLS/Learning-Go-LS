package main

import (
	"fmt"
	"sort"
)

func main() {
	closures()
	shadowingInClosures()
	createPerson()
	callingMakeMult()
}

// Returning Functions from Functions

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

// And here is how the function is used:
func callingMakeMult() {
	fmt.Println("make mult")

	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 5; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
	fmt.Println("")

}

// Passing Functions as Parameters

func createPerson() {
	fmt.Println("Create Person")

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	// sort by last name
	fmt.Println("sort by lastname")

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName // people is captured by the closure.
	})
	fmt.Println(people)

	// sort by age
	fmt.Println("sort by age")

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
	fmt.Println("")

}

// Just as with any inner scope, you can shadow a variable inside a closure. If you change the code to:

func shadowingInClosures() {
	fmt.Println("Shadowing In Closures")

	a := 20
	f := func() {
		fmt.Println(a)
		a := 30
		fmt.Println(a)
	}
	f()
	fmt.Println(a)
	fmt.Println("")

}

// Closures
// Functions declared inside functions are special; they are closures. This is a computer science word that means that functions declared inside functions are able to access and modify variables declared in the outer function.

func closures() {
	fmt.Println("Closures")

	a := 20
	f := func() {
		fmt.Println(a)
		a = 30
	}
	f()
	fmt.Println(a)
	fmt.Println("")

}
