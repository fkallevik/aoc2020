package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
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

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

type day func()

func open(fn day) {
	start := time.Now()
	fmt.Printf("************************************************\n")
	fmt.Printf("%s Answers:\n", getFunctionName(fn))
	fmt.Printf("************************************************\n")
	fn()
	elapsed := time.Since(start)
	fmt.Printf("\nComputation took %s\n", elapsed)
	fmt.Printf("************************************************\n")
}

func main() {
	//open(Day1)
	//open(Day2)
	//open(Day3)
	open(Day4)
}
