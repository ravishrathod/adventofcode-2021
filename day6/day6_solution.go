package main

import (
	"adventoccode2021/commons"
	"strconv"
	"strings"
	"time"
)

func main() {
	const partitionSize = 8000
	lines, err := commons.ReadFile("input/day6_sample.txt")
	if err != nil {
		panic(err)
	}
	dayOneCounters := toInt(lines[0])
	fishes := counterToFishes(dayOneCounters)
	startTime := time.Now()
	for i := 1;i<=256;i++ {
		var allNewFishes []*LanternFish
		partition := partitionList(fishes, partitionSize)
		listenerChannel := make(chan []*LanternFish, len(partition))
		for _, subList := range partition {
			go processSubList(subList, listenerChannel)
		}
		for i:=0;i<len(partition);i++ {
			select {
				case newFishes := <- listenerChannel:
					allNewFishes = append(allNewFishes, newFishes...)
			}
		}
		fishes = append(fishes, allNewFishes...)
		//var newFishes []*LanternFish
		//for _, fish := range fishes {
		//	newFish := fish.DayPassed()
		//	if newFish != nil {
		//		newFishes = append(newFishes, newFish)
		//	}
		//}
		//fishes = append(fishes, newFishes...)
	}
	println("Total time taken: ", time.Since(startTime) / time.Millisecond)
	println("Total fishes: ", len(fishes))
}

func processSubList(fishes []*LanternFish, listenerChannel chan []*LanternFish) {
	var newFishes []*LanternFish
	for _, fish := range fishes {
		newFish := fish.DayPassed()
		if newFish != nil {
			newFishes = append(newFishes, newFish)
		}
	}
	listenerChannel <- newFishes
}

func counterToFishes(counters []int) []*LanternFish {
	var fishes []*LanternFish
	for _, counter := range counters {
		fish := CreateFish(counter)
		fishes = append(fishes, fish)
	}
	return fishes
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
