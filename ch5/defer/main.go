package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	simpleCat()
	deferExample()
}

func deferExample() int { // last in defer first out, last apear the defer the first to be executed. defer bakal cuman mau ngerun function dan nutupnya kalau semua dependencies di engin bersih.
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("exiting:", a)
	return a
}

// Normally, a function call runs immediately, but defer delays the invocation until the surrounding function exits.
func simpleCat() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
