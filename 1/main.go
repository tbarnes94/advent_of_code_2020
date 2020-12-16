package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	arr := strings.Split(sc, "\n")

	// numbers
	a := 0
	b := 0
	c := 0
	// indices
	i := 0
	j := 0
	k := 0

	// part 1
	//for a + b != 2020 {
	//	a, _ = strconv.Atoi(arr[i])
	//	b, _ = strconv.Atoi(arr[j])
	//	if j == len(arr)-1 {
	//		j = 0
	//		i++
	//	} else {
	//		j++
	//	}
	//}

	// part 2
	for a + b + c != 2020 {
		a, _ = strconv.Atoi(arr[i])
		b, _ = strconv.Atoi(arr[j])
		c, _ = strconv.Atoi(arr[k])
		if k == len(arr)-1 {
			k = 0
			if j == len(arr)-1 {
				j = 0
				i++
			} else {
				j++
			}
		} else {
			k++
		}
	}


	fmt.Println("a: "+strconv.Itoa(a), "b: "+strconv.Itoa(b), "c: "+strconv.Itoa(c), a + b + c, a * b * c)
}
