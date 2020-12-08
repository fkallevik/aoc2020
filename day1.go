package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var validCombos = make(map[string]int)

func rec(c []int, l int, expenses *[]int, number int) {
	for _, currentExp := range *expenses {

		comboArr := append(c, currentExp)

		// Sort array so we dont get duplicates in our validCombos map.
		sort.Ints(comboArr)

		sum := 0
		product := 1
		for _, p := range c {
			sum += p
			product *= p
		}

		if sum+currentExp == number {
			validCombos[ArrayToString(comboArr, "*")] = product * currentExp
		}

		if len(comboArr) < l {
			rec(comboArr, l, expenses, number)
		}
	}

}

func findProductOfNExpensesThatEqualNumber(n int, number int) {
	data := ReadFile("input/day1")
	expensesStrings := strings.Split(data, "\n")
	var expensesInts []int

	for _, s := range expensesStrings {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		expensesInts = append(expensesInts, num)
	}

	for _, v := range expensesInts {
		rec([]int{v}, n, &expensesInts, number)
	}

	fmt.Printf("Valid combinations %v\n", validCombos)
}

func Day1() {
	findProductOfNExpensesThatEqualNumber(3, 2020)
}
