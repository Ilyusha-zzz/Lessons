package main

import (
	"fmt"
)

func DigitCounter() {
	var number int
	fmt.Println("Enter a number to check the number of digits")
	fmt.Scanf("%d", &number)

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

	fmt.Printf("numberMap: %v\n", numberMap)
}