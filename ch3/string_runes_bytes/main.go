package main

import (
	"fmt"
)

func main() {
	extractingString()
	sunSlicing()
	stringConversion()
	stringToSlice()
}

func stringToSlice() {
	fmt.Println("String to Slice")

	var s string = "Hello 🌞"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
	fmt.Println("")

}

func stringConversion() {
	fmt.Println("String Conversion")

	// Because of this complicated relationship among runes, strings, and bytes, Go has some interesting type conversions between these types. A single rune or byte can be converted to a string:
	var a rune = 'x'
	var s string = string(a)
	fmt.Println(s)

	var b byte = 'y'
	var s2 string = string(b)
	fmt.Println(s2)

	// A common bug for new Go developers is to try to make an int into a string by using a type conversion:
	// var x int = 65
	var x rune = 65   //conversion from int to string yields a string of one rune, not a string of digits (did you mean fmt.Sprint(x)?)
	var y = string(x) //conversion from int to string yields a string of one rune, not a string of digits (did you mean fmt.Sprint(x)?)
	var z = fmt.Sprint(x)
	fmt.Println(y)
	fmt.Println(z)
	// This results in y having the value “A,” not “65.” As of Go 1.15, go vet blocks a type conversion to string from any integer type other than rune or byte.

	fmt.Println("")

}

func sunSlicing() {
	fmt.Println("Sun Slicing")

	var s string = "Hello 🌞"
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(len(s))
	fmt.Println("")

}

func extractingString() {
	fmt.Println("Extracting String")

	// Just as you can extract a single value from an array or a slice, you can extract a single value from a string by using an index expression:
	var r string = "Hello there"
	var b byte = r[6]
	var bst string = r[6:8]
	fmt.Printf("%q \n", b)
	fmt.Printf("%s \n", bst)

	s := "Wow look at me"
	bs := []byte(s)
	fmt.Printf("%s \n", bs) // Output: Wow look at me
	fmt.Printf("%d \n", bs) // Output: [87 111 119 32 108 111 111 107 32 97 116 32 109 101]

	// The slice expression notation that you used with arrays and slices also works with strings:
	var s2 string = r[4:7]
	var s3 string = r[:5]
	var s4 string = r[6:]
	fmt.Printf("%s \n", s2)
	fmt.Printf("%s \n", s3)
	fmt.Printf("%s \n", s4)

	fmt.Println("")

}
