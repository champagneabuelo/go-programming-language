package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

// Exercise 1.7
func fetchStdout() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%v", b)
	}
}

// Exercise 1.8
func fetchCheckPrefix() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			fmt.Printf("Prefix is missing, add it now.\n")
			fixedURL := "http://" + url
			fetchURL(fixedURL)
		} else {
			fetchURL(url)
		}
	}
}

// Execise 1.9
func fetchWithStatusCode() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		status := resp.Status
		fmt.Printf("fetch: Status code: %s\n", status)
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

// Helper function for fetch calls.
func fetchURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}
