package main

import (
	"adventoccode2021/commons"
)

func main() {
	lines, err := commons.ReadFile("input/day6.txt")
	if err != nil {
		panic(err)
	}
	dayOneCounters := commons.LinetoInt(lines[0])
	fishesByCounter := make(map[int]int)

	for _, counter := range dayOneCounters {
		fishesByCounter[counter]++
	}
	for day := 1; day <= 256; day++ {
		existingDayZeroFishes := fishesByCounter[0]
		for counter := 1; counter <= 8; counter++ {
			fishesByCounter[counter-1] = fishesByCounter[counter]
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
