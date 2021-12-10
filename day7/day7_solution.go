package main

import (
	"adventoccode2021/commons"
	"math"
)
import "github.com/montanaflynn/stats"

func main() {
	lines, err := commons.ReadFile("input/day7.txt")
	if err != nil {
		panic(err)
	}
	positions := commons.LinetoInt(lines[0])
	part1(positions)
	part2(positions)
}

func part1(positions []int) {
	median := findMedian(positions)
	coast := calculateLinearCostToMove(positions, median)
	println(coast)
}

func part2(positions []int) {
	mean := findMean(positions)
	targetsToTry := []int{mean - 1, mean, mean + 1}
	var costs []float64
	for _, target := range targetsToTry {
		cost := calculateExponentialCostToMove(positions, target)
		costs = append(costs, float64(cost))
	}
	minCost, _ := stats.Min(costs)
	print(int(minCost))
}

func calculateExponentialCostToMove(positions []int, targetPosition int) int {
	totalCost := 0
	for _, pos := range positions {
		totalCost += calculateExponentialCost(pos, targetPosition)
	}
	return totalCost
}

func calculateExponentialCost(from int, to int) int {
	distance := mod(from - to)
	coast := (distance) * (distance + 1) / 2
	return coast
}

func calculateLinearCostToMove(positions []int, targetPosition int) int {
	totalCost := 0
	for _, pos := range positions {
		cost := mod(pos - targetPosition)
		totalCost += cost

	}
	return totalCost
}

func mod(value int) int {
	if value < 0 {
		value = value * -1
	}
	return value
}

func findMedian(positions []int) int {
	var floatValues []float64
	for _, pos := range positions {
		fValue := float64(pos)
		floatValues = append(floatValues, fValue)
	}
	median, _ := stats.Median(floatValues)
	return int(median)
}

func findMean(positions []int) int {
	var floatValues []float64
	for _, pos := range positions {
		fValue := float64(pos)
		floatValues = append(floatValues, fValue)
	}
	mean, _ := stats.Mean(floatValues)
	mean = math.Round(mean)
	return int(mean)
}
