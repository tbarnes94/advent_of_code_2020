package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func isolateColor(s string) string {
	s = strings.ReplaceAll(s, " bags", "")
	s = strings.ReplaceAll(s, " bag", "")
	return s
}

type Color struct {
	Value int
	Name string
}

func parseRules(bags []string) map[string][]Color {
	rules := make(map[string][]Color)
	for _, bag := range bags {
		if bag == "" {
			break
		}
		bag = strings.ReplaceAll(bag, "contain", "")
		bag = strings.ReplaceAll(bag, ".", "")
		kv := strings.Split(bag, "  ")
		bagColor := isolateColor(kv[0])
		var bagContent []Color
		for _, item := range strings.Split(kv[1], ", ") {
			parts := strings.Split(item, " ")
			value, _ := strconv.Atoi(parts[0])
			name := strings.Join(parts[1:len(parts)-1], " ")
			bagContent = append(bagContent, Color{
				Value: value,
				Name: name,
			})
		}
		rules[bagColor] = bagContent
	}
	return rules
}

func recursiveCheck(goalColor string, rules map[string][]Color, containsGoalColor map[string]struct{}) {
	for bagColor := range getContains(rules, goalColor) {
		containsGoalColor[bagColor] = struct{}{}
		recursiveCheck(bagColor, rules, containsGoalColor)
	}
}

// Returns all the bags that contain the given bag
func getContains(rules map[string][]Color, bag string) map[string]struct{} {
	containsGoalColor := make(map[string]struct{})
	for bagColor, bagContent := range rules {
		for _, item := range bagContent {
			if item.Name == bag {
				containsGoalColor[bagColor] = struct{}{}
			}
		}
	}
	return containsGoalColor
}

func getTotalNumberOfBags(goalColor string, rules map[string][]Color) int {
	ctr := 0
	for _, color := range rules[goalColor] {
		ctr += color.Value * (1 + getTotalNumberOfBags(color.Name, rules))
	}
	return ctr
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	bags := strings.Split(sc, "\n")
	rules := parseRules(bags)
	const GoalColor = "shiny gold"

	//part 1
	containsGoalColor := make(map[string]struct{})
	recursiveCheck(GoalColor, rules, containsGoalColor)
	fmt.Println(fmt.Sprintf("Number of bags that contain %s bags: %d", GoalColor, len(containsGoalColor)))

	//part 2
	total := getTotalNumberOfBags(GoalColor, rules)
	fmt.Println(fmt.Sprintf("Number of bags that a %s bag contains: %d", GoalColor, total))
}
