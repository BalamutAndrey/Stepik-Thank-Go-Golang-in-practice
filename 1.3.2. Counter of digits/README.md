## 1.3.2. Счетчик цифр

Напишите программу, которая считает, сколько раз каждая цифра встречается в числе. Гарантируется, что на вход подаются только положительные целые числа, не выходящие за диапазон `int`.
___
**Напишите программу. Тестируется через stdin → stdout**

**Time Limit:** 8 секунд

**Memory Limit:** 256 MB
___
**Sample Input:**
> **12823**

**Sample Output:**
> **1:1 2:2 3:1 8:1**
___

```Go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var number int
	fmt.Scanf("%d", &number)

	// Посчитайте, сколько раз каждая цифра встречается
	// в числе `number`. Запишите результат в карту `counter`,
	// где ключом является цифра числа, а значением -
	// сколько раз она встречается
	counter := make(map[int]int)
	// ...

	printCounter(counter)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// printCounter печатает карту в формате
// key1:val1 key2:val2 ... keyN:valN
func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
```