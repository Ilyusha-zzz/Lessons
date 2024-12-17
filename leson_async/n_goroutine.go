package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func n_Goroutine() {
	r := rand.New(rand.NewSource((time.Now().Unix())))
	var wg sync.WaitGroup
	maxVar := r.Intn(100)

	wg.Add(maxVar)
	for i := 0; i < maxVar; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("i: %v\n", i)
		}(i)
	}

	wg.Wait()
	fmt.Printf("maxVar: %v\n", maxVar)

	// commit
}
