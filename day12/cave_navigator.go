package main

import (
	"strings"
)

type CaveNavigator struct {
	AllowTwoVisits bool
}

func (this *CaveNavigator) FindAllPaths(caveSystem CaveSystem) int {
	return this.findAllPaths(caveSystem.Start, Path{})

}

func (this *CaveNavigator) findAllPaths(cave *Cave, path Path) int {
	if cave.Name == "end" {
		path.AddNode(cave)
		println(strings.Join(path.nodes, ","))
		return 1
	}
	totalPaths := 0
	path.AddNode(cave)
	connections := this.getConnectionsToVisit(cave, path)
	for _, connection := range connections {
		totalPaths += this.findAllPaths(connection, path)
	}
	return totalPaths
}


func (this *CaveNavigator) getConnectionsToVisit(node *Cave, path Path) []*Cave {
	var eligibleConnections []*Cave
	for _, cave := range node.connections {
		if cave.Name == "start" {
			continue
		}
		if cave.IsSmall()  {
			if !this.isVisited(cave.Name, path) {
				eligibleConnections = append(eligibleConnections, cave)
			}
		} else {
			eligibleConnections = append(eligibleConnections, cave)
		}
	}
	return eligibleConnections
}

func (this *CaveNavigator) isVisited(name string, path Path) bool {
	if !this.AllowTwoVisits {
		for _, node := range path.nodes {
			if node == name {
				return true
			}
		}
	} else {
		for _, node := range path.nodes {
			if node == name && path.visitedTwice {
				return true
			}
		}
	}
	return false
}

type Path struct {
	nodes []string
	visitedTwice bool
}

func (this *Path) AddNode(cave *Cave) {
	if cave.IsSmall() {
		for _, node := range this.nodes {
			if node == cave.Name {
				this.visitedTwice = true
			}
		}
	}
	this.nodes = append(this.nodes, cave.Name)
}