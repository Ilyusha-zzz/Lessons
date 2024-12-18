package main

import (
	"fmt"
	"math"
)

func DistanceBetweenPoints() {
	var x1, x2, y1, y2 float64
	fmt.Println("Enter the coordinates of the first point (x1 y1)")
	fmt.Scan(&x1, &y1)
	fmt.Println("Enter the coordinates of the second point (x2 y2)")
	fmt.Scan(&x2, &y2)

	distance := math.Sqrt(math.Pow(x1-x2, 2)+math.Pow(y1-y2, 2))

	fmt.Printf("Distance between points: %v\n", distance)
}