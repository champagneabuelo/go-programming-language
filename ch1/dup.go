package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// dup function that only reads from stdin
func dup1() {
	counts := make(map[string]int)
	countLines(os.Stdin, counts)
	printLines(counts)
}

// dup function that reads from stdin or list of named files
func dup2() {
	files := os.Args[1:]
	if len(files) == 0 {
		dup1()
	} else {
		counts := make(map[string]int)
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
		printLines(counts)
	}
}

// dup function that only reads named files
func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	printLines(counts)
}

// helper function to print duplicated lines
func printLines(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// helper function to count duplicated lines and add them to map
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
