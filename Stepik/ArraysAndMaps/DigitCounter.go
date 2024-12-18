package main

import (
	"fmt"
	"sort"
)

func DigitCounter() {
	var number int
	fmt.Println("Enter a number to check the number of digits")
	fmt.Scanf("%d", number)

	numberMap := map[int]int{}
	var nNumber []int
	for number > 0 {
		nNumber = append(nNumber, number%10)
		number /= 10
	}

	for _, v := range nNumber {
		if _, ok := numberMap[v]; !ok {
			numberMap[v] = 1
		} else {
			numberMap[v]++
		}
	}

	printCounter(numberMap)
}

func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}