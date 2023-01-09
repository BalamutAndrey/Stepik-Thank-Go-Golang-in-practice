## 1.2.4. Язык по коду

Напишите программу, которая определяет название языка по его коду. Правила:

**`en` →** English

**`fr` →** French

**`ru` или `rus` →** Russian

**иначе →** Unknown

___
**Напишите программу. Тестируется через stdin → stdout**

**Time Limit:** 8 секунд

**Memory Limit:** 256 MB
___
**Sample Input:**
> **en**

**Sample Output:**
> **English**
___

```Go
package main

import (
	"fmt"
)

func main() {
	var code string
	fmt.Scan(&code)

	// определите полное название языка по его коду
	// и запишите его в переменную `lang`
	// ...

	fmt.Println(lang)
}
```