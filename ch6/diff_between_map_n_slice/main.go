package main

import (
	"fmt"
	"os"
)

func main() {
	justAnother("makefile")
}

func justAnother(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	fmt.Println(file)
	return nil
	// defer file.Close()
	// data := make([]byte, 100)
	// for {
	// 	count, err := file.Read(data)
	// 	process(data[:count])
	// 	if err != nil {
	// 		if errors.Is(err, io.EOF) {
	// 			return nil
	// 		}
	// 		return err
	// 	}
	// }
}
