package main

import "adventoccode2021/commons"

func main() {
	lines, err := commons.ReadFile("input/day12.txt")
	if err != nil {
		panic(err)
	}
	caveSystems := CreateCaveSystem(lines)
	part1(caveSystems)
	part2(caveSystems)
}

func part1(caveSystems *CaveSystem) {
	navigator := CaveNavigator{}
	paths := navigator.FindAllPaths(*caveSystems)
	println(paths)
}

func part2(caveSystems *CaveSystem) {
	navigator := CaveNavigator{
		AllowTwoVisits: true,
	}
	paths := navigator.FindAllPaths(*caveSystems)
	println(paths)
}