package main

import (
	"fmt"
	"time"
	"sync"
	"strconv"
)

// 1. Напиши функцию, которая использует `time.Sleep` для создания таймера. После ожидания функция должна выводить сообщение в консоль.
func TimeSleep() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func () {
		for i := 1; i < 10; i++ {		
			t, _ := time.ParseDuration(strconv.Itoa(i)+"s")
			fmt.Printf("second: %v\n", t.Seconds())
			time.Sleep(1000*time.Millisecond)
		}
	} ()

	go func ()  {
		fmt.Println("Start")
		time.Sleep(3000*time.Millisecond)
		fmt.Println("End")
		wg.Done()
	} ()
	wg.Wait()	
}