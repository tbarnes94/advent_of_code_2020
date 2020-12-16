package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	arr := strings.Split(sc, "\n")

	// number of trees
	ctr := 0

	// vertical/horiztontal index in arr content
	y := 0
	ndx := 0

	// arrays of slopes to use
	run := []int{1, 3, 5, 7, 1}
	drop := []int{1, 1, 1, 1, 2}

	// index of run/drop arrays
	i := 0

	var trees []int
	for i < len(run) {
		ctr = 0
		ndx = 0
		y = 0
		for y < len(arr) {
			if len(arr[y]) == 0 {
				break
			}
			if string(arr[y][ndx]) == "#" {
				ctr++
			}
			ndx = (ndx + run[i]) % len(arr[ndx])
			y = y + drop[i]
		}
		trees = append(trees, ctr)
		i++
	}
	out := 1
	for _, t := range trees {
		out = out * t
	}

	fmt.Println("Number of Trees Encountered:", out, trees)
}
