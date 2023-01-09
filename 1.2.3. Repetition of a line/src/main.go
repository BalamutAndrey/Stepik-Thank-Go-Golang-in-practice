package main

import (
	"fmt"
)

func main() {
	var source, result string
	var times int
	fmt.Scan(&source, &times)
	for i := 0; i < times; i++ {
		result += source
	}
	fmt.Println(result)
}
