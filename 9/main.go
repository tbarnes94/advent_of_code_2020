package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func parseInts (strs []string) []int {
	var out []int
	for _, s := range strs {
		n, _ := strconv.Atoi(s)
		out = append(out, n)
	}
	return out
}

func validNumber(preamble []int, num int) bool {
	for i:=0; i<len(preamble); i++ {
		for j:=1; j<len(preamble); j++ {
			if preamble[i] + preamble[j] == num {
				return true
			}
		}
	}
	return false
}

func findInvalidNumber(numbers []int, preambleLength int) int {
	var invalidNum int
	for i, num := range numbers {
		if i < preambleLength {
			continue
		}
		if !validNumber(numbers[i-preambleLength:i], num) {
			invalidNum = num
			break
		}
	}
	return invalidNum
}

func findEncryptWeakness(numbers []int, invalidNum int) int {
	for i := range numbers {
		for j:=i+1; j<len(numbers); j++ {
			if check(numbers[i:j], invalidNum) {
				nums := numbers[i:j]
				sort.Ints(nums)
				return nums[0] + nums[len(nums)-1]
			}
		}
	}
	return 0
}

func check(numbers []int, invalidNum int) bool {
	var nums []int
	for _, num := range numbers {
		nums = append(nums, num)
		if sum(nums) == invalidNum {
			return true
		}
	}
	return false
}

func sum(n []int) int {
	o := 0
	for _, num := range n {
		o += num
	}
	return o
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	arr := strings.Split(sc, "\n")
	numbers := parseInts(arr)
	const PreambleLength int = 25

	// part 1 - invalid number
	invalidNum := findInvalidNumber(numbers, PreambleLength)
	fmt.Println("Invalid Number:", invalidNum)

	// part 2 - encryption weakness
	ew := findEncryptWeakness(numbers, invalidNum)
	fmt.Println("Encryption Weakness:", ew)
}
