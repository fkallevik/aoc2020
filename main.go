package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func ReadFile(name string) string {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(b))
}

func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

type day func()

func open(fn day) {
	start := time.Now()
	fmt.Printf("************************************************\n")
	fmt.Printf("Day 1 Answers:\n")
	fmt.Printf("************************************************\n")
	fn()
	elapsed := time.Since(start)
	fmt.Printf("Computation took %s\n", elapsed)
	fmt.Printf("************************************************\n")
}

func main() {
	open(Day1)
}
