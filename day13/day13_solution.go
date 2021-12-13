package main

import "adventoccode2021/commons"

func main() {
	lines, err := commons.ReadFile("input/day13.txt")
	if err != nil {
		panic(err)
	}
	originalMatrix := CreateInstructionMatrix(lines)
	foldInstructions := ParseFoldInstructions(lines)

	part1(originalMatrix, foldInstructions)
	part2(originalMatrix, foldInstructions)
}

func part2(originalMatrix *InstructionMatrix, foldInstructions []FoldInstruction) {
	var foldedMatrix = originalMatrix
	for _, foldInstruction := range foldInstructions {
		foldedMatrix = foldedMatrix.Fold(foldInstruction)
	}
	foldedMatrix.Print()
}

func part1(matrix *InstructionMatrix, foldInstructions []FoldInstruction) {
	foldedMatrix := matrix.Fold(foldInstructions[0])

	sum := 0
	for y:=0;y<len(foldedMatrix.matrix);y++ {
		for x:=0;x<len(foldedMatrix.matrix[0]);x++ {
			sum += foldedMatrix.matrix[y][x]
		}
	}

	println(sum)
}

