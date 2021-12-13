package main

import (
	"adventoccode2021/commons"
	"fmt"
	"strings"
)

type InstructionMatrix struct {
	matrix [][]int
}

func (this *InstructionMatrix) Fold(foldInstruction FoldInstruction) *InstructionMatrix {
	if foldInstruction.Axis == "y" {
		return this.foldUp(foldInstruction)
	}
	return this.foldLeft(foldInstruction)
}

func (this *InstructionMatrix) foldUp(instruction FoldInstruction) *InstructionMatrix {
	upper := this.matrix[:instruction.At]
	lower := this.matrix[instruction.At+1:]

	height := len(upper)
	width := len(upper[0])
	matrix := make([][]int, height)
	for y,_ := range matrix {
		matrix[y] = make([]int, width)
	}
	for y:=0;y<height;y++ {
		for x:=0;x<width;x++ {
			sum := upper[y][x] + lower[height-1-y][x]
			if sum > 1 {
				sum = 1
			}
			matrix[y][x] = sum
		}
	}
	return &InstructionMatrix{
		matrix: matrix,
	}
}

func (this *InstructionMatrix) foldLeft(instruction FoldInstruction) *InstructionMatrix {
	left := make([][]int, len(this.matrix))
	right := make([][]int, len(this.matrix))

	for y:=0; y<len(this.matrix); y++ {
		left[y] = this.matrix[y][:instruction.At]
		right[y] = this.matrix[y][instruction.At+1:]
	}

	height := len(left)
	width := len(left[0])
	matrix := make([][]int, height)
	for y,_ := range matrix {
		matrix[y] = make([]int, width)
	}
	for y:=0;y<height;y++ {
		for x:=0;x<width;x++ {
			sum := left[y][x] + right[y][width-1-x]
			if sum > 1 {
				sum = 1
			}
			matrix[y][x] = sum
		}
	}
	return &InstructionMatrix{
		matrix: matrix,
	}
}

func (this *InstructionMatrix) Print() {
	height := len(this.matrix)
	width := len(this.matrix[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			value := "."
			if this.matrix[y][x] > 0 {
				value = "#"
			}
			print(value + " ")
		}
		fmt.Print("\n")
	}
}
type Point struct {
	X int
	Y int
}

func CreateInstructionMatrix(lines []string) *InstructionMatrix  {
	maxX := 0
	maxY := 0
	var points []Point
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			break
		}
		xy := commons.LinetoIntArray(line)
		x := xy[0]
		y := xy[1]
		point := Point{
			X: x,
			Y: y,
		}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		points = append(points, point)
	}
	matrix := make([][]int, maxY+1)
	for y, _ := range matrix {
		matrix[y] = make([]int, maxX+1)
	}
	for _, point := range points {
		matrix[point.Y][point.X] = 1
	}
	return &InstructionMatrix{
		matrix: matrix,
	}
}

