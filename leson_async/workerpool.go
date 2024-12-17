package main

import (
	"fmt"
	"sync"
)

// 4. Создай worker pool(не знаешь что такое worker pool - в интернет), где несколько рабочих (горутин) выполняют задачи из канала.
func WorkerPool() {
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