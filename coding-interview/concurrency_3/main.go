package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://www.lamoda.ru",
		"https://www.yandex.ru",
		"https://www.mail.ru",
		"https://www.google.ru",
	}
	start := time.Now()
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fmt.Printf("Fetching %s...\n", url)

			err := fetchURL(url)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", url, err)
				return
			}

			fmt.Printf("Fetched %s\n", url)
		}(url)
	}
	wg.Wait()

	fmt.Println("All requests launched!")
	fmt.Println("Program finished.")
	fmt.Printf("Сколько времени это заняло: %s\n", time.Since(start).String())
}

func fetchURL(url string) error {
	_, err := http.Get(url)
	return err
}
