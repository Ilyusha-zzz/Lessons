package main

import "fmt"

// 3. Создай простую программу, где несколько горутин отправляют сообщения в канал, а главная функция их читает.
func SeveralGoroutin() {
	ch := make(chan int)
	var wg = 0
	
	go func(wg *int) {
		*wg++
		for i := 0; i < 5; i++ {
			ch <- i
		}
		*wg--
		if *wg == 0 {
			close(ch)
		} 
	}(&wg)

	go func(wg *int) {
		*wg++
		for i := 5; i < 10; i++ {
			ch <- i
		}
		*wg--
		if *wg == 0 {
			close(ch)
		} 
	}(&wg)

	go func(wg *int) {
		*wg++
		for i := 10; i < 15; i++ {
			ch <- i
		}
		*wg--
		if *wg == 0 {
			close(ch)
		} 
	}(&wg)
		
	for c := range ch {
		fmt.Printf("c: %v\n", c)
	}
}
