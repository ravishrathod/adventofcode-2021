package main

import (
	"strconv"
	"strings"
)

type FoldInstruction struct {
	Axis string
	At int
}

func ParseFoldInstructions(lines []string) []FoldInstruction {
	var foldInstructions []FoldInstruction
	for _, line := range lines {
		if strings.HasPrefix(line, "fold along ") {
			trimmed := strings.ReplaceAll(line, "fold along ", "")
			parts := strings.Split(trimmed, "=")
			at,_ := strconv.Atoi(parts[1])
			foldInstruction := FoldInstruction{
				Axis: parts[0],
				At: at,
			}
			foldInstructions = append(foldInstructions, foldInstruction)
		}
	}
	return foldInstructions
}
