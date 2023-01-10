package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	var res []int
	for i := range iterable {
		if predicate(iterable[i]) {
			res = append(res, iterable[i])
		}
	}
	return res
}

func main() {
	src := readInput()
	res := filter(func(i int) bool { return i%2 == 0 }, src)
	fmt.Println(res)
}

func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
