package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	/*
	n1()
	n2()
	n3()
	n4()
	*/
	n_Goroutine()
}

	
// 1. Напиши функцию, которая использует `time.Sleep` для создания таймера. После ожидания функция должна выводить сообщение в консоль.
func n1() {
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

// 2. Напиши программу, которая симулирует загрузку данных с нескольких "сайтов". Используй горутины для одновременного выполнения задач.
func n2() {
	var urlToDownload []string = []string{"url_1...","url_2...","url_3...","url_n..."}
	var wg sync.WaitGroup
	err := true
	
	for _, url := range urlToDownload {
		go func(url string) {
			wg.Add(1)
			defer wg.Done()
			fmt.Println("Start Download from ",url )
			if !err{
				fmt.Println("Download is not completed. Error:", err)
			}
			fmt.Println("Download complited from ", url)
			}(url)
	}
	wg.Wait()
}

// 3. Создай простую программу, где несколько горутин отправляют сообщения в канал, а главная функция их читает.
func n3() {
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

// 4. Создай worker pool(не знаешь что такое worker pool - в интернет), где несколько рабочих (горутин) выполняют задачи из канала.
func n4() {
	const numJob = 13
	const numWorker = 3
	chJob := make(chan int, numJob)
	chResult := make(chan int, numJob)
	var wg sync.WaitGroup

	for i := 1; i <= numWorker; i++ {
		go func (i int, chJob <- chan int, chResult chan <- int)  {
			for v := range chJob {
				current := v 
				fmt.Println("Worker ", i, " complete job ", current)
				chResult <- current * 2
				
			}
		}(i, chJob, chResult)
	}
	
	for i := 1; i <= numJob; i++ {
		chJob <- i
	}
	close(chJob)

	for j := 1; j <= numJob; j++ {
		<-chResult
	}
	close(chResult)

	wg.Wait()
}