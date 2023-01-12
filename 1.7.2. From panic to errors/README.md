## 1.7.2. От паники к ошибкам

Я написал программу, которая считывает исходную строку вида:

>balance/overdraft t1 t2 t3 ... tn

И превращает ее в счет типа `account` и список транзакций типа `[]int`.

К сожалению, я поленился явно обработать ошибки, положившись на отлов паники через `defer` и `recover()`. Исправьте программу, заменив их на явный возврат и обработку ошибок.

В получившейся программе не должно быть конструкций `defer` и `recover()`. Если в исходной строке есть ошибка — программа должна вывести только ошибку. Иначе — вывести счет и транзакции.

Например:

* **на входе:** 80/10 10 -20 30
* **на выходе:** -> {80 10} [10 -20 30]

Или:

* **на входе:** 80/10 10 z 30
* **на выходе:** -> strconv.Atoi: parsing "z": invalid syntax

Если в исходной строке несколько ошибок, программа должна вывести только первую.

Не меняйте текст ошибок, используйте те же строки, что принимали вызовы `panic()` в исходной программе.

Гарантируется, что общий формат исходной строки соблюдается (то есть значения разделены пробелами, овердрафт отделен от баланса дробью, и тому подобное), но могут быть ошибки в отдельных значениях (как в примере выше).

Программа не должна завершаться паникой или `os.Exit(1)` ни при каких обстоятельствах.

___
**Напишите программу. Тестируется через stdin → stdout**

**Time Limit:** 8 секунд

**Memory Limit:** 256 MB
___
**Sample Input:**
> **80/10 10 -20 30**

**Sample Output:**
> **-> {80 10} [10 -20 30]**
___

```Go
package main

// не меняйте импорты, они нужны для проверки
import (
    "bufio"
    "errors"
    "fmt"
    "io/ioutil"
    "os"
    "reflect"
    "runtime"
    "strconv"
    "strings"
)

// account представляет счет
type account struct {
    balance   int
    overdraft int
}

func main() {
    var acc account
    var trans []int
    defer func() {
        fmt.Print("-> ")
        err := recover()
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(acc, trans)
    }()
    acc, trans = parseInput()
}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput() (account, []int) {
    accSrc, transSrc := readInput()
    acc := parseAccount(accSrc)
    trans := parseTransactions(transSrc)
    return acc, trans
}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInput() (string, []string) {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    scanner.Scan()
    accSrc := scanner.Text()
    var transSrc []string
    for scanner.Scan() {
        transSrc = append(transSrc, scanner.Text())
    }
    return accSrc, transSrc
}

// parseAccount парсит счет из строки
// в формате balance/overdraft.
func parseAccount(src string) account {
    parts := strings.Split(src, "/")
    balance, _ := strconv.Atoi(parts[0])
    overdraft, _ := strconv.Atoi(parts[1])
    if overdraft < 0 {
        panic("expect overdraft >= 0")
    }
    if balance < -overdraft {
        panic("balance cannot exceed overdraft")
    }
    return account{balance, overdraft}
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactions(src []string) []int {
    trans := make([]int, len(src))
    for idx, s := range src {
        t, _ := strconv.Atoi(s)
        trans[idx] = t
    }
    return trans
}
```