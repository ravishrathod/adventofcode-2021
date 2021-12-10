package main

import (
	"adventoccode2021/commons"
	"strconv"
	"strings"
)

func main() {
	lines, err := commons.ReadFile("input/day8.txt")
	if err != nil {
		panic(err)
	}
	readings := parseSignals(lines)
	part1(readings)
	part2(readings)
}

func part2(readings []Reading) {
	sum := 0
	for _, reading := range readings {
		value := decode(reading)
		println(value)
		sum += value
	}
	println(sum)
}

func decode(reading Reading) int {
	numberToPattern := make(map[int]SignalPattern)
	segmentsToNumber := make(map[string]int)

	for _, signalPattern := range reading.Patterns {
		possibleValues := signalPattern.PossibleValues()
		if len(possibleValues) == 1 {
			numberToPattern[possibleValues[0]] = *signalPattern
			segmentsToNumber[signalPattern.SortedSegments] = possibleValues[0]
		}
	}
	onePattern := numberToPattern[1]
	fourPatterns := numberToPattern[4]
	lPattern := *computeL(fourPatterns, onePattern)
	segmentsToNumber[lPattern.SortedSegments] = 11
	numberToPattern[11] = lPattern

	finalReading := ""
	for _, outputPatterns := range reading.Output {
		if value, found := segmentsToNumber[outputPatterns.SortedSegments]; found {
			finalReading = finalReading + strconv.Itoa(value)
		} else if len(outputPatterns.SortedSegments) == 5 {
			if outputPatterns.ContainsAll(numberToPattern[1].SortedSegments) {
				finalReading = finalReading + "3"
			} else if outputPatterns.ContainsAll(numberToPattern[11].SortedSegments) {
				finalReading = finalReading + "5"
			} else {
				finalReading = finalReading + "2"
			}
		} else {
			if outputPatterns.ContainsAll(numberToPattern[4].SortedSegments) {
				finalReading = finalReading + "9"
			} else if outputPatterns.ContainsAll(numberToPattern[11].SortedSegments) {
				finalReading = finalReading + "6"
			} else {
				finalReading = finalReading + "0"
			}
		}
	}
	intValue, _ := strconv.Atoi(finalReading)
	return intValue
}

func computeL(four SignalPattern, one SignalPattern) *SignalPattern {
	segments := one.Segments
	result := ""
	for _, char := range four.SegmentArray() {
		if !strings.Contains(segments, char) {
			result = result + char
		}
	}
	return CreateSignalPattern(result)
}

func part1(readings []Reading) {
	count := 0

	for _, readings := range readings {
		for _, output := range readings.Output {
			if len(output.PossibleValues()) == 1 {
				count++
			}
		}
	}

	println(count)
}
