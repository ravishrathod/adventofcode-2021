package main

import "adventoccode2021/commons"

func main() {
	lines, err := commons.ReadFile("input/day11.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	grid := parseGrid(lines)
	octopusGrid := OctopusGrid{
		grid: grid,
	}
	octopusGrid.Tick(100)
	println(octopusGrid.GetFlashes())
}

func part2(lines []string) {
	grid := parseGrid(lines)
	octopusGrid := OctopusGrid{
		grid: grid,
	}
	println(octopusGrid.FirstSynchronousFlash())
}

func parseGrid(lines []string) [][]int {
	heightMap := make([][]int, len(lines))
	for i, line := range lines {
		numbers := commons.LineToIntArrayNoSeparator(line)
		heightMap[i] = numbers
	}
	return heightMap
}
