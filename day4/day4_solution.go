package main

import (
	"adventoccode2021/commons"
	"fmt"
	"strings"
)

func main() {
	lines, err := commons.ReadFile("input/day4.txt")
	if err != nil {
		panic(err)
	}
	draws := parseDraws(lines)
	part1(draws, parseBoards(lines))
	part2(draws, parseBoards(lines))

}

func part1(draws []int, bingoBoards []BingoBoard) {
	solved := false
	for _, draw := range draws {
		if solved {
			break
		}
		for i, board := range bingoBoards {
			solved = board.mark(draw)
			if solved {
				fmt.Printf("Draw: %v . Board number: %v", draw, i)
				println("")
				println("Score: ", board.score())
				break
			}
		}
	}
}

func part2(draws []int, bingoBoards []BingoBoard)  {
	totalBoardsSolved := 0
	complete := false
	for _, draw := range draws {
		if complete {
			break
		}
		for _, board := range bingoBoards {
			if board.isSolved() {
				continue
			}
			solved := board.mark(draw)
			if solved {
				totalBoardsSolved += 1
				if totalBoardsSolved == len(bingoBoards)  {
					complete = true
					println(board.score())
					break
				}
			}

		}
	}
}

func parseBoards(lines []string) []BingoBoard {
	var boardLines []string
	var bingoBoards []BingoBoard
	for i, line := range lines {
		if i <= 1 {
			continue
		}
		if strings.TrimSpace(line) == "" {
			bingoBoard := CreateBoard(boardLines)
			bingoBoards = append(bingoBoards, bingoBoard)
			boardLines = []string{}
			continue
		}
		boardLines = append(boardLines, line)
	}
	return bingoBoards
}

func parseDraws(lines []string) []int {
	var draws []int
	firstLine := lines[0]
	draws = commons.LinetoInt(firstLine)
	return draws
}