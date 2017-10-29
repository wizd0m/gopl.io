package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	proc3()
	proc1()
	proc2()
	proc1()
	proc3()
	proc2()
}

func proc1() {
	start := time.Now()
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("Proc1: %.10f seconds elapsed\n", time.Since(start).Seconds())
}

func proc2() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("Proc2: %.10f seconds elapsed\n", time.Since(start).Seconds())
}

func proc3() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("Proc3: %.10f seconds elapsed\n", time.Since(start).Seconds())
}
