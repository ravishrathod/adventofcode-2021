package main

import (
	"strconv"
	"strings"
)

func parseLines(fileLines []string) []Line {
	var lines []Line
	for _, fileLine := range fileLines {
		startAndEnd := strings.Split(fileLine, " -> ")
		startPoint := parsePoint(startAndEnd[0])
		endPoint := parsePoint(startAndEnd[1])
		line := CreateLine(startPoint, endPoint)
		lines = append(lines, *line)
	}

	return lines
}

func parsePoint(startAndEnd string) Point {
	xy := strings.Split(startAndEnd, ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])
	return Point{
		X: x,
		Y: y,
	}
}
