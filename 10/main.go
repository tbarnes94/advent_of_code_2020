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

func getDistinctConfigs(jolts []int, thisVal int, thisNdx int, cache []int) int {
	out := 0
	if thisNdx>0 && thisNdx<len(jolts) && cache[thisNdx]>0 {
		return cache[thisNdx]
	}
	for nextNdx:=thisNdx+1; nextNdx<=nextNdx+3 && nextNdx<len(jolts); nextNdx++ {
		nextVal := jolts[nextNdx]
		if nextVal-thisVal<=3 {
			cachedVal := getDistinctConfigs(jolts, nextVal, nextNdx, cache)
			cache[nextNdx] = cachedVal
			out += cachedVal
		}

	}
	return out
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	arr := strings.Split(sc, "\n")
	jolts := parseInts(arr)
	jolts = append(jolts, 0, max(jolts)+3)
	sort.Ints(jolts)

	// part 1
	out := getNJoltDiffs(jolts, 1) * getNJoltDiffs(jolts, 3)
	fmt.Println("Number of 1 Jolt Diffs * Number of 3 Jolt Diffs:", out)

	// part 2
	jolts = jolts[1:]
	cache := make([]int, len(jolts))
	cache[len(jolts)-1] = 1
	fmt.Println("Distinct Adapter Configurations:", getDistinctConfigs(jolts, 0, -1, cache))
}
