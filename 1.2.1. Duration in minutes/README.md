## 1.2.1. Продолжительность в минутах

Напишите программу, которая считает количество минут во временном отрезке.
>**1h30m = 90 min
>300s = 5 min
>10m = 10 min**

Используйте для этого [time.Duration.Minutes()](https://golang.org/pkg/time/#Duration.Minutes) из стандартной библиотеки.
___
**Напишите программу. Тестируется через stdin → stdout**

**Time Limit:** 8 секунд
**Memory Limit:** 256 MB
___
**Sample Input:**
> **1h30m**

**Sample Output:**
> **1h30m = 90 min**
___

```Go
package main

import (
	"fmt"
	"time"
)

func main() {
	// считываем временной отрезок из os.Stdin
	// гарантируется, что значение корректное
	// не меняйте этот блок
	var s string
	fmt.Scan(&s)
	d, _ := time.ParseDuration(s)

	// выведите исходное значение
	// и количество минут в нем
	// в формате "исходное = X min"
	// используйте метод .Minutes() объекта d
	fmt.Println(/* ... */)
}
```