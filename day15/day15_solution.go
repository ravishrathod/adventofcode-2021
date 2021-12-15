package main

import (
	"adventoccode2021/commons"
)

func main() {
	lines, err := commons.ReadFile("input/day15.txt")
	if err != nil {
		panic(err)
	}

	matrix := make([][]int, 0)
	for _, line := range lines {
		row := commons.LineToIntArrayNoSeparator(line)
		matrix = append(matrix, row)

	}
	pathFinder := NewPathFinder(matrix)
	pathFinder.FindShortestPath()
	print(pathFinder.GetShortestDistance())
}