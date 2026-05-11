<p align="center">
  <img src="./assets/gopher_interview_image.png" alt="Gopher готовится к собеседованию по Go" width="720">
</p>

<h1 align="center">Golang Interview Questions 🚀</h1>

<p align="center">
  Вопросы по Go для подготовки к техническим собеседованиям: от базового синтаксиса до слайсов, мап, указателей, горутин, каналов, строк и практических задач.
</p>

<p align="center">
  <img alt="GitHub stars" src="https://img.shields.io/github/stars/yakomisar/golang-questions?style=social">
  <img alt="Go" src="https://img.shields.io/badge/Go-Interview%20Prep-00ADD8?logo=go&logoColor=white">
  <img alt="Language" src="https://img.shields.io/badge/language-Russian-blue">
  <img alt="License" src="https://img.shields.io/badge/license-MIT-green">
</p>

---

## О репозитории

Добро пожаловать в репозиторий с вопросами по Golang для подготовки к интервью.

Здесь собраны вопросы, которые могут встретиться на собеседовании, с разъяснениями, примерами кода и разбором corner cases. Представьте, что это ваши билеты на экзамен: вытягиваете билет, разбираете тему и постепенно закрываете пробелы.

## Как готовиться 💪

Перед началом подготовки лучше создать спокойную рабочую обстановку без отвлекающих факторов.

Не пугайтесь, если не можете ответить на вопрос сразу. Цель подготовки не в том, чтобы заучить ответы, а в том, чтобы понять механику языка и научиться рассуждать вслух.

Перед прохождением вопросов желательно повторить базовые темы Go:

- slices;
- maps;
- structs;
- pointers;
- interfaces;
- goroutines;
- channels;
- context;
- garbage collector;
- profiling.

---

## Оглавление 📚

- [Общие вопросы](#общие-вопросы)
- [Slices](#slices)
- [Maps](#maps)
- [Указатели](#указатели)
- [Goroutines и Channels](#goroutines-и-channels)
- [Работа со строками](#работа-со-строками)
- [Вклад](#вклад-)
- [Лицензия](#лицензия)

---

# Общие вопросы

## Вопрос 1. Чем горутины отличаются от потоков операционной системы?

Можете ли вы объяснить разницу между горутинами и потоками операционной системы?

<details>
<summary><strong>Ответ</strong></summary>

Горутины (`goroutines`) в Go — это легковесные потоки выполнения, которыми управляет рантайм Go.

Они разделяют общее адресное пространство и планируются внутри Go-процесса. В отличие от потоков операционной системы, горутины дешевле по памяти и проще масштабируются на большое количество параллельных задач.

Можно представить, что горутины — это падаваны в мире Star Wars. Они молоды, гибки и могут выполнять множество заданий, управляемые Силой, то есть рантаймом Go. Падаваны быстро переключаются между задачами и могут работать вместе.

Потоки операционной системы — это джедаи. Они мощнее и имеют прямой доступ к системным ресурсам, но они тяжелее, дороже в обслуживании и требуют больше внимания со стороны операционной системы.

Итого:

- горутины управляются рантаймом Go;
- OS threads управляются операционной системой;
- горутины дешевле по памяти;
- множество горутин могут выполняться поверх меньшего количества OS threads;
- рантайм Go сам занимается планированием горутин.

</details>

---

## Вопрос 2. Что такое пустой интерфейс `interface{}` и когда он полезен?

Можете привести пример кода, где пустой интерфейс пригодился бы для обработки разных типов данных?

<details>
<summary><strong>Ответ</strong></summary>

Пустой интерфейс `interface{}` в Go может хранить значение любого типа, потому что он не требует реализации каких-либо методов.

Он полезен, когда нужно работать с разными типами данных. Например, при логировании, сериализации, универсальных контейнерах или обработке данных неизвестного типа.

Пример:

```go
package main

import "fmt"

func printValue(value interface{}) {
    fmt.Println(value)
}

func main() {
    printValue(42) // Можно передать целое число
    printValue("Hello") // Можно передать строку
    printValue(3.14) // Можно передать число с плавающей точкой
}
```

Начиная с Go 1.18, для многих задач, где раньше использовали `interface{}`, можно использовать `any`.

</details>

---

## Вопрос 3. Как бы вы реализовали свой примитив для управления памятью?

Можете предоставить код, который демонстрирует создание, выделение и освобождение памяти с использованием указателей?

<details>
<summary><strong>Ответ</strong></summary>

В Go управление памятью автоматизировано: память выделяется рантаймом, а освобождается garbage collector. В обычном прикладном коде вручную управлять памятью не требуется.

Но если представить простой примитив управления памятью, он может выглядеть так:

```go
type MemoryManager struct {
    data []byte
}

func NewMemoryManager(size int) *MemoryManager {
    return &MemoryManager{
        data: make([]byte, size),
    }
}

func (mm *MemoryManager) Allocate(size int) []byte {
    if len(mm.data) < size {
        return nil
    }

    allocated := mm.data[:size]
    mm.data = mm.data[size:]

    return allocated
}

func (mm *MemoryManager) Free(allocated []byte) {
    mm.data = append(mm.data, allocated...)
}
```

Важно понимать, что это учебный пример. В реальном коде такой менеджер памяти может быть небезопасным и неэффективным без дополнительного контроля, синхронизации и учета фрагментации.

</details>

---

## Вопрос 4. Как в общих чертах работает Garbage Collector в Go?

<details>
<summary><strong>Ответ</strong></summary>

Garbage Collector в Go автоматически освобождает память, которая больше не используется программой.

Если использовать аналогию со Star Wars, можно представить GC как R2-D2, который летает по галактике памяти и проверяет, какие объекты еще нужны, а какие уже можно удалить.

### Mark phase

На фазе разметки GC проходит по объектам, которые достижимы из корней: стеков горутин, глобальных переменных и других runtime-структур.

В упрощенном виде используется трехцветная маркировка:

- **белый** — объект еще не посещался;
- **серый** — объект найден, но его ссылки еще не обработаны;
- **черный** — объект и его ссылки обработаны.

### Sweep phase

После разметки GC очищает память, занятую объектами, которые остались белыми, то есть недостижимыми.

### Главное

GC в Go работает конкурентно с программой и старается минимизировать stop-the-world паузы. Это не значит, что пауз нет совсем, но рантайм Go старается делать их короткими.

</details>

---

## Вопрос 5. В какой момент рантайм решает запустить сборщик мусора?

<details>
<summary><strong>Ответ</strong></summary>

Сборщик мусора в Go стремится найти баланс между использованием CPU и потреблением памяти.

Основной триггер запуска GC — рост heap-памяти после предыдущего цикла сборки мусора.

### Основные триггеры

1. **Выделение памяти**

   Когда программа выделяет память и размер heap достигает целевого порога, рантайм запускает новый цикл GC.

2. **Ручной запуск**

   Программист может вызвать:

   ```go
   runtime.GC()
   ```

   Но в большинстве случаев лучше не делать этого вручную и довериться рантайму.

3. **GC pacer**

   Рантайм Go использует механизм GC pacer, который рассчитывает, когда лучше начать следующий цикл GC, исходя из текущего объема live heap и скорости выделения памяти.

### GOGC

Переменная окружения `GOGC` управляет тем, насколько должен вырасти heap после прошлого GC, прежде чем будет запущен следующий цикл.

По умолчанию:

```sh
GOGC=100
```

Это означает, что если после прошлого GC осталось 4 MB живых объектов, следующий GC примерно запустится, когда heap достигнет 8 MB.

Примеры:

- `GOGC=50` — GC будет запускаться чаще;
- `GOGC=100` — значение по умолчанию;
- `GOGC=200` — GC будет запускаться реже, heap может вырасти примерно до трехкратного размера live heap.

</details>

---

## Вопрос 6. Какие недостатки есть у подхода Mark-and-Sweep?

<details>
<summary><strong>Ответ</strong></summary>

У подхода Mark-and-Sweep есть несколько недостатков.

### 1. Паузы

Классический Mark-and-Sweep может приводить к паузам выполнения программы, потому что рантайму нужно пройти по объектам и определить, какие из них живые.

### 2. Накладные расходы

Разметка и очистка требуют CPU-ресурсов. В высоконагруженных приложениях это может влиять на latency и throughput.

### 3. Фрагментация памяти

После очистки память может стать фрагментированной: свободные участки будут разбросаны по heap.

### 4. Сложность реализации

Нужно корректно отслеживать корни, ссылки между объектами, write barriers и взаимодействие с выполняющейся программой.

</details>

---

## Вопрос 7. Как интегрировать `pprof` и анализировать производительность приложения?

Вопрос может звучать по-разному:

> Как вы интегрируете pprof в приложение для сбора CPU и memory profiles?
>
> Какие команды используете для анализа профилей?
>
> Как находите узкие места в производительности?
>
> Как оптимизируете память на основе данных pprof?

<details>
<summary><strong>Ответ</strong></summary>

`pprof` — это инструмент для профилирования Go-приложений. Он помогает понять, где программа тратит CPU, где выделяет память и какие функции являются bottleneck.

### Пример CPU и memory profiling

```go
package main

import (
    "fmt"
    "os"
    "runtime/pprof"
    "time"
)

func main() {
    cpuFile, err := os.Create("cpu.pprof")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer cpuFile.Close()

    if err := pprof.StartCPUProfile(cpuFile); err != nil {
        fmt.Println(err)
        return
    }
    defer pprof.StopCPUProfile()

    memFile, err := os.Create("mem.pprof")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer memFile.Close()

    for i := 0; i < 1_000_000; i++ {
        s := make([]byte, 1000)

        if i%1000 == 0 {
            time.Sleep(500 * time.Millisecond)
        }

        for j := 0; j < len(s); j++ {
            s[j] = byte(j % 256)
        }
    }

    if err := pprof.WriteHeapProfile(memFile); err != nil {
        fmt.Println(err)
        return
    }
}
```

### Анализ CPU profile

```sh
go tool pprof cpu.pprof
```

Внутри интерактивного режима:

```sh
top
list functionName
web
```

### Анализ memory profile

```sh
go tool pprof mem.pprof
```

Команда `top` покажет функции, где выделяется больше всего памяти.

### Для web-сервисов

Часто `pprof` подключают через `net/http/pprof`:

```go
import _ "net/http/pprof"
```

И запускают debug HTTP server:

```go
go func() {
    http.ListenAndServe("localhost:6060", nil)
}()
```

После этого можно открыть:

```text
http://localhost:6060/debug/pprof/
```

</details>

---

## Вопрос 8. Что такое `GOGC` и `GODEBUG`?

<details>
<summary><strong>Ответ</strong></summary>

`GOGC` и `GODEBUG` — это переменные окружения, которые позволяют менять поведение рантайма Go.

### GOGC

`GOGC` управляет частотой запуска сборщика мусора.

Пример:

```sh
export GOGC=200
go run main.go
```

На Windows:

```sh
set GOGC=200
go run main.go
```

Чем выше `GOGC`, тем реже запускается GC, но тем больше памяти может использовать приложение.

Чем ниже `GOGC`, тем чаще запускается GC, но тем больше CPU может уходить на сборку мусора.

### GODEBUG

`GODEBUG` включает разные debug-режимы рантайма.

Например:

```sh
export GODEBUG=gctrace=1
go run main.go
```

Это выведет информацию о каждом цикле GC.

### В Dockerfile

```dockerfile
ENV GOGC=200
ENV GODEBUG=gctrace=1
```

### В коде

```go
package main

import "os"

func main() {
    _ = os.Setenv("GOGC", "200")
    _ = os.Setenv("GODEBUG", "gctrace=1")
}
```

</details>

---

## Вопрос 9. Что такое middleware и router в Go?

<details>
<summary><strong>Ответ</strong></summary>

### Middleware

Middleware — это функция, которая выполняется до или после основного HTTP handler.

Middleware часто используют для:

- логирования;
- аутентификации;
- авторизации;
- CORS;
- recovery после panic;
- добавления request ID;
- метрик;
- rate limiting.

Упрощенно middleware можно представить как checkpoint на дороге. Перед тем как запрос попадет в основной handler, он проходит через несколько проверок.

### Router

Router отвечает за то, чтобы направить HTTP-запрос в правильный handler на основе URL, HTTP method и других параметров.

Например:

- `GET /users` → получить пользователей;
- `POST /users` → создать пользователя;
- `GET /users/{id}` → получить конкретного пользователя.

В Go для routing часто используют стандартный `net/http`, `chi`, `gorilla/mux`, `httprouter` и другие библиотеки.

</details>

---

# Slices

## Вопрос 1. Какие значения будут у `len(s)` и `cap(s)`?

```go
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]
```

<details>
<summary><strong>Ответ</strong></summary>

`a` — это массив из пяти элементов.

`s` — это слайс, который ссылается на массив `a` и содержит элементы:

```go
[]int{2, 3, 4}
```

Значения:

```go
len(s) == 3
cap(s) == 4
```

Почему `cap(s) == 4`?

Слайс начинается с индекса `1` исходного массива и может расширяться до конца массива:

```text
[1, 2, 3, 4, 5]
    ^-------- до конца массива 4 элемента
```

</details>

---

## Вопрос 2. Как изменится `base` после выполнения кода?

```go
base := []int{10, 20, 30, 40}
newSlice := base[1:3]
newSlice[1] = 50
```

<details>
<summary><strong>Ответ</strong></summary>

`newSlice` ссылается на тот же underlying array, что и `base`.

До изменения:

```go
base     // []int{10, 20, 30, 40}
newSlice // []int{20, 30}
```

После:

```go
newSlice[1] = 50
```

Изменится элемент `base[2]`.

Итог:

```go
base     // []int{10, 20, 50, 40}
newSlice // []int{20, 50}
```

</details>

---

## Вопрос 3. Какие будут `len(original)` и `cap(original)`?

```go
original := make([]int, 3, 5)
original = append(original, 1, 2, 3)
```

<details>
<summary><strong>Ответ</strong></summary>

После создания:

```go
original := make([]int, 3, 5)
```

Слайс выглядит так:

```go
[]int{0, 0, 0}
```

Имеет:

```go
len(original) == 3
cap(original) == 5
```

После:

```go
original = append(original, 1, 2, 3)
```

Нужно добавить три элемента. Текущая емкость `5`, а новая длина будет `6`, поэтому произойдет реаллокация.

Итог:

```go
original // []int{0, 0, 0, 1, 2, 3}
len(original) == 6
cap(original) >= 6
```

Часто в такой ситуации capacity может стать `10`, но точный алгоритм роста capacity является деталью реализации рантайма и может зависеть от версии Go и размера слайса.

</details>

---

## Вопрос 4. Чем отличаются nil slice и empty slice?

```go
var nilSlice []int
emptySlice := make([]int, 0)
```

Что вернут проверки?

```go
nilSlice == nil
emptySlice == nil
```

<details>
<summary><strong>Ответ</strong></summary>

`nilSlice` — это nil-слайс. Он не указывает на underlying array.

```go
var nilSlice []int
```

У него:

```go
len(nilSlice) == 0
cap(nilSlice) == 0
nilSlice == nil // true
```

`emptySlice` — это пустой, но инициализированный слайс.

```go
emptySlice := make([]int, 0)
```

У него:

```go
len(emptySlice) == 0
cap(emptySlice) == 0
emptySlice == nil // false
```

В большинстве случаев оба можно использовать одинаково: по ним можно итерироваться и в них можно делать `append`.

</details>

---

## Вопрос 5. Каким будет значение `slices` после выполнения кода?

```go
slices := [][]int{
    {1, 2},
    {3, 4},
}

slices[0] = append(slices[0], 3)
```

<details>
<summary><strong>Ответ</strong></summary>

Мы добавляем значение `3` в первый вложенный слайс.

Итог:

```go
slices := [][]int{
    {1, 2, 3},
    {3, 4},
}
```

</details>

---

## Вопрос 6. Почему `append` иногда требует новый участок памяти, а иногда нет?

Как это связано с capacity слайса?

<details>
<summary><strong>Ответ</strong></summary>

Слайс в Go — это структура, которая содержит:

- указатель на underlying array;
- длину `len`;
- емкость `cap`.

Если при `append` в текущем underlying array достаточно места, новый элемент будет добавлен в тот же массив.

Если места недостаточно, Go создаст новый underlying array, скопирует туда старые элементы и добавит новые.

Пример:

```go
s := make([]int, 0, 2)
s = append(s, 1)
s = append(s, 2)

// Здесь capacity уже заполнена.
s = append(s, 3) // скорее всего будет новый underlying array
```

Реаллокация может быть дорогой, потому что нужно выделить новую память и скопировать элементы.

</details>

---

## Вопрос 7. Nil slice vs empty slice: когда какой использовать?

<details>
<summary><strong>Ответ</strong></summary>

`nil` slice:

```go
var s []int
```

Пустой slice:

```go
s := []int{}
// или
s := make([]int, 0)
```

Оба имеют длину `0`, и в оба можно делать `append`.

Разница может быть важна при сериализации, например в JSON:

```go
var nilSlice []int = nil
emptySlice := []int{}
```

При JSON marshaling результат может отличаться:

```json
null
```

и

```json
[]
```

В API часто предпочитают возвращать пустой массив `[]`, а не `null`, если клиент ожидает список.

</details>

---

## Вопрос 8. Как удалить элемент из слайса без стандартной библиотеки, не нарушив порядок?

<details>
<summary><strong>Ответ</strong></summary>

Если порядок важен, можно использовать `append`:

```go
func removeOrdered(s []int, i int) []int {
    return append(s[:i], s[i+1:]...)
}
```

Если порядок не важен, можно заменить удаляемый элемент последним:

```go
func removeUnordered(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}
```

Второй вариант быстрее, потому что не нужно сдвигать элементы.

</details>

---

## Вопрос 9. Освободится ли память после `largeSlice = largeSlice[:5]`?

<details>
<summary><strong>Ответ</strong></summary>

Нет, не обязательно.

Если большой слайс был обрезан до маленького, новый слайс все еще может ссылаться на тот же underlying array.

Пример:

```go
largeSlice := make([]byte, 10<<20) // 10 MB
smallSlice := largeSlice[:5]
```

Даже если нам нужны только первые 5 байт, `smallSlice` может удерживать весь массив на 10 MB.

Чтобы дать GC возможность освободить большой массив, можно скопировать нужные данные в новый слайс:

```go
smallCopy := make([]byte, len(smallSlice))
copy(smallCopy, smallSlice)
largeSlice = nil
```

Или использовать:

```go
smallCopy := append([]byte(nil), smallSlice...)
```

</details>

---

## Вопрос 10. Что выведет этот код?

```go
taskList := []string{
    "Проснуться",
    "Покушать",
    "Поработать",
}

wakeup := taskList[0:2]
work := taskList[2:3]

wakeup = append(wakeup, "Погулять с собакой")

fmt.Println("Wakeup staff:", wakeup)
fmt.Println("Workstaff:", work)
```

<details>
<summary><strong>Ответ</strong></summary>

Исходный слайс:

```go
[]string{"Проснуться", "Покушать", "Поработать"}
```

После:

```go
wakeup := taskList[0:2]
```

Получаем:

```go
wakeup == []string{"Проснуться", "Покушать"}
len(wakeup) == 2
cap(wakeup) == 3
```

После:

```go
work := taskList[2:3]
```

Получаем:

```go
work == []string{"Поработать"}
len(work) == 1
cap(work) == 1
```

Так как у `wakeup` есть свободная capacity, `append` перезапишет следующий элемент underlying array.

```go
wakeup = append(wakeup, "Погулять с собакой")
```

Итоговый вывод:

```text
Wakeup staff: [Проснуться Покушать Погулять с собакой]
Workstaff: [Погулять с собакой]
```

</details>

---

# Maps

## Вопрос 1. Что произойдет при выполнении кода с nil map?

```go
var m map[string]int
m["key"] = 42
```

Почему это происходит и как исправить?

<details>
<summary><strong>Ответ</strong></summary>

`map` в Go — это ссылочный тип. Нулевое значение map — `nil`.

В nil map можно читать, но нельзя записывать.

Этот код вызовет panic:

```text
panic: assignment to entry in nil map
```

Исправление:

```go
m := make(map[string]int)
m["key"] = 42
```

Или:

```go
m := map[string]int{}
m["key"] = 42
```

</details>

---

## Вопрос 2. Что будет выведено и почему?

```go
func modifyMap(m map[int]string) {
    m[2] = "changed"
}

func main() {
    myMap := map[int]string{1: "one", 2: "two", 3: "three"}
    modifyMap(myMap)
    fmt.Println(myMap)
}
```

<details>
<summary><strong>Ответ</strong></summary>

Map — ссылочный тип. При передаче map в функцию копируется descriptor, но он указывает на те же данные.

Поэтому изменение внутри функции будет видно снаружи.

Вывод:

```go
map[1:one 2:changed 3:three]
```

Порядок вывода ключей не гарантирован.

</details>

---

## Вопрос 3. Безопасно ли удалять элементы из map во время итерации?

```go
m := map[int]bool{1: true, 2: true, 3: true}

for k := range m {
    if k == 2 {
        delete(m, k)
    }
}
```

<details>
<summary><strong>Ответ</strong></summary>

Да, удалять элементы из map во время итерации в этой же горутине безопасно.

Но важно помнить:

- порядок итерации по map не гарантирован;
- конкурентное чтение и запись из разных горутин без синхронизации небезопасны;
- concurrent map writes могут привести к panic.

</details>

---

## Вопрос 4. Как отличить отсутствующий ключ от нулевого значения?

<details>
<summary><strong>Ответ</strong></summary>

В Go при чтении из map можно получить два значения:

```go
value, ok := m[key]
```

- `value` — значение по ключу или zero value типа;
- `ok` — `true`, если ключ есть, и `false`, если ключа нет.

Пример:

```go
checkMap := map[string]int{
    "count": 0,
}

if val, ok := checkMap["count"]; ok {
    fmt.Println("Ключ существует, значение:", val)
} else {
    fmt.Println("Ключ отсутствует")
}
```

</details>

---

## Вопрос 5. Гарантирован ли порядок итерации по map?

<details>
<summary><strong>Ответ</strong></summary>

Нет.

Порядок итерации по map в Go не гарантирован.

Нельзя писать код, который зависит от порядка обхода map.

Если нужен стабильный порядок:

1. Соберите ключи в слайс.
2. Отсортируйте ключи.
3. Итерируйтесь по отсортированному слайсу ключей.

Пример:

```go
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}

sort.Strings(keys)

for _, k := range keys {
    fmt.Println(k, m[k])
}
```

</details>

---

# Указатели

## Вопрос 1. Что выведет следующий код?

```go
func setLinkHome(link *string) {
    *link = "http://home"
}

link := "http://other"
setLinkHome(&link)
fmt.Println(link)
```

<details>
<summary><strong>Ответ</strong></summary>

Код выведет:

```text
http://home
```

Функция `setLinkHome` принимает указатель на строку и меняет значение переменной, на которую этот указатель указывает.

Важно: строка в Go immutable, но мы не меняем содержимое старой строки. Мы меняем значение переменной `link`, чтобы она указывала на другую строку.

</details>

---

## Вопрос 2. Что содержится в `i`?

```go
var ptr *int
i := 10
ptr = &i
*ptr++
```

<details>
<summary><strong>Ответ</strong></summary>

В Go выражение:

```go
*ptr++
```

означает увеличение значения, на которое указывает `ptr`.

То есть это эквивалентно:

```go
(*ptr)++
```

Итог:

```go
i == 11
```

</details>

---

## Вопрос 3. Указатель на массив vs указатель на слайс

```go
arr := [3]int{1, 2, 3}
s := arr[:]
```

Можно ли получить указатель на массив? А указатель на слайс? В чем разница?

<details>
<summary><strong>Ответ</strong></summary>

Массив в Go — value type.

Слайс — descriptor, который содержит указатель на underlying array, длину и capacity.

Указатель на массив:

```go
fmt.Printf("%p\n", &arr)
```

Адрес первого элемента слайса:

```go
fmt.Printf("%p\n", &s[0])
```

Указатель на сам slice descriptor:

```go
fmt.Printf("%p\n", &s)
```

Важно различать:

- `&arr` — указатель на массив;
- `&s` — указатель на переменную-слайс;
- `&s[0]` — указатель на первый элемент underlying array.

</details>

---

## Вопрос 4. Что будет выведено?

```go
func modifyValue(x *int) {
    *x = 5
}

func main() {
    var num int = 2
    modifyValue(&num)
    fmt.Println(num)
}
```

<details>
<summary><strong>Ответ</strong></summary>

Будет выведено:

```text
5
```

В функцию передается адрес переменной `num`. Внутри функции происходит разыменование указателя и изменение значения по этому адресу.

</details>

---

## Вопрос 5. Как Go управляется с указателями в контексте GC?

<details>
<summary><strong>Ответ</strong></summary>

В Go есть garbage collector, который автоматически освобождает память объектов, на которые больше нет ссылок.

Go не использует ручное освобождение памяти, как `free` в C.

GC отслеживает достижимость объектов. Если объект больше недостижим из корней программы, он может быть очищен.

Если у вас есть указатель на большой объект, этот объект будет считаться живым, пока указатель остается достижимым.

</details>

---

## Вопрос 6. Что такое двойное разыменование в связном списке?

```go
type Node struct {
    value int
    next  *Node
}

first := &Node{value: 1}
second := &Node{value: 2}
first.next = second

fmt.Println(first.next.value)
```

<details>
<summary><strong>Ответ</strong></summary>

`first` — указатель на `Node`.

`first.next` — указатель на следующий `Node`.

`first.next.value` — обращение к полю `value` следующего узла.

Go автоматически разыменовывает указатели при доступе к полям структуры, поэтому не нужно писать:

```go
(*first).next.value
```

Итоговый вывод:

```text
2
```

</details>

---

## Вопрос 7. Будет ли напечатан `ok`?

```go
func main() {
    defer func() {
        recover()
    }()

    panic("test panic")
    fmt.Println("ok")
}
```

<details>
<summary><strong>Ответ</strong></summary>

Нет, `ok` не будет напечатан.

После вызова `panic` выполнение текущей функции останавливается, затем начинают выполняться отложенные функции `defer`.

`recover()` остановит панику, но выполнение не вернется к строке после `panic`.

</details>

---

## Вопрос 8. Исправьте код с `WaitGroup`

Функция должна вывести:

```text
one
two
three
Done!
```

Первые три строки могут быть в любом порядке, но `Done!` обязательно должно быть в конце.

```go
func printText(data []string) {
    wg := sync.WaitGroup{}

    for _, v := range data {
        go func(v string) {
            wg.Add(1)
            fmt.Println(v)
            wg.Done()
        }()
    }

    fmt.Println("done!")
}

func main() {
    data := []string{"one", "two", "three"}
    printText(data)
}
```

<details>
<summary><strong>Ответ</strong></summary>

Проблемы:

- `wg.Add(1)` вызывается внутри горутины;
- основная функция не ждет завершения горутин;
- параметр `v` не передается в анонимную функцию;
- `Done!` печатается сразу.

Исправленный вариант:

```go
func printText(data []string) {
    var wg sync.WaitGroup

    for _, v := range data {
        wg.Add(1)

        go func(v string) {
            defer wg.Done()
            fmt.Println(v)
        }(v)
    }

    wg.Wait()
    fmt.Println("Done!")
}

func main() {
    data := []string{"one", "two", "three"}
    printText(data)
}
```

</details>

---

## Вопрос 9. Что может пойти не так при подсчете операций?

```go
var callCounter uint

func main() {
    for i := 0; i < 10000; i++ {
        go func() {
            time.Sleep(time.Second)
            callCounter++
        }()
    }

    fmt.Println("Call counter value =", callCounter)
}
```

<details>
<summary><strong>Ответ</strong></summary>

Проблемы:

1. `callCounter++` не атомарная операция.
2. Несколько горутин могут одновременно читать и писать в одну переменную.
3. Возникает data race.
4. `fmt.Println` выполнится раньше, чем большинство горутин завершится.

Можно исправить через `sync.WaitGroup` и `atomic`:

```go
var callCounter atomic.Uint64

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 10000; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()
            time.Sleep(time.Second)
            callCounter.Add(1)
        }()
    }

    wg.Wait()
    fmt.Println("Call counter value =", callCounter.Load())
}
```

</details>

---

## Вопрос 10. Какие недостатки у кода с `context` и timeout?

```go
func (s *Service) ProcessData(timeoutCtx context.Context, r io.Reader) error {
    errCh := make(chan error)

    go func() {
        errCh <- s.processDataInternal(r)
    }()

    select {
    case err := <-errCh:
        return err
    case <-timeoutCtx.Done():
        return timeoutCtx.Err()
    }
}
```

<details>
<summary><strong>Ответ</strong></summary>

Проблемы:

1. Если `timeoutCtx` завершится раньше, `ProcessData` вернет ошибку, но горутина продолжит работать.
2. Если `processDataInternal` завершится после timeout, отправка в `errCh` заблокируется, потому что читать из канала уже некому.
3. `processDataInternal` не получает `context.Context`, поэтому не может корректно остановиться.

Улучшенный вариант:

```go
func (s *Service) ProcessData(ctx context.Context, r io.Reader) error {
    errCh := make(chan error, 1)

    go func() {
        errCh <- s.processDataInternal(ctx, r)
    }()

    select {
    case err := <-errCh:
        return err
    case <-ctx.Done():
        return ctx.Err()
    }
}
```

И важно, чтобы `processDataInternal` тоже уважала context:

```go
func (s *Service) processDataInternal(ctx context.Context, r io.Reader) error {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
    }

    // long-running work
    return nil
}
```

Буферизированный канал `errCh := make(chan error, 1)` нужен, чтобы горутина не зависла при отправке ошибки, если основной код уже вышел по timeout.

</details>

---

## Вопрос 11. Что выведет программа?

```go
func a() {
    x := []int{}
    x = append(x, 0)
    x = append(x, 1)
    x = append(x, 2)
    y := append(x, 3)
    z := append(x, 4)
    fmt.Println(y, z)
}

func main() {
    a()
}
```

<details>
<summary><strong>Ответ</strong></summary>

Разберем по шагам:

```go
x := []int{}
x = append(x, 0) // len=1, cap=1, x=[0]
x = append(x, 1) // len=2, cap=2, x=[0 1]
x = append(x, 2) // len=3, cap=4, x=[0 1 2]
```

У `x` есть свободная capacity.

```go
y := append(x, 3)
```

`y` использует тот же underlying array:

```go
y == [0 1 2 3]
```

Затем:

```go
z := append(x, 4)
```

`z` снова пишет в тот же underlying array на позицию после `x`.

Значение `3` будет перезаписано на `4`.

Вывод:

```text
[0 1 2 4] [0 1 2 4]
```

</details>

---

## Вопрос 12. Что выведет код со строкой?

```go
s := "test"
println(s[0])

var newS string = "R"
counter := 0
for _, item := range s {
    counter++
    if counter == 1 {
        continue
    }
    newS = strings.Join([]string{newS, string(item)}, "")
}
println(newS)
```

<details>
<summary><strong>Ответ</strong></summary>

```go
println(s[0])
```

Выведет ASCII-код первого байта строки:

```text
116
```

Потому что `'t'` имеет код `116`.

Далее код собирает новую строку, пропуская первый символ `test` и добавляя `R` в начало.

Итог:

```text
Rest
```

Важно: строки в Go immutable. Нельзя сделать так:

```go
s[0] = 'R'
```

</details>

---

## Вопрос 13. Как ускорить SQL-запрос по slowlog?

DevOps говорит, что в slowlog есть запрос, который выполняется дольше 10 секунд.

Запрос выбирает клиентов определенного менеджера, у которых указаны два этапа сделки.

```sql
SELECT m.*
FROM members m
LEFT JOIN manager_clients mc ON m.user_id = mc.client_id
WHERE mc.manager_id = '152734'
  AND m.user_id IN (
      SELECT client_id
      FROM client_deal_phases cdp
      WHERE cdp.phase_id IN (45, 47)
      GROUP BY client_id
      HAVING count(client_id) = 2
  );
```

`EXPLAIN`:

| id | select_type        | table | type   | possible_keys                   | key                             | rows   | Extra                    |
|----|--------------------|-------|--------|---------------------------------|---------------------------------|--------|--------------------------|
| 1  | PRIMARY            | mc    | ref    | idx_manager_id_client_id_uindex | idx_manager_id_client_id_uindex | 1      | Using where; Using index |
| 1  | PRIMARY            | m     | eq_ref | idx_user_id                     | idx_user_id                     | 1      | Using where              |
| 2  | DEPENDENT SUBQUERY | cdp   | index  | idx_client_id                   | idx_client_id                   | 189480 | Using where              |

<details>
<summary><strong>Ответ</strong></summary>

Главная проблема — `DEPENDENT SUBQUERY`.

Такой подзапрос может выполняться многократно для строк из внешнего запроса. Кроме того, по таблице `client_deal_phases` просматривается много строк.

Можно переписать запрос через `JOIN`:

```sql
SELECT m.*
FROM members m
JOIN manager_clients mc
  ON m.user_id = mc.client_id
JOIN (
    SELECT client_id
    FROM client_deal_phases
    WHERE phase_id IN (45, 47)
    GROUP BY client_id
    HAVING COUNT(DISTINCT phase_id) = 2
) cdp
  ON m.user_id = cdp.client_id
WHERE mc.manager_id = '152734';
```

Также стоит проверить индексы.

Например, для `client_deal_phases` может быть полезен составной индекс:

```sql
CREATE INDEX idx_phase_id_client_id
ON client_deal_phases (phase_id, client_id);
```

Или, в зависимости от кардинальности и плана выполнения:

```sql
CREATE INDEX idx_client_id_phase_id
ON client_deal_phases (client_id, phase_id);
```

Важно проверять итоговый `EXPLAIN` после изменения запроса и индексов.

</details>

---

# Goroutines и Channels

## Вопрос 1. Блокируется ли этот код?

```go
ch := make(chan int)

go func() {
    <-ch
}()

ch <- 1
```

<details>
<summary><strong>Ответ</strong></summary>

Канал `ch` небуферизированный.

Отправка в небуферизированный канал блокируется до тех пор, пока другая горутина не будет готова прочитать значение.

В этом примере есть горутина, которая читает из канала:

```go
<-ch
```

Поэтому отправка:

```go
ch <- 1
```

разблокируется, когда чтение будет готово.

</details>

---

## Вопрос 2. Что произойдет после выполнения кода?

```go
ch := make(chan int)
close(ch)
ch <- 1
```

<details>
<summary><strong>Ответ</strong></summary>

Будет panic.

В закрытый канал нельзя отправлять значения.

Ошибка:

```text
panic: send on closed channel
```

Из закрытого канала можно читать. Если значений больше нет, чтение вернет zero value типа.

</details>

---

## Вопрос 3. Что содержится в `val`?

```go
ch1 := make(chan chan int)
ch2 := make(chan int)

go func() {
    ch2 <- 1
}()

ch1 <- ch2
val := <-ch1
```

<details>
<summary><strong>Ответ</strong></summary>

Этот код приведет к deadlock.

Почему:

```go
ch1 <- ch2
```

Это отправка в небуферизированный канал `ch1`. Она заблокируется, пока кто-то не начнет читать из `ch1`.

Но чтение:

```go
val := <-ch1
```

находится ниже и никогда не будет достигнуто, потому что main goroutine уже заблокировалась на отправке.

Также горутина с:

```go
ch2 <- 1
```

тоже заблокируется, потому что из `ch2` никто не читает.

Чтобы исправить, можно отправку в `ch1` сделать в отдельной горутине:

```go
go func() {
    ch1 <- ch2
}()

val := <-ch1
fmt.Println(<-val)
```

</details>

---

## Вопрос 4. Как узнать, что канал закрыт?

<details>
<summary><strong>Ответ</strong></summary>

При чтении из канала можно использовать второе значение:

```go
val, ok := <-ch
```

- `ok == true` — значение получено, канал открыт или в нем еще были значения;
- `ok == false` — канал закрыт и значений больше нет.

Пример:

```go
val, ok := <-ch
if !ok {
    fmt.Println("Канал закрыт")
    return
}

fmt.Println("Значение:", val)
```

Также можно использовать `range`:

```go
for val := range ch {
    fmt.Println(val)
}
```

Цикл завершится, когда канал будет закрыт и все значения будут прочитаны.

</details>

---

## Вопрос 5. Есть ли способ узнать, сколько элементов находится в канале?

<details>
<summary><strong>Ответ</strong></summary>

Да, для канала можно использовать `len`:

```go
len(ch)
```

Она вернет количество элементов, которые сейчас находятся в буфере канала.

Пример:

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2

fmt.Println(len(ch)) // 2
```

Но важно помнить: в конкурентной программе это значение может устареть сразу после чтения, потому что другая горутина может отправить или прочитать значение.

</details>

---

## Вопрос 6. Как избежать race condition в этом коде?

```go
counter := 0
var mu sync.Mutex

go func() {
    mu.Lock()
    counter++
    mu.Unlock()
}()

go func() {
    mu.Lock()
    counter++
    mu.Unlock()
}()
```

Какие еще методы синхронизации вы знаете?

<details>
<summary><strong>Ответ</strong></summary>

Синхронизация достигается с помощью `sync.Mutex`.

Перед изменением `counter` горутина захватывает lock:

```go
mu.Lock()
```

После изменения освобождает lock:

```go
mu.Unlock()
```

Лучше использовать `defer`, чтобы lock точно освободился:

```go
mu.Lock()
defer mu.Unlock()
counter++
```

Другие способы синхронизации:

- `sync.WaitGroup`;
- channels;
- `sync.RWMutex`;
- `sync.Once`;
- `sync.Cond`;
- `sync/atomic`;
- context cancellation.

</details>

---

## Вопрос 7. Что произойдет при выполнении кода?

```go
ch := make(chan int)
ch <- 1
```

Как решить проблему?

<details>
<summary><strong>Ответ</strong></summary>

Программа заблокируется.

`ch` — небуферизированный канал. Отправка в него требует, чтобы кто-то одновременно читал из канала.

Вариант 1: использовать буферизированный канал.

```go
ch := make(chan int, 1)
ch <- 1
```

Вариант 2: читать из канала в другой горутине.

```go
ch := make(chan int)

go func() {
    fmt.Println(<-ch)
}()

ch <- 1
```

</details>

---

## Вопрос 8. Как читать из двух input channels и писать в output channel?

```go
in1 := make(chan int)
in2 := make(chan int)
out := make(chan int)
```

<details>
<summary><strong>Ответ</strong></summary>

Можно использовать `for-select`:

```go
func merge(in1, in2 <-chan int, out chan<- int) {
    for in1 != nil || in2 != nil {
        select {
        case v, ok := <-in1:
            if !ok {
                in1 = nil
                continue
            }
            out <- v

        case v, ok := <-in2:
            if !ok {
                in2 = nil
                continue
            }
            out <- v
        }
    }

    close(out)
}
```

Почему мы присваиваем `nil` закрытому каналу?

Чтение из закрытого канала всегда готово, поэтому `select` будет постоянно выбирать этот case. Присваивание `nil` отключает этот case.

</details>

---

## Вопрос 9. Как определить, что канал закрыт, если `0` — допустимое значение?

```go
ch := make(chan int, 1)
ch <- 0
close(ch)
```

<details>
<summary><strong>Ответ</strong></summary>

Нужно использовать второе значение при чтении:

```go
val, ok := <-ch
```

Пример:

```go
val, ok := <-ch
if !ok {
    fmt.Println("Канал закрыт")
    return
}

fmt.Println("Получено значение:", val)
```

Если `val == 0`, это еще не значит, что канал закрыт. Нужно смотреть именно на `ok`.

</details>

---

## Вопрос 10. Как узнать количество элементов в буферизированном канале?

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
```

<details>
<summary><strong>Ответ</strong></summary>

Можно использовать:

```go
fmt.Println(len(ch))
```

В этом примере:

```go
len(ch) == 2
cap(ch) == 3
```

Но в конкурентной программе это значение является моментальным снимком. Оно может измениться сразу после вызова `len(ch)`.

</details>

---

# Работа со строками

## Вопрос 1. Как эффективно конкатенировать большое количество строк?

<details>
<summary><strong>Ответ</strong></summary>

Для большого количества строк лучше использовать `strings.Builder`.

Пример:

```go
var b strings.Builder

b.WriteString("hello")
b.WriteString(" ")
b.WriteString("world")

result := b.String()
```

Если заранее известен примерный размер результата, можно вызвать `Grow`:

```go
var b strings.Builder
b.Grow(1024)

for _, part := range parts {
    b.WriteString(part)
}

result := b.String()
```

Это помогает уменьшить количество выделений памяти.

</details>

---

## Вопрос 2. Чем отличаются `len` и `utf8.RuneCountInString`?

<details>
<summary><strong>Ответ</strong></summary>

`len(s)` возвращает количество байт в строке.

`utf8.RuneCountInString(s)` возвращает количество Unicode code points, то есть рун.

Пример:

```go
s := "привет"

fmt.Println(len(s))                    // количество байт
fmt.Println(utf8.RuneCountInString(s)) // количество рун
```

Кириллические символы в UTF-8 занимают больше одного байта, поэтому результаты будут отличаться.

Когда использовать:

- `len` — когда важен размер в байтах;
- `utf8.RuneCountInString` — когда нужно посчитать символы в Unicode-строке.

</details>

---

## Вопрос 3. Как сравниваются строки в Go?

Что учитывать при сравнении строк в разных языках и кодировках?

<details>
<summary><strong>Ответ</strong></summary>

Строки в Go можно сравнивать через оператор `==`:

```go
if a == b {
    fmt.Println("строки равны")
}
```

Также есть функция:

```go
strings.Compare(a, b)
```

Она возвращает:

- `0`, если строки равны;
- `-1`, если `a < b`;
- `1`, если `a > b`.

Важно: `strings.Compare` не возвращает `true` или `false`.

Что учитывать:

- строки сравниваются по байтам;
- визуально одинаковые Unicode-строки могут иметь разное представление;
- для сложных случаев может понадобиться Unicode normalization.

Пример проблемы:

```go
s1 := "é"        // один Unicode code point
s2 := "e\u0301" // e + combining accent
```

Визуально строки похожи, но байтовое представление может отличаться.

</details>

---

## Вопрос 4. Как извлечь подстроку из строки?

Как обработать строки с многобайтовыми символами?

<details>
<summary><strong>Ответ</strong></summary>

Если строка содержит только ASCII, можно использовать срез по байтам:

```go
s := "hello"
sub := s[1:4] // "ell"
```

Но для Unicode-строк так делать опасно, потому что можно разрезать символ посередине.

Для Unicode лучше преобразовать строку в `[]rune`:

```go
s := "привет"
runes := []rune(s)
sub := string(runes[1:4])

fmt.Println(sub)
```

Важно: в стандартной библиотеке Go нет функции `strings.Substring`.

</details>

---

## Вопрос 5. Как преобразовать строку в число и число в строку?

<details>
<summary><strong>Ответ</strong></summary>

Для преобразования строки в число можно использовать `strconv.Atoi`:

```go
n, err := strconv.Atoi("42")
if err != nil {
    return err
}

fmt.Println(n)
```

Для преобразования числа в строку:

```go
s := strconv.Itoa(42)
fmt.Println(s)
```

Для других типов есть функции:

```go
strconv.ParseInt
strconv.ParseFloat
strconv.FormatInt
strconv.FormatFloat
```

Ошибки нужно обрабатывать, потому что строка может быть некорректной:

```go
n, err := strconv.Atoi("not-a-number")
if err != nil {
    fmt.Println("Некорректное число:", err)
    return
}
```

</details>

---

## Бонусный вопрос. Как безопасно обрабатывать строки в веб-сервере?

Как предотвращать XSS и SQL injection?

<details>
<summary><strong>Ответ</strong></summary>

Для безопасной обработки пользовательского ввода важно не пытаться вручную экранировать все подряд, а использовать правильные инструменты для конкретного контекста.

### SQL injection

Не собирайте SQL через конкатенацию строк:

```go
query := "SELECT * FROM users WHERE name = '" + name + "'"
```

Используйте параметры:

```go
rows, err := db.QueryContext(ctx,
    "SELECT * FROM users WHERE name = ?",
    name,
)
```

### XSS

Для HTML используйте `html/template`, а не `text/template`.

```go
tmpl, err := template.ParseFiles("page.html")
if err != nil {
    return err
}

err = tmpl.Execute(w, data)
```

`html/template` учитывает HTML-контекст и помогает безопасно экранировать данные.

### Общие рекомендации

- валидируйте входные данные;
- используйте prepared statements;
- не доверяйте пользовательскому вводу;
- кодируйте данные под конкретный контекст: HTML, URL, JSON, SQL;
- ограничивайте размер входных данных;
- логируйте аккуратно, чтобы не сохранять секреты.

</details>

---

## Вклад ❤️

Если у вас есть интересные вопросы, идеи для улучшения или вы нашли ошибку, создавайте Pull Request или Issue.

Вместе мы сделаем этот ресурс полезнее для всех, кто готовится к собеседованиям на Go-разработчика.

Перед PR желательно:

- проверить форматирование Markdown;
- добавить минимальный пример кода, если вопрос технический;
- объяснить не только «что будет выведено», но и почему;
- по возможности указать corner cases.

---

## Лицензия

MIT License. См. файл [`LICENSE`](./LICENSE) для деталей.

---

<p align="center">
  Удачи на собеседованиях! 🍀
</p>

