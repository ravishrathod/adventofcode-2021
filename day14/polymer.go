package main

import "math"

type Polymer struct {
	sequence string
	pairCount map[string]int
}

func CreatePolymer(sequence string) *Polymer {
	polymer := &Polymer{
		sequence:  sequence,
		pairCount: make(map[string]int),
	}
	var array []string
	for _, char := range []rune(polymer.sequence) {
		array = append(array, string(char))
	}

	length := len(array)
	for i:=0;i<length-1;i++ {
		pair := array[i] + array[i+1]
		polymer.pairCount[pair] =  polymer.pairCount[pair] + 1
	}
	return polymer
}

func (polymer *Polymer) ApplyRules(rules map[string]string) {
	updatedCounts := make(map[string]int)
	for pair := range polymer.pairCount {
		if insertChar, found := rules[pair];found {
			firstNewPair := pair[:len(pair)-1] + insertChar
			//polymer.pairCount[pair] new pairs will be generated
			updatedCounts[firstNewPair] = updatedCounts[firstNewPair] + polymer.pairCount[pair]

			secondNewPair := insertChar + pair[len(pair)-1:]
			updatedCounts[secondNewPair] = updatedCounts[secondNewPair] + polymer.pairCount[pair]
		}
	}
	polymer.pairCount = updatedCounts
}

func (polymer *Polymer) GetMaxAndMinCounts() (int, int) {
	charCounts := make(map[string]int)
	for pair, count := range polymer.pairCount {
		firstChar := pair[:len(pair)-1]
		charCounts[firstChar] = charCounts[firstChar] + count
	}
	templateRune := []rune(polymer.sequence)
	lastTemplateChar := string(templateRune[len(templateRune)-1])
	charCounts[lastTemplateChar] = charCounts[lastTemplateChar] + 1

	maxCount := -1
	minCount := math.MaxInt
	for _, count := range charCounts {
		if count < minCount {
			minCount = count
		}
		if count > maxCount {
			maxCount =  count
		}
	}
	return maxCount, minCount
}