package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// response хранит URL и результат его обработки
type response struct {
	url    string
	status bool
}

func main() {
	// Список URL для проверки
	urls := []string{
		"https://www.lamoda.ru",
		"https://www.yandex.ru",
		"https://www.mail.ru",
		"https://www.google.ru",
	}
	// Фиксируем время начала выполнения программы
	start := time.Now()

	// Используем WaitGroup для ожидания завершения всех горутин
	wg := sync.WaitGroup{}

	// Буферизованный канал для хранения ответов
	ch := make(chan response, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fmt.Printf("Fetching %s...\n", url)

			res, err := fetchURL(url)
			if err != nil {
				// Отправляем результат в канал
				ch <- res
				// fmt.Printf("Error fetching %s: %v\n", url, err)
				return
			}

			// Отправляем результат в канал
			ch <- res

			// fmt.Printf("Fetched %s\n", url)
		}(url)
	}

	// Горутина для закрытия канала после завершения всех запросов
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Чтение результатов из канала
	for val := range ch {
		// Дополнительная логика по обработке результатов
		fmt.Printf("url: %s, status: %t\n", val.url, val.status)
	}

	fmt.Println("All requests launched!")
	fmt.Println("Program finished.")
	fmt.Printf("Сколько времени это заняло: %s\n", time.Since(start).String())
}

// fetchURL выполняет HTTP-запрос к указанному URL и возвращает результат
func fetchURL(url string) (response, error) {
	_, err := http.Get(url)
	if err != nil {
		return response{
			url:    url,
			status: false,
		}, err
	}
	return response{
		url:    url,
		status: true,
	}, nil
}
