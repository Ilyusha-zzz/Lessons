package main

import (
	"fmt"
	"sync"
)

// 2. Напиши программу, которая симулирует загрузку данных с нескольких "сайтов". Используй горутины для одновременного выполнения задач.
func DownloadFromURL() {
	var urlToDownload []string = []string{"url_1...","url_2...","url_3...","url_n..."}
	var wg sync.WaitGroup
	err := true
	
	for _, url := range urlToDownload {
		wg.Add(1)
		go func(url string) {
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