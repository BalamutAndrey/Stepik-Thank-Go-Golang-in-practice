## 1.4.2. Перетасовка

Напишите функцию `shuffle()`, которая тасует элементы среза в случайном порядке. Функция должна отрабатывать in-place, то есть менять содержимое переданного среза, а не создавать новый срез. Чтобы перетасовать элементы, используйте функцию [rand.Shuffle()](https://golang.org/pkg/math/rand/#Shuffle):

>func Shuffle(n int, swap func(i, j int))

Гарантируется, что на вход подаются только целые числа.

Обратите внимание на комментарий к функции `init()` — на реальных задачах она может вам пригодиться.

___
**Напишите программу. Тестируется через stdin → stdout**

**Time Limit:** 8 секунд

**Memory Limit:** 256 MB
___
**Sample Input:**
> **1 2 3 4 5 6**

**Sample Output:**
> **[2 4 5 6 1 3]**
___

```Go
package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
)

// shuffle перемешивает элементы nums in-place.
func shuffle(nums []int) {
    // перетасуйте nums с помощью rand.Shuffle()
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// init - это специальная функция, которую Go вызывает до main()
// обычно используется для инициализации начального состояния приложения
func init() {
    // Функция rand.Seed() инициализирует генератор случайных чисел
    // здесь мы используем константу 42, чтобы программу
    // можно было проверить тестами.
    //
    // В реальных задачах не используйте константы!
    // Используйте, например, время в наносекундах:
    // rand.Seed(time.Now().UnixNano())
    rand.Seed(42)
}

func main() {
    nums := readInput()
    shuffle(nums)
    fmt.Println(nums)
}

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
    var nums []int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        num, _ := strconv.Atoi(scanner.Text())
        nums = append(nums, num)
    }
    return nums
}
```