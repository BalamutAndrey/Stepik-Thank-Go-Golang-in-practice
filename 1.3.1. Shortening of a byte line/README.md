## 1.3.1. Укорот байтовой строки

Напишите программу, которая укорачивает строку до указанной длины и добавляет в конце многоточие:

`text` = Eyjafjallajokull, `width` = 6 **→ Eyjafj...**

Если строка не превышает указанной длины, менять ее не следует:

`text` = hello, `width` = 6 **→ hello**

Гарантируется, что в исходной строке `text` используются только однобайтовые символы без пробелов, а длина `width` строго больше 0.
___
**Напишите программу. Тестируется через stdin → stdout**

**Time Limit:** 8 секунд

**Memory Limit:** 256 MB
___
**Sample Input:**
> **Eyjafjallajokull**

**Sample Output:**
> **Eyjafj...**
___

```Go
package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...

	fmt.Println(res)
}
```