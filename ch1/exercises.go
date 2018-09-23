package main

import (
	"bufio"
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

// Exercise 1.4
func dupWithFileName() {
	args := os.Args[1:]

	if len(args) == 0 {
		dup1()
	} else {
		counts := make(map[string]int)
		files := make(map[string][]string)
		for _, arg := range args {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLinesAndFile(f, counts, files)
			f.Close()
		}
		printLinesAndFile(counts, files)
	}
}

func countLinesAndFile(f *os.File, counts map[string]int, files map[string][]string) {
	input := bufio.NewScanner(f)
	name := f.Name()
	for input.Scan() {
		counts[input.Text()]++
		if !stringInSlice(name, files[input.Text()]) {
			files[input.Text()] = append(files[input.Text()], name)
		}
	}
}

func printLinesAndFile(counts map[string]int, files map[string][]string) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d - %s\t%s\n", n, files[line], line)
		}
	}
}

func stringInSlice(name string, files []string) bool {
	for _, file := range files {
		if file == name {
			return true
		}
	}
	return false
}
