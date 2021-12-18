package main

import (
	"adventoccode2021/commons"
	"fmt"
	_ "net/http/pprof"
	"time"
)

func main() {
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()
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
	start := time.Now()
	pathFinder.FindShortestPath()
	println(pathFinder.GetShortestDistance())
	fmt.Printf("Time : %d(ms)", time.Since(start)/time.Millisecond)
	println("")

	extrapolater := NewCaveExtrapolater(matrix)
	biggerCave := extrapolater.Extrapolate()

	start = time.Now()
	pathFinder = NewPathFinder(biggerCave)
	pathFinder.FindShortestPath()
	println(pathFinder.GetShortestDistance())
	fmt.Printf("Time : %d(ms)", time.Since(start)/time.Millisecond)
	println("")
}