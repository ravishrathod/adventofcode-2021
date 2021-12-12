package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCave_CreateCaveSystem(t *testing.T) {
	lines := []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}

	caveSystem := CreateCaveSystem(lines)

	assert.Equal(t, 6, len(caveSystem.Caves))
	assert.Equal(t, 2, len(caveSystem.Start.connections))
	caveBConnections := caveSystem.Caves["b"].connections
	assert.Equal(t, 4, len(caveBConnections))
	assert.Contains(t, caveBConnections, "start")
	assert.Contains(t, caveBConnections, "A")
	assert.Contains(t, caveBConnections, "d")
	assert.Contains(t, caveBConnections, "end")
}

