package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func validByr(byr string) bool {
	birthYear, err := strconv.Atoi(byr)
	if err != nil {
		return false
	}
	return birthYear >= 1920 && birthYear <= 2002
}

func validIyr(iyr string) bool {
	issueYear, err := strconv.Atoi(iyr)
	if err != nil {
		return false
	}
	return issueYear >= 2010 && issueYear <= 2020
}

func validEyr(eyr string) bool {
	expYear, err := strconv.Atoi(eyr)
	if err != nil {
		return false
	}
	return expYear >= 2020 && expYear <= 2030
}

func validHgt(hgt string) bool {
	number := hgt[:len(hgt)-2]
	unit := hgt[len(hgt)-2:]
	num, err := strconv.Atoi(number)
	if err != nil {
		return false
	}
	if unit == "in" {
		return num >= 59 && num <= 76
	} else if unit == "cm" {
		return num >= 150 && num <= 193
	}
	return false
}

func validHcl(hcl string) bool {
	matched, err := regexp.MatchString("#[A-Za-z0-9]{6}", hcl)
	if err != nil {
		return false
	}
	return matched
}

func validEcl(ecl string) bool {
	if len(ecl) != 3 {
		return false
	}
	mEyeColors := []string {
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth",
	}
	for _, c := range mEyeColors {
		if ecl == c {
			return true
		}
	}
	return false
}

func validPid(pid string) bool {
	if len(pid) != 9 {
		return false
	}
	_, err := strconv.Atoi(pid)
	if err != nil {
		return false
	}
	return true
}


func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	passports := strings.Split(sc, "\n\n")

	// number of valid passports
	numValid := 0
	for _, p := range passports {
		// replace newlines with whitespace for reliable split
		p = strings.ReplaceAll(p, "\n", " ")
		// split on whitespace
		pairs := strings.Split(p, " ")

		// passport map
		mPassports := make(map[string]string)
		for _, l := range pairs {
			kv := strings.Split(l, ":")
			if len(kv) < 2 {
				break
			}
			mPassports[kv[0]] = kv[1]
		}

		// validation
		validCtr := 0
		for k, v := range mPassports {
			switch k {
			case "byr":
				if !validByr(v) {
					continue
				}
				validCtr++
			case "iyr":
				if !validIyr(v) {
					continue
				}
				validCtr++
			case "eyr":
				if !validEyr(v) {
					continue
				}
				validCtr++
			case "hgt":
				if !validHgt(v) {
					continue
				}
				validCtr++
			case "hcl":
				if !validHcl(v) {
					continue
				}
				validCtr++
			case "ecl":
				if !validEcl(v) {
					continue
				}
				validCtr++
			case "pid":
				if !validPid(v) {
					continue
				}
				validCtr++
			}
		}
		// if we have 7 and cid is the key missing a value, assume valid
		if validCtr == 7 {
			numValid++
		}
	}
	fmt.Println("Number of Valid Passports:", numValid)
}
