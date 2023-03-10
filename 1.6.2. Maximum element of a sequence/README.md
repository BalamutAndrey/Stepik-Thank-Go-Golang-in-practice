## 1.6.2. Максимальный элемент последовательности

Определите интерфейс универсального итератора (``iterator``), который можно использовать в функции выбора максимального элемента (``max``). Реализуйте интерфейс для итератора по срезу целых чисел.

Подробности — по коду задания.

___
**Напишите программу. Тестируется через stdin → stdout**

**Time Limit:** 8 секунд

**Memory Limit:** 256 MB
___
**Sample Input:**
> **1 4 5 2 3**

**Sample Output:**
> **5**
___

```Go
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

// element - интерфейс элемента последовательности
type element interface{}

// weightFunc - функция, которая возвращает вес элемента
type weightFunc func(element) int

// iterator - интерфейс, который умеет
// поэлементно перебирать последовательность
type iterator interface {
    // чтобы понять сигнатуры методов - посмотрите,
    // как они используются в функции max() ниже
}

// intIterator - итератор по целым числам
// (реализует интерфейс iterator)
type intIterator struct {
    // поля структуры
}

// методы intIterator, которые реализуют интерфейс iterator

// конструктор intIterator
func newIntIterator(src []int) *intIterator {
    // ...
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// main находит максимальное число из переданных на вход программы.
func main() {
    nums := readInput()
    it := newIntIterator(nums)
    weight := func(el element) int {
        return el.(int)
    }
    m := max(it, weight)
    fmt.Println(m)
}

// max возвращает максимальный элемент в последовательности.
// Для сравнения элементов используется вес, который возвращает
// функция weight.
func max(it iterator, weight weightFunc) element {
    var maxEl element = it.val()
    for it.next() {
        curr := it.val()
        if maxEl == nil || weight(curr) > weight(maxEl) {
            maxEl = curr
        }
    }
    return maxEl
}

// readInput считывает последовательность целых чисел из os.Stdin.
func readInput() []int {
    var nums []int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        nums = append(nums, num)
    }
    return nums
}
```