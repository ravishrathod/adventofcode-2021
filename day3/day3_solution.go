package main

import (
	"adventoccode2021/commons"
	"strconv"
)

func main() {
	lines, err := commons.ReadFile("input/day3.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var oneCounts []int = make([]int, 12)
	var zeroCounts []int = make([]int, 12)
	var gamma = ""
	var epsilon = ""
	for _, input := range lines {
		digits := []rune(input)
		if len(digits) != 12 {
			panic("unexpected input " + input)
		}
		for i, digit := range digits {
			if string(digit) == "0" {
				zeroCounts[i] = zeroCounts[i] + 1
			} else {
				oneCounts[i] = oneCounts[i] + 1
			}
		}
	}

	for i, zeroCount := range zeroCounts {
		oneCount := oneCounts[i]
		if zeroCount > oneCount {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		}
	}
	gammaInt, _ := strconv.ParseInt(gamma, 2, 0)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 0)

	println(gammaInt * epsilonInt)
}

func part2(lines []string) {
	const inputLength = 12
	var oxygenValues = lines
	var co2Values = lines

	for position := 0 ;position<inputLength;position++ {
		if len(oxygenValues) != 1 {
			oxygenCounts := calculateCountsForPosition(oxygenValues, position)
			oxygenValues = calculateOxygenRating(oxygenValues, oxygenCounts, position)
		}
		if len(co2Values) != 1 {
			co2Counts := calculateCountsForPosition(co2Values, position)
			co2Values = calculateCo2Rating(co2Values, co2Counts, position)
		}
		if len(oxygenValues) == 1 && len(co2Values) == 1 {
			break
		}
	}
	oxygenRating, _ := strconv.ParseInt(oxygenValues[0], 2, 0)
	co2Rating, _ := strconv.ParseInt(co2Values[0], 2, 0)

	println(oxygenRating * co2Rating)
}

func calculateOxygenRating(values []string, counts digitCounts, position int) []string {
	mostCommonChar := "0"
	if counts.OneCount >= counts.ZeroCount {
		mostCommonChar = "1"
	}
	return filterValues(values, position, mostCommonChar)
}

func calculateCo2Rating(values []string, counts digitCounts, position int) []string {
	leastCommonChar := "1"
	if counts.ZeroCount <= counts.OneCount {
		leastCommonChar = "0"
	}
	return filterValues(values, position, leastCommonChar)
}

func filterValues(values []string, position int, selectionChar string) []string {
	var filteredValues []string
	for _, value := range values {
		charAtPosition := string(value[position])
		if charAtPosition == selectionChar {
			filteredValues = append(filteredValues, value)
		}
	}
	return filteredValues
}

func calculateCountsForPosition(lines []string, position int) digitCounts  {
	count := &digitCounts{}
	for _, value := range lines {
		charAtPosition := string(value[position])
		if charAtPosition == "0" {
			count.ZeroCount = count.ZeroCount + 1
		} else {
			count.OneCount = count.OneCount + 1
		}
	}
	return *count
}

type digitCounts struct {
	ZeroCount int
	OneCount int
}
