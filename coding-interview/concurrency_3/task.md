```go
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
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s...\n", url)

			err := fetchURL(url)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", url, err)
				return
			}

			fmt.Printf("Fetched %s\n", url)
		}(url)
	}

	fmt.Println("All requests launched!")
	time.Sleep(10*time.Millisecond)
	fmt.Println("Program finished.")
	fmt.Printf("Сколько времени это заняло: %s\n", time.Since(start).String())
}

func fetchURL(url string) error {
	_, err := http.Get(url)
	return err
}
```
