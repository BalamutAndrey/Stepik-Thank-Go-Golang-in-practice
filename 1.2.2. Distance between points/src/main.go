package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)
	d := math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
	fmt.Println(d)
}
