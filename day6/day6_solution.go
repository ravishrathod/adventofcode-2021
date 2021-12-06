package main

import (
	"adventoccode2021/commons"
	"strconv"
	"strings"
)

func main() {
	lines, err := commons.ReadFile("input/day6.txt")
	if err != nil {
		panic(err)
	}
	dayOneCounters := toInt(lines[0])
	fishesByCounter := make(map[int]int)

	for _, counter := range dayOneCounters {
		fishesByCounter[counter]++
	}
	for i := 1;i<=256;i++ {
		existingDayZeroFishes := fishesByCounter[0]
		for day := 1;day <=8;day++ {
			fishesByCounter[day-1] = fishesByCounter[day]
		}
		fishesByCounter[8] = existingDayZeroFishes
		fishesByCounter[6] += existingDayZeroFishes
	}
	totalFishes := 0
	for _, fishes := range fishesByCounter {
		totalFishes = totalFishes + fishes
	}
	println("Total fishes: ", totalFishes)
}

func toInt(line string) []int {
	values := strings.Split(line, ",")
	var days []int
	for _, val := range values {
		day, _ := strconv.Atoi(val)
		days = append(days, day)
	}
	return days
}
