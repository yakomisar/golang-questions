package main

import (
	"fmt"
	"math/rand"
	"time"
)

// simpleGenerator создает канал, в который будет отправлено n случайных чисел.
func simpleGenerator(n int) <-chan int {
	// тут пишем код
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- rand.Intn(100) + 1 // 1 -- 100 }
			// Небольшая задержка, чтобы почуствовать магический процесс
			time.Sleep(time.Millisecond * 100)
		}
	}()

	return out
}

func main() {
	// Вызываем шляпу
	generate := simpleGenerator(5) // Создаем генератор, который сгенерирует 5 чисел

	// Здесь необходимо вывести все числа полученные через генератор
	for number := range generate {
		fmt.Println("Счастливое число:", number)
	}
}
