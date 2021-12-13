package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateInstructionMatrix(t *testing.T) {
	lines := []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
		"",
		"fold along y=7",
		"fold along x=5",
	}
	instructionMatrix := CreateInstructionMatrix(lines)
	assert.Equal(t, 15, len(instructionMatrix.matrix))
	assert.Equal(t, 11, len(instructionMatrix.matrix[0]))

	assert.Equal(t, 1, instructionMatrix.matrix[10][6])
	assert.Equal(t, 0, instructionMatrix.matrix[10][5])
}

func TestFolding(t *testing.T) {
	lines := []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
		"",
		"fold along y=7",
		"fold along x=5",
	}
	instructionMatrix := CreateInstructionMatrix(lines)
	foldUp := FoldInstruction{
		Axis: "y",
		At: 7,
	}

	foldLeft := FoldInstruction{
		Axis: "x",
		At: 5,
	}
	foldedMatrix := instructionMatrix.Fold(foldUp)
	foldedMatrix.Print()

	println("-------------------------------------")
	foldedMatrix = foldedMatrix.Fold(foldLeft)
	foldedMatrix.Print()
}


