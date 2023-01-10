package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func shuffle(nums []int) {
	f := func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}
	rand.Shuffle(len(nums), f)
}

func init() {
	rand.Seed(42)
}

func main() {
	nums := readInput()
	shuffle(nums)
	fmt.Println(nums)
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
