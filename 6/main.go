package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func distinctString(s string) string {
	seen := make(map[string]bool)
	out := ""
	for _, c := range s {
		if !seen[string(c)] {
			seen[string(c)] = true
			out += string(c)
		}
	}
	return out
}

func consensusString(s string, numPeople int) string {
	seen := make(map[string]int)
	out := ""
	for _, c := range s {
		seen[string(c)]++
	}
	for ndx := range seen {
		if seen[ndx] == numPeople {
			out += ndx
		}
	}
	return out
}

//func sortString(s string) string {
//	arr := strings.Split(s, "")
//	sort.Strings(arr)
//	return strings.Join(arr, "")
//}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	groups := strings.Split(sc, "\n\n")

	// part 1
	count := 0
	for _, g := range groups {
		g = strings.ReplaceAll(g, "\n", "")
		g = strings.ReplaceAll(g, " ", "")
		count += len(distinctString(g))
	}
	fmt.Println("Total Count of Distinct Yes Answers:", count)

	// part 2
	count = 0
	for _, g := range groups {
		g = strings.ReplaceAll(g, " ", "")
		people := strings.Split(g, "\n")
		//fmt.Println(fmt.Sprintf("\nGroup #%d:", i+1))
		for _, p := range people {
			if p == "" {
				people = people[:len(people)-1]
				break
			}
			//fmt.Println(fmt.Sprintf("Person #%d: [ %s ]", j+1, p))
		}
		g = strings.ReplaceAll(g, "\n", "")
		sharedYeses := consensusString(g, len(people))
		//fmt.Println("Shared Yeses:", sharedYeses, "Count:", len(sharedYeses))
		count += len(sharedYeses)
	}
	fmt.Println("Total Count of Group-Shared Yes Answers:", count)
}
