package main

import (
	"fmt"
	"math/rand"
	"time"
)

// simpleGenerator создает канал, в который будет отправлено n случайных чисел.
// Представьте, что это «лотерейные билеты».
func simpleGenerator(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			// Выдаем случайное число от 1 до 100
			out <- rand.Intn(100) + 1
			// Небольшая пауза для реалистичности
			time.Sleep(time.Millisecond * 100)
		}
	}()
	return out
}

// multiplier получает данные (наши «лотерейные билеты») и выполняет операцию умножения на x.
// Это как будто наш «бонус-множитель» или «увеличитель выигрыша».
func multiplier(val <-chan int, x int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range val {
			out <- num * x
		}
	}()
	return out
}

func main() {
	// Генерируем 10 случайных «билетов» и умножаем их значение на 2
	pipeline := multiplier(simpleGenerator(10), 2)

	// Выводим на экран итоговые «выигрышные значения»
	for ticket := range pipeline {
		fmt.Println("Выигрышный билет:", ticket)
	}
}
