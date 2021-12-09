package main

import (
	"sort"
	"strings"
)

type SignalPattern struct {
	Segments string
	SortedSegments string
}

func (this *SignalPattern) SegmentArray() []string {
	var array []string
	for _, val := range []rune(this.Segments) {
		array = append(array, string(val))
	}
	return array
}

func (this *SignalPattern) ContainsAll(input string) bool  {
	containsAll := true
	for _, val := range []rune(input) {
		containsAll = containsAll && strings.Contains(this.SortedSegments, string(val))
		if !containsAll {
			break
		}
	}
	return containsAll
}

func (this *SignalPattern) PossibleValues() []int {
	var possibleValues []int
	if len(this.Segments) == 2 {
		possibleValues = append(possibleValues, 1)
		return possibleValues
	}
	if len(this.Segments) == 7 {
		possibleValues = append(possibleValues, 8)
		return possibleValues
	}
	if len(this.Segments) == 4 {
		possibleValues = append(possibleValues, 4)
		return possibleValues
	}
	if len(this.Segments) == 3 {
		possibleValues = append(possibleValues, 7)
		return possibleValues
	}
	return possibleValues
}

type Reading struct {
	Patterns []*SignalPattern
	Output   []*SignalPattern
}

func CreateSignalPattern(input string) *SignalPattern {
	if len(input) > 7 {
		panic("Invalid input" + input)
	}
	var segments []string
	for _, val := range []rune(input) {
		segments = append(segments, string(val))
	}
	sort.Strings(segments)
	return &SignalPattern{
		Segments: input,
		SortedSegments: strings.Join(segments, ""),
	}
}

func parseSignals(lines []string) []Reading {
	var readings []Reading
	for _, line := range lines {
		var patterns []*SignalPattern
		var outputs []*SignalPattern

		patternsAndReadings := strings.Split(line, "|")
		patternStrings := strings.Split(strings.TrimSpace(patternsAndReadings[0]), " ")
		for _, patternString := range patternStrings {
			patterns = append(patterns, CreateSignalPattern(patternString))
		}

		outputStrings := strings.Split(strings.TrimSpace(patternsAndReadings[1]), " ")
		for _, outputString := range outputStrings {
			outputs = append(outputs, CreateSignalPattern(outputString))
		}
		readings = append(readings, Reading{
			Patterns: patterns,
			Output:   outputs,
		})
	}
	return readings
}