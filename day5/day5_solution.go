package main

import (
	"adventoccode2021/commons"
)

func main() {
	fileLines, err := commons.ReadFile("input/day5.txt")
	if err != nil {
		panic(err)
	}
	part1(fileLines)
	part2(fileLines)

}

func part1(fileLines []string) {
	println("----Part1----")
	lines := parseLines(fileLines)
	var strightLines []Line
	for _, line := range lines {
		if line.IsHorizontal() || line.IsVertical() {
			strightLines = append(strightLines, line)
		}
	}
	calculateIntersectionsV2(strightLines)
}

func part2(fileLines []string) {
	println("----Part2----")
	lines := parseLines(fileLines)
	calculateIntersectionsV2(lines)
}

func calculateIntersectionsV2(lines []Line) {
	intersectionCounts := make(map[Point]int)
	pointsWithMoreThanOneIntersections := 0
	for _, line := range lines {
		points := line.getPoints()
		for _, point := range points {
			intersectionCounts[point] ++
		}
	}
	for _, count := range intersectionCounts {
		if count > 1 {
			pointsWithMoreThanOneIntersections++
		}
	}

	println(pointsWithMoreThanOneIntersections)
}