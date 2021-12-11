package main

import (
	"adventoccode2021/commons"
	"sort"
)

func main() {
	lines, err := commons.ReadFile("input/day9.txt")
	if err != nil {
		panic(err)
	}
	heightMap := parseHeightMap(lines)
	part1(heightMap)
	part2(heightMap)
}

func part1(heightMap [][]int) {
	lowPoints := findLowPoints(heightMap)
	risk := 0
	for _, lowPoint := range lowPoints {
		risk += lowPoint.Value + 1
	}
	println("Risk ", risk)
}

func part2(heightMap [][]int) {
	lowPoints := findLowPoints(heightMap)
	var basins []Basin

	for _, lowPoint := range lowPoints {
		basin := findBasin(lowPoint, heightMap)
		basins = append(basins, basin)
	}

	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i].Points) > len(basins[j].Points)
	})

	product := 1
	for i, basin := range basins {
		if i > 2 {
			break
		}
		product *= len(basin.Points)
	}
	println("Product ", product)
}

func findBasin(lowPoint LowPoint, heightMap [][]int) Basin {
	basin := &Basin{
		LowPoint: lowPoint,
		Points:   make(map[string]bool),
	}
	addPointsToBasin(basin, lowPoint.X, lowPoint.Y, lowPoint.Value-1, heightMap)

	return *basin
}

func addPointsToBasin(basin *Basin, x int, y int, sourceValue int, heightMap [][]int) {
	height := len(heightMap)
	width := len(heightMap[0])
	currentValue := heightMap[y][x]
	if currentValue < sourceValue || currentValue == 9 {
		return
	}
	if !basin.HasPoint(x, y) {
		basin.AddPoint(x, y)
	}
	if x > 0 && !basin.HasPoint(x-1, y) {
		addPointsToBasin(basin, x-1, y, currentValue, heightMap)
	}
	if x < width-1 && !basin.HasPoint(x+1, y) {
		addPointsToBasin(basin, x+1, y, currentValue, heightMap)
	}
	if y > 0 && !basin.HasPoint(x, y-1) {
		addPointsToBasin(basin, x, y-1, currentValue, heightMap)
	}
	if y < height-1 && !basin.HasPoint(x, y+1) {
		addPointsToBasin(basin, x, y+1, currentValue, heightMap)
	}

}

func findLowPoints(heightMap [][]int) []LowPoint {
	var lowPoints []LowPoint
	height := len(heightMap)
	width := len(heightMap[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if isLowest(x, y, width, height, heightMap) {
				lowPoint := LowPoint{
					Value: heightMap[y][x],
					X:     x,
					Y:     y,
				}
				lowPoints = append(lowPoints, lowPoint)
			}
		}
	}
	return lowPoints
}

func isLowest(x int, y int, width int, height int, heightMap [][]int) bool {
	depth := heightMap[y][x]
	lowPoint := true
	if x > 0 {
		lowPoint = lowPoint && depth < heightMap[y][x-1]
	}
	if x < width-1 {
		lowPoint = lowPoint && depth < heightMap[y][x+1]
	}
	if y > 0 {
		lowPoint = lowPoint && depth < heightMap[y-1][x]
	}
	if y < height-1 {
		lowPoint = lowPoint && depth < heightMap[y+1][x]
	}
	return lowPoint
}

func parseHeightMap(lines []string) [][]int {
	heightMap := make([][]int, len(lines))
	for i, line := range lines {
		numbers := commons.LineToIntArrayNoSeparator(line)
		heightMap[i] = numbers
	}
	return heightMap
}
