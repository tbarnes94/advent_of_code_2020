package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"math"
)

func getRow(rowCode string) int {
	start := 0
	end := 127
	for i, c := range rowCode {
		if string(c) == "B" {
			if string(rowCode[6]) == "F" {
				continue
			}
			start = start + int(math.Pow(2, float64(7-(i+1))))
			continue
		}
		if string(rowCode[6]) == "B" {
			continue
		}
		end = end - int(math.Pow(2, float64(7-(i+1))))
	}
	if string(rowCode[6]) == "B" {
		//fmt.Println("row:", start)
		return start
	}
	//fmt.Println("row:", end)
	return end
}

func getCol(colCode string) int {
	start := 0
	end := 7
	for i, c := range colCode {
		if string(c) == "R" {
			start = start + int(math.Pow(2, float64(3-(i+1))))
		} else if string(c) == "L" {
			end = end - int(math.Pow(2, float64(3-(i+1))))
		}
	}
	if string(colCode[2]) == "R" {
		//fmt.Println("col:", start)
		return start
	}
	//fmt.Println("col:", end)
	return end
}

func getSeatId(seat string) int {
	rowCode := seat[:7]
	colCode := seat[7:]
	return getRow(rowCode) * 8 + getCol(colCode)
}

func getMax(arr []int) int {
	max := 0
	for _, id := range arr {
		if id > max {
			max = id
		}
	}
	return max
}

func findMySeat(seatIds []int) int {
	nextSeat := 1
	mySeatId := 0
	for i, thisSeatId := range seatIds {
		if thisSeatId != seatIds[nextSeat]-1 {
			mySeatId = thisSeatId + 1
			break
		}
		if i == len(seatIds) - 2 {
			break
		}
		nextSeat++
	}
	return mySeatId
}

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	sc := string(content)
	seats := strings.Split(sc, "\n")

	var seatIds []int
	for _, seat := range seats {
		if seat == "" {
			break
		}
		//fmt.Println(fmt.Sprintf("Seat #%d: [ %s ]", i+1, seat))
		seatIds = append(seatIds, getSeatId(seat))
	}
	sort.Ints(seatIds)
	//fmt.Println(seatIds)
	fmt.Println("My Seat ID:", findMySeat(seatIds))
	fmt.Println("Max Seat ID:", getMax(seatIds))
}
