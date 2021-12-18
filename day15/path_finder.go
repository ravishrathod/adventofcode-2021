package main

import (
	"math"
)

//PathFinder : Dijkstra’s Shortest Path
type PathFinder struct {
	matrix         [][]int
	width          int
	height         int
	nodes          map[Point]int
	visitedNodes   map[Point]bool
	unVisitedQueue *VertexQueue
}

func NewPathFinder(matrix [][]int) *PathFinder {
	distanceFromStart := make(map[Point]int)
	height := len(matrix)
	width := len(matrix[0])
	var points []Point

	zeroPoint := Point{
		X:        0,
		Y:        0,
	}
	points = append(points, zeroPoint)
	distanceFromStart[zeroPoint] = 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			point := Point{
				X:        x,
				Y:        y,
			}
			if !(x == 0 && y == 0) {
				points = append(points, point)
				distanceFromStart[point] = math.MaxInt
			}
		}
	}
	vertexQueue := NewVertexQueue(points)

	pathFinder := &PathFinder{
		matrix:         matrix,
		height:         height,
		width:          width,
		nodes:          distanceFromStart,
		visitedNodes:   make(map[Point]bool),
		unVisitedQueue: vertexQueue,
	}
	return pathFinder
}

func (pathFinder *PathFinder) FindShortestPath() {
	node, _ := pathFinder.unvisitedClosestNode()
	pathFinder.findDistanceForNode(node)
}

func (pathFinder *PathFinder) findDistanceForNode(point Point) {
	pathFinder.visitedNodes[point] = true
	neighbours := pathFinder.unvisitedNeighbours(point.X, point.Y)
	currentDistance := pathFinder.getDistanceFromStart(point)
	for _, neighbour := range neighbours {
		if currentDistance+pathFinder.matrix[neighbour.Y][neighbour.X] < pathFinder.getDistanceFromStart(neighbour) {
			updatedDistance := currentDistance + pathFinder.matrix[neighbour.Y][neighbour.X]
			pathFinder.updateDistance(neighbour, updatedDistance)
		}
	}
	node, found := pathFinder.unvisitedClosestNode()
	if found {
		pathFinder.findDistanceForNode(node)
	}
}

func (pathFinder *PathFinder) getDistanceFromStart(point Point) int {
	if distance, found := pathFinder.nodes[point]; found {
		return distance
	}
	return math.MaxInt
}

func (pathFinder *PathFinder) updateDistance(point Point, distance int) {
	pathFinder.nodes[point] = distance
	pathFinder.unVisitedQueue.Update(point, distance)
}

func (pathFinder *PathFinder) unvisitedClosestNode() (Point, bool) {
	point := pathFinder.unVisitedQueue.Pop()
	if point == nil {
		return Point{}, false
	}
	return *point, true
}

func (pathFinder *PathFinder) unvisitedNeighbours(x int, y int) []Point {
	var points []Point
	if x > 0 {
		n := Point{
			X: x - 1,
			Y: y,
		}
		if _, found := pathFinder.visitedNodes[n]; !found {
			points = append(points, n)
		}
	}
	if x < pathFinder.width-1 {
		n := Point{
			X: x + 1,
			Y: y,
		}
		if _, found := pathFinder.visitedNodes[n]; !found {
			points = append(points, n)
		}
	}
	if y > 0 {
		n := Point{
			X: x,
			Y: y - 1,
		}
		if _, found := pathFinder.visitedNodes[n]; !found {
			points = append(points, n)
		}
	}
	if y < pathFinder.height-1 {
		n := Point{
			X: x,
			Y: y + 1,
		}
		if _, found := pathFinder.visitedNodes[n]; !found {
			points = append(points, n)
		}
	}

	return points
}

func (pathFinder *PathFinder) GetShortestDistance() int {
	point := Point{
		X: pathFinder.width - 1,
		Y: pathFinder.height - 1,
	}
	return pathFinder.nodes[point]
}

type Point struct {
	X        int
	Y        int
}
