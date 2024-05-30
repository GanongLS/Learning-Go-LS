package main

import (
	"fmt"
)

func main() {
	callDivN()

}

func callDivN() {
	fmt.Println("divN")
	var n = divN(2, 1)
	fmt.Println("that is n:", n)
	fmt.Println("")
}
func divN(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}
