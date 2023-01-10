package main

import (
	"fmt"
)

func main() {
	var text, res string
	var width int
	fmt.Scanf("%s %d", &text, &width)
	if len(text) <= width {
		res = text
	} else {
		str := []byte(text)
		str = str[:width]
		res = string(str) + "..."
	}
	fmt.Println(res)
}
