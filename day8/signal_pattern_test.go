package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSignalPattern(t *testing.T) {
	lines := []string{"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe", "edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc"}
	readings := parseSignals(lines)
	assert.Equal(t, 2, len(readings))

	reading := readings[0]
	assert.Equal(t, 10, len(reading.Patterns))
	assert.Equal(t, 4, len(reading.Output))
}

func Test_SignalPattern_PossibleValues(t *testing.T) {
	assertPossibleValue(t, "edb", 7)
	assertPossibleValue(t, "fdgacbe", 8)
	assertPossibleValue(t, "cg", 1)
	assertPossibleValue(t, "fcge", 4)
}

func assertPossibleValue(t *testing.T, pattern string, possibleValue int) {
	signalPattern := CreateSignalPattern(pattern)
	possibleValues := signalPattern.PossibleValues()
	assert.Equal(t, 1, len(possibleValues))
	assert.Equal(t, possibleValue, possibleValues[0])
}
