```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// simpleGenerator создает канал, в который будет отправлено n случайных чисел.
func simpleGenerator(n int) <-chan int {
    // тут пишем код
}

func main() {
	generate := simpleGenerator(5) // Создаем генератор, который сгенерирует 5 чисел

    // Здесь необходимо вывести все числа полученные через генератор
}
```
