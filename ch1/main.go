package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println(testing.Benchmark(echo1Benchmark))
	fmt.Println(testing.Benchmark(echo2Benchmark))
	fmt.Println(testing.Benchmark(echo3Benchmark))
}
