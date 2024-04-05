package main

import (
	"fmt"
	"slices"
)

func main() {
	sliceInitiation()
	sliceComp()
	sliceAppend()
	variadicAppend()
	understandingCap()
	usingMake()
	emptyingSlice()
	declaringSlice()
	slicingSlice()
	sliceOverlapping()
	trickyAppend()
	independetSlice()
	overlappingAndCopy()
	arraySliceCopy()
	arrayToSlice()
	sliceToArray()
	arrayPointer()
}

func arrayPointer() {
	fmt.Println("Array Pointers")

	// I haven’t talked about pointers yet, but you can also use a type conversion to convert a slice into a pointer to an array:
	xSlice := []int{1, 2, 3, 4}
	xArrayPointer := (*[4]int)(xSlice)
	// After converting a slice to an array pointer, the storage between the two is shared. A change to one will change the other:
	xSlice[0] = 10
	xArrayPointer[1] = 20
	fmt.Println(xSlice) // prints [10 20 3 4] fmt.Println(xArrayPointer) // prints &[10 20 3 4]
	// Pointers are covered in Chapter 6.
	fmt.Println("")

}

func sliceToArray() {

	fmt.Println("Slice to Array")

	// Use a type conversion to make an array variable from a slice. You can convert an entire slice to an array of the same type, or you can create an array from a subset of the slice.
	// When you convert a slice to an array, the data in the slice is copied to new memory. That means that changes to the slice won’t affect the array, and vice versa.

	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)

	// panicArray := [5]int(xSlice) // will panic
	// fmt.Println(panicArray)

	fmt.Println("")

}

func arrayToSlice() {
	fmt.Println("Array to slice")

	// To convert an entire array into a slice, use the [:] syntax:
	xArray := [4]int{5, 6, 7, 8}
	xSlice := xArray[:]

	fmt.Println(xArray)
	fmt.Println(xSlice)

	// You can also convert a subset of an array into a slice:
	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// Be aware that taking a slice from an array has the same memory-sharing properties as taking a slice from a slice.

	x[0] = 10
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	fmt.Println("")

}

func arraySliceCopy() {
	fmt.Println("Array Slice and Copy")
	// You can use copy with arrays by taking a slice of the array. You can make the array either the source or the destination of the copy. You can try out the following code on The Go Playground or in the sample_code/copy_array directory in the Chapter 3 repository:

	x := []int{1, 2, 3, 4}
	d := [4]int{5, 6, 7, 8}
	y := make([]int, 2)
	copy(y, d[:])
	fmt.Println(y)
	copy(d[:], x)
	fmt.Println(d)
	fmt.Println("")

}

func overlappingAndCopy() {
	fmt.Println("overlapping Slice and Copy")

	// The copy function allows you to copy between two slices that cover overlapping sections of an underlying slice:
	x := []int{1, 2, 3, 4}
	num := copy(x[:3], x[1:])
	fmt.Println(x, num)
	fmt.Println("")

}

func independetSlice() {
	fmt.Println("Independent Slice")
	x := []int{16, 1, 2, 4, 8}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)

	// You can also copy a subset of a slice. The following code copies the first two elements of a four-element slice into a two-element slice
	a := []int{1, 2, 3, 4}
	b := make([]int, 2)
	numa := copy(b, a)
	fmt.Println(b, numa)

	copy(b, a[2:]) // copy ngereturn value tapi kalau ga ditangkap ga error. beda sama append

	fmt.Println("")
}

func trickyAppend() {
	fmt.Println("Tricky Append")

	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	fmt.Println(cap(x), cap(y))
	y = append(y, "z")
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Println("")

}

func sliceOverlapping() {
	fmt.Println("Slice Overlapping")

	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	fmt.Println("")

}

func slicingSlice() {
	fmt.Println("Slicing Slice")

	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	fmt.Println("")

}

func declaringSlice() {
	fmt.Println("declaring slice")
	// declaring nill slice
	var data []int
	fmt.Println("nil slice:", data)

	// You can create a slice using an empty slice literal:
	var x = []int{}
	fmt.Println("empty slice:", x)

	// This creates a slice with zero length and zero capacity. It is confus‐ ingly different from a nil slice. Because of implementation reasons, comparing a zero-length slice to nil returns false, while compar‐ ing a nil slice to nil returns true. For simplicity, favor nil slices. A zero-length slice is useful only when converting a slice to JSON. You’ll look at this more in “encoding/json” on page 327.

	fmt.Println("")

}

func emptyingSlice() {
	fmt.Println("emptying slice")

	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s))
	clear(s)
	fmt.Println(s, len(s))

	fmt.Println("")

}

func usingMake() {
	fmt.Println("Using make")
	x := make([]int, 5) //This creates an int slice with a length of 5
	fmt.Println(x, len(x), cap(x))

	// One common beginner mistake is to try to populate those initial elements using append:
	x = append(x, 10)

	fmt.Println(x, len(x), cap(x))

	y := make([]int, 5, 10) // This creates an int slice with a length of 5 and a capacity of 10.
	fmt.Println(y, len(y), cap(y))

	// You can also create a slice with zero length but a capacity that’s greater than zero:
	z := make([]int, 0, 10)
	fmt.Println(z, len(z), cap(z))
	// In this case, you have a non-nil slice with a length of 0 but a capacity of 10. Since the length is 0, you can’t directly index into it, but you can append values to it:

	z = append(z, 5, 6, 7, 8)

	fmt.Println(z, len(z), cap(z))

	fmt.Println("")
}

func understandingCap() {
	fmt.Println("Understanding Capacity")

	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	fmt.Println("")
}

func variadicAppend() {
	fmt.Println("variadic append")
	var u = [6]int{1, 2, 3, 4, 5, 6}
	var x = u[:3]
	y := []int{20, 30, 40}
	var z = append(x, y...) // append dengan value variadic seperti ini juga mengubah array dari slice yang diberikan.
	fmt.Println(u)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	// append(x, y...) //comp time error, value return dari append tidak digunakan.
	fmt.Println("")
}

func sliceAppend() {
	// The built-in append function is used to grow slices:
	var x []int
	x = append(x, 10) // assign result to the variable that's passed in
	fmt.Println(x)

	var a = []int{1, 2, 3}
	var b = append(a, 4)
	fmt.Println(a)
	fmt.Println(b)

	var c = [5]int{1, 2, 3, 4, 5}
	var d = c[1:3]
	var e = append(d, 6, 7, 8) // ketika append kepanjangan, go akan summon array baru sebagai dasar slice baru, array lama ga akan diubah, kalau append ga kepanjangan (ga out of cap, array lama akan dimodifikasi)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}

func sliceComp() {
	// you can't compare slice to other slice with comparison like == or != or other, you only could compare slice with nil
	var aa []int
	fmt.Println(aa == nil) // prints true
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	s := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(x, y)) // prints true
	fmt.Println(slices.Equal(x, z)) // prints false
	// fmt.Println(slices.Equal(x, s)) // does not compile, slice int ga bisa dibandingin dengan slice string.
	fmt.Println(s == nil) // prints false
}

func sliceInitiation() {
	var x = []int{10, 20, 30}

	var y = []int{1, 5: 4, 6, 10: 100, 15}

	fmt.Println(x)
	fmt.Println(y)

	// You can simulate multidimensional slices and make a slice of slices:
	var z [][]int
	fmt.Println(z)

	// You read and write slices using bracket syntax, and, just as with arrays, you can’t read
	// or write past the end or use a negative index:

	x[0] = 10
	fmt.Println(x[2])

	// zero slice value is nill,

	var aa []int

	fmt.Println(aa)

}
