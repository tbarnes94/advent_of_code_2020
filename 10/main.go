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
		if s == "" {
			break
		}
		n, _ := strconv.Atoi(s)
		out = append(out, n)
	}
	return out
}

func getNJoltDiffs(nums []int, n int) int {
	ctr := 0
	j := 1
	for i, _ := range nums {
		if j==len(nums) {
			break
		}
		if nums[j]-nums[i] == n || nums[j]-nums[i] == -n {
			ctr++
		}
		j++
	}
	return ctr
}

func max(nums []int) int {
	out := 0
	for _, n := range nums {
		if n > out {
			out = n
		}
	}
	return out
}

func contains(list []int, value int) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// inspired by https://www.reddit.com/r/adventofcode/comments/ka8z8x/2020_day_10_solutions/gfcxuxf
func getPaths(adapters []int) int {
	goal := max(adapters)
	paths := make([]int, goal+1)
	validJoltDiffs := []int{1, 2, 3}
	paths[0] = 1
	for _, adapter := range adapters {
		for _, diff := range validJoltDiffs {
			nextAdapter := adapter + diff
			if contains(adapters, nextAdapter) {
				paths[nextAdapter] += paths[adapter]
			}
		}
	}
	return paths[goal]
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	arr := strings.Split(sc, "\n")
	jolts := parseInts(arr)
	jolts = append(jolts, 0, max(jolts)+3)
	sort.Ints(jolts)

	// part 1
	fmt.Println("Number of 1 Jolt Diffs * Number of 3 Jolt Diffs:", getNJoltDiffs(jolts, 1) * getNJoltDiffs(jolts, 3))

	// part 2
	fmt.Println(fmt.Sprintf("Distinct Adapter Configurations for an Adapter with %d Joltage:", max(jolts)), getPaths(jolts))
}
