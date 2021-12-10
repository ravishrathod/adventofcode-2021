package main

import (
	"strconv"
	"strings"
)

type BoardEntry struct {
	Value  int
	Marked bool
}

type BingoBoard interface {
	mark(value int) bool
	isSolved() bool
	score() int
}

type BingoBoardStruct struct {
	entries           [][]*BoardEntry
	lastMarkedValue   int
	lastMarkedRows    []int
	lastMarkedColumns []int
	solved            bool
}

func (this *BingoBoardStruct) mark(value int) bool {
	this.lastMarkedRows = []int{}
	this.lastMarkedColumns = []int{}
	for i, row := range this.entries {
		for j, entry := range row {
			if !entry.Marked && entry.Value == value {
				entry.Marked = true
				this.lastMarkedValue = value
				this.lastMarkedColumns = append(this.lastMarkedColumns, j)
				this.lastMarkedRows = append(this.lastMarkedRows, i)
			}
		}
	}
	return this.isSolvedFast()
}

func (this *BingoBoardStruct) isSolvedFast() bool {
	if len(this.lastMarkedRows) == 0 && len(this.lastMarkedColumns) == 0 {
		return false
	}
	solved := true
	for _, rowNum := range this.lastMarkedRows {
		solved = true
		for _, entry := range this.entries[rowNum] {
			if !entry.Marked {
				solved = false
				break
			}
		}
		if solved {
			break
		}
	}
	if solved {
		this.solved = solved
		return solved
	}

	for _, colNum := range this.lastMarkedColumns {
		solved = true
		for _, row := range this.entries {
			if !row[colNum].Marked {
				solved = false
				break
			}
		}
		if solved {
			break
		}
	}
	this.solved = solved
	return solved
}

func (this *BingoBoardStruct) isSolved() bool {
	return this.solved
}

func (this *BingoBoardStruct) score() int {
	unmarkedSum := 0
	for _, row := range this.entries {
		for _, entry := range row {
			if !entry.Marked {
				unmarkedSum += entry.Value
			}
		}
	}
	return unmarkedSum * this.lastMarkedValue
}

func CreateBoard(lines []string) BingoBoard {
	entries := make([][]*BoardEntry, 5)
	for i := range entries {
		entries[i] = make([]*BoardEntry, 5)
	}

	if len(lines) != 5 {
		panic("Invalid rows")
	}
	for i, line := range lines {
		digits := parseDigits(line)
		if len(digits) != 5 {
			panic("invalid input " + line)
		}
		for col, digit := range digits {
			boardEntry := &BoardEntry{
				Value:  digit,
				Marked: false,
			}
			entries[i][col] = boardEntry
		}
	}

	return &BingoBoardStruct{
		entries: entries,
	}
}

func parseDigits(line string) []int {
	var digits []int
	digitStr := strings.Split(line, " ")
	for _, value := range digitStr {
		if strings.TrimSpace(value) == "" {
			continue
		}
		num, _ := strconv.Atoi(value)
		digits = append(digits, num)
	}
	return digits
}
