package main

import "strconv"
import "adventoccode2021/commons"

func main() {
	lines, err := commons.ReadFile("input/day1.txt")
	if err != nil {
		panic(err)
	}
	input := parserToInt(lines)

	part1(input)
	part2(input)

}

func part1(input []int) {
	depthIncreaseCount := calculateTotalIncrements(input)
	println(depthIncreaseCount)
}

func calculateTotalIncrements(input []int) int {
	var depthIncreaseCount = 0
	for i, depth := range input {
		if i > 0 {
			if depth > input[i-1] {
				depthIncreaseCount = depthIncreaseCount + 1
			}
		}
	}
	return depthIncreaseCount
}

func part2(input []int) {
	var windows []int
	for i, depth := range input {
		if i >1 {
			windows = append(windows, depth + input[i-1] + input[i-2])
		}
	}
	depthIncreaseCount := calculateTotalIncrements(windows)
	println(depthIncreaseCount)
}

func parserToInt(lines []string) []int {
	var numbers []int
	for _, input := range lines {
		parsedValue, _ := strconv.Atoi(input)
		numbers = append(numbers, parsedValue)
	}
	return numbers
}