package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	arr := strings.Split(sc, "\n")

	// part 1
	//ctr := 0
	//for _, line := range arr {
	//	if line == "" {
	//		break
	//	}
	//	pw := strings.Split(line, " ")
	//	ndx := strings.Split(pw[0], "-")
	//	start, _ := strconv.Atoi(ndx[0])
	//	end, _ := strconv.Atoi(ndx[1])
	//	rule := pw[1][:len(pw[1])-1]
	//	body := pw[2]
	//	count := 0
	//	for _, c := range body {
	//		if string(c) == rule {
	//			count++
	//		}
	//	}
	//	matched := count >= start && count <= end
	//	if matched {
	//		ctr++
	//	}
	//
	//	//fmt.Println(start, end, rule, body, matched)
	//}

	// part 2
	ctr := 0
	for _, line := range arr {
		if line == "" {
			break
		}
		pw := strings.Split(line, " ")
		ndx := strings.Split(pw[0], "-")
		start, _ := strconv.Atoi(ndx[0])
		end, _ := strconv.Atoi(ndx[1])
		rule := pw[1][:len(pw[1])-1]
		body := pw[2]

		matched := string(body[start-1]) == rule && string(body[end-1]) != rule ||
			string(body[start-1]) != rule && string(body[end-1]) == rule
		if matched {
			ctr++
		}

		fmt.Println(start, end, rule, body, matched)
	}

	fmt.Println("Number of Valid Passwords: ", ctr)
}
