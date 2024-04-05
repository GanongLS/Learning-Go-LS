package main

import (
	"fmt"
	"maps"
)

func main() {
	mapsDeclaration()
	readingAndWrittingAMap()
	commaOKIdiom()
	deletingFromMap()
	emptyingMap()
	comparingMaps()
	usingMapsAsSets()
}

func usingMapsAsSets() {
	fmt.Println("Using Maps as Sets")

	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	// Some people prefer to use struct{} for the value when a map is being used to implement a set. (I’ll discuss structs in the next section.) The advantage is that an empty struct uses zero bytes, while a boolean uses one byte.
	// The disadvantage is that using a struct{} makes your code clums‐ ier. You have a less obvious assignment, and you need to use the comma ok idiom to check if a value is in the set:
	intSetStruct := map[int]struct{}{}
	valsStruct := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range valsStruct {
		intSetStruct[v] = struct{}{}
	}
	if _, ok := intSetStruct[100]; ok {
		fmt.Println("5 is in the set")
	}

	// Unless you have very large sets, the difference in memory usage will not likely be significant enough to outweigh the disadvantages.

	fmt.Println("")

}

func comparingMaps() {
	fmt.Println("Comparing maps")

	m := map[string]int{"hello": 5,
		"world": 10,
	}
	n := map[string]int{
		"world": 10,
		"hello": 5}
	fmt.Println("is m and n equal?:", maps.Equal(m, n)) // prints true

	fmt.Println("")

}

func emptyingMap() {
	fmt.Println("Emptying Map")
	m := map[string]int{"hello": 5,
		"world": 10,
	}
	fmt.Println("m:", m, len(m))
	clear(m)
	fmt.Println("m:", m, len(m))

	fmt.Println("")

}

func deletingFromMap() {
	fmt.Println("Deleting from Map")
	// Key-value pairs are removed from a map via the built-in delete function:
	m := map[string]int{"hello": 5,
		"world": 10,
	}
	delete(m, "hello")
	fmt.Println("m:", m)

	fmt.Println("")

	// The delete function takes a map and a key and then removes the key-value pair with the specified key. If the key isn’t present in the map or if the map is nil, nothing happens. The delete function doesn’t return a value.
}

func commaOKIdiom() {

	//As you’ve seen, a map returns the zero value if you ask for the value associated with a key that’s not in the map. This is handy when implementing things like the totalWins counter you saw earlier. However, you sometimes do need to find out if a key is in a map. Go provides the comma ok idiom to tell the difference between a key that’s associated with a zero value and a key that’s not in the map:

	fmt.Println("comma ok idiom")

	m := map[string]int{"hello": 5,
		"world": 0}
	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["world"]
	fmt.Println(v, ok)
	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	fmt.Println("")

}

func readingAndWrittingAMap() {
	fmt.Println("Reading and Writting Map")

	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println("orcas: ", totalWins["Orcas"])
	fmt.Println("kittens: ", totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println("kittens: ", totalWins["Kittens"])
	fmt.Println("lions: ", totalWins["Lions"])
	totalWins["Lions"] = 3
	fmt.Println("lions: ", totalWins["Lions"])

	fmt.Println("")

	// When you try to read the value assigned to a map key that was never set, the map returns the zero value for the map’s value type. In this case, the value type is an int, so you get back a 0. You can use the ++ operator to increment the numeric value for a map key. Because a map returns its zero value by default, this works even when there’s no existing value associated with the key.

}

func mapsDeclaration() {
	fmt.Println("Maps Declaration")

	// First, you can use a var declaration to create a map variable that’s set to its zero value:
	var nilMap map[string]int
	// nilMap["Ogre"] = 1
	fmt.Println(nilMap)

	// However, attempting to write to a nil map variable causes a panic.
	// You can use a := declaration to create a map variable by assigning it a map literal:
	totalWins := map[string]int{}
	totalWins["Ogre"] = 1

	fmt.Println(totalWins)

	// Here’s what a nonempty map literal looks like:
	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}

	orcas := teams["Orcas"]

	fmt.Println(teams)
	fmt.Println("orcas: ", orcas)

	// If you know how many key-value pairs you intend to put in the map but don’t know the exact values, you can use make to create a map with a default size:
	ages := make(map[int][]string, 10)
	ages[33] = []string{"Luki"}

	fmt.Println("ages: ", ages)

	fmt.Println("")

}
