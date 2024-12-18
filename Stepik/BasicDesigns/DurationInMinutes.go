package main

import (
	"fmt"
	"time"
)

func DurationInMinutes() {
	var importTime string
	fmt.Println("Inter time to convert to minutes")
	fmt.Scan(&importTime)
	randTime, err := time.ParseDuration(importTime)

	if err != nil {
		fmt.Printf("Encorrect data entry: %v\n", err)
	} else {
		fmt.Printf("Time in minutes: %vm\n", randTime.Minutes())
	}
}