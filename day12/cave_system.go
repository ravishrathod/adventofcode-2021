package main

import "strings"

type CaveSystem struct {
	Start *Cave
	End *Cave
	Caves map[string]*Cave
}

func CreateCaveSystem(lines []string) *CaveSystem {
	caves := make(map[string]*Cave)

	for _, line := range lines {
		entries := strings.Split(line, "-")
		addCaves(entries, caves)

	}
	return &CaveSystem{
		Start: caves["start"],
		End: caves["end"],
		Caves: caves,
	}
}

func addCaves(entries []string, caves map[string]*Cave) {
	var firstCave, secondCave *Cave
	var found bool
	if firstCave, found = caves[entries[0]]; !found {
		firstCave = createCave(entries[0])
		caves[entries[0]] = firstCave
	}
	if secondCave, found = caves[entries[1]]; !found {
		secondCave = createCave(entries[1])
		caves[entries[1]] = secondCave
	}
	firstCave.AddConnection(secondCave)
	secondCave.AddConnection(firstCave)
}

type Cave struct {
	Name string
	connections map[string]*Cave
}

func createCave(name string) *Cave {
	return &Cave{
		Name: name,
		connections: map[string]*Cave{},
	}
}

func (this *Cave) AddConnection(connection *Cave) {
	if _, found := this.connections[connection.Name]; ! found  {
		this.connections[connection.Name] = connection
	}
}

func (this *Cave) IsSmall() bool {
	return (this.Name != "start" && this.Name != "end") && this.Name == strings.ToLower(this.Name)
}