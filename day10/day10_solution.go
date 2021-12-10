package main

import (
	"adventoccode2021/commons"
	"github.com/montanaflynn/stats"
)

var errorScoreForSymbol = map[string] int {
	")" : 3,
	"]" : 57,
	"}" : 1197,
	">" : 25137,
}

var completionScoreForSymbol = map[string] int {
	")" : 1,
	"]" : 2,
	"}" : 3,
	">" : 4,
}

func main() {
	lines, err:= commons.ReadFile("input/day10.txt")
	if err != nil {
		panic(err)
	}
	process(lines)
}

func process(lines []string)  {
	errorScore := 0
	var completionScores  []float64
	for _, line := range lines {
		parser := CreateSubSystemParser(line)
		parser.Parse()
		if parser.IsCorrupted() {
			errorScore += errorScoreForSymbol[parser.IllegalSymbol()]
		} else if parser.IsIncomplete() {
			completionScore := 0
			for _, symbol := range parser.CompletionSymbols() {
				completionScore *= 5
				completionScore += completionScoreForSymbol[symbol]
			}
			completionScores = append(completionScores, float64(completionScore))
		}
	}
	median, _ := stats.Median(completionScores)
	println("Total error score ", errorScore)
	println("Median completion score ", int(median))
}


