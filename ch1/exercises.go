package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// Exercise 1.1
func echoWithCommand() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

// Exercise 1.2
func echoWithIndex() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s %d\n", os.Args[i], i)
	}
}

// Exercise 1.3
func echo1Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
	}
}

func echo2Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
	}
}

func echo3Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(os.Args[1:], " ")
	}
}
