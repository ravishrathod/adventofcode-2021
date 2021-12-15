package main

import "math"
//PathFinder : Dijkstra’s Shortest Path
type PathFinder struct {
	matrix            [][]int
	width             int
	height            int
	distanceFromStart map[Point]int
	visitedNodes      map[Point]bool
}

func NewPathFinder(matrix [][]int) *PathFinder {
	distanceFromStart := make(map[Point]int)

	height := len(matrix)
	width := len(matrix[0])
	for y:=0;y<height;y++ {
		for x:=0;x<width;x++ {
			point := Point{
				X: x,
				Y: y,
			}
			distanceFromStart[point] = math.MaxInt
		}
	}
	point := Point{
		X: 0,
		Y: 0,
	}
	distanceFromStart[point] = 0

	pathFinder := &PathFinder{
		matrix: matrix,
		height: height,
		width:  width,
		distanceFromStart: distanceFromStart,
		visitedNodes: make(map[Point]bool),
	}
	return pathFinder
}

func (pathFinder *PathFinder) FindShortestPath() {
	pathFinder.findDistanceForNode(0, 0)
}

func (pathFinder *PathFinder) findDistanceForNode(x int, y int) {
	point := Point{
		X:x,
		Y:y,
	}
	pathFinder.visitedNodes[point] = true
	//if x == 0 && y == 0 {
	//	pathFinder.distanceFromStart[point] = 0
	//}
	neighbours := pathFinder.unvisitedNeighbours(x, y)
	currentDistance := pathFinder.distanceFromStart[point]
	for _, neighbour := range neighbours {
		if currentDistance + pathFinder.matrix[neighbour.Y][neighbour.X] < pathFinder.distanceFromStart[neighbour] {
			pathFinder.distanceFromStart[neighbour] = currentDistance + pathFinder.matrix[neighbour.Y][neighbour.X]
		}
	}
	node, found := pathFinder.unvisitedClosestNode()
	if found {
		pathFinder.findDistanceForNode(node.X, node.Y)
	}

}

func (pathFinder *PathFinder) unvisitedClosestNode() (Point, bool) {
	smallestDistance := math.MaxInt
	var point Point
	hasResult := false
	for node, distance := range pathFinder.distanceFromStart {
		if _, found := pathFinder.visitedNodes[node]; !found {
			if distance < smallestDistance {
				smallestDistance = distance
				point = node
				hasResult = true
			}
		}
	}
	return point, hasResult
}

func (pathFinder *PathFinder) unvisitedNeighbours(x int, y int) []Point {
	var points []Point
	if x > 0 {
		n := Point{
			X: x-1,
			Y: y,
		}
		if _, found := pathFinder.visitedNodes[n]; !found {
			points = append(points, n)
		}
	}
	if x < pathFinder.width-1 {
		n := Point{
			X: x+1,
			Y: y,
		}
		if _, found := pathFinder.visitedNodes[n]; !found {
			points = append(points, n)
		}
	}
	if y > 0 {
		n := Point{
			X: x,
			Y: y-1,
		}
		if _, found := pathFinder.visitedNodes[n]; !found {
			points = append(points, n)
		}
	}
	if y < pathFinder.height-1 {
		n := Point{
			X: x,
			Y: y+1,
		}
		if _, found := pathFinder.visitedNodes[n]; !found {
			points = append(points, n)
		}
	}

	return points
}

func (pathFinder *PathFinder) GetShortestDistance() int {
	point := Point{
		X: pathFinder.width-1,
		Y: pathFinder.height-1,
	}
	return pathFinder.distanceFromStart[point]
}

type Point struct {
	X int
	Y int
}
