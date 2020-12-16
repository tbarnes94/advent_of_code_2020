package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	arr := strings.Split(sc, "\n\n")

	fmt.Println(arr)
}
