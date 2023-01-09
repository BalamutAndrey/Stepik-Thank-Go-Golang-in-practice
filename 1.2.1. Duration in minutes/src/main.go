package main

import (
	"fmt"
	"time"
)

func main() {
	var s string
	fmt.Scan(&s)
	d, _ := time.ParseDuration(s)
	fmt.Println(s, "=", d.Minutes(), "min")
}
