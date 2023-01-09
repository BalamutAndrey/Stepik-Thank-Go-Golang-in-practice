## 1.2.3. Повтор строки

Программа принимает на вход строку `source` и число `times`. Требуется склеить `source` саму с собой `times` раз и вернуть результат:

**`source = x`, `times = 3` → `xxx`**
**`source = omm`, `times = 2` → `ommomm`**

___
**Напишите программу. Тестируется через stdin → stdout**

**Time Limit:** 8 секунд

**Memory Limit:** 256 MB
___
**Sample Input:**
> **a 5**

**Sample Output:**
> **aaaaa**
___

```Go
package main

import (
	"fmt"
)

func main() {
	var source string
	var times int
	// гарантируется, что значения корректные
	fmt.Scan(&source, &times)

	// возьмите строку `source` и повторите ее `times` раз
	// запишите результат в `result`
    // ...

	fmt.Println(result)
}
```