package main

import (
	"fmt"
	"strconv"
)

func RepeatLine() {
	var line string
	var newLine string
	var scaned string
	fmt.Println("Enter the line to repeat")
	fmt.Scan(&line)
	fmt.Println("Enter the number of repetitions")
	fmt.Scan(&scaned)
	current, err := strconv.Atoi(scaned)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		for i := 0; i < current; i++ {
			newLine += line
		}
		fmt.Printf("New line: %v\n", newLine)
	}
}