package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	/*
		another form of the for loop iterates over a range of values
		from a data type like a string or a slice.

		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
	*/

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
