package main

type CaveExtrapolater struct {
	original [][]int
	xChunkSize int
	yChunkSize int
}

func NewCaveExtrapolater(original [][]int) *CaveExtrapolater {
	yChunkSize := len(original)
	xChunkSize := len(original[0])

	return &CaveExtrapolater{
		original: original,
		xChunkSize: xChunkSize,
		yChunkSize: yChunkSize,
	}
}

func (ce *CaveExtrapolater) Extrapolate() [][]int {

	height := ce.yChunkSize * 5
	width := ce.xChunkSize * 5

	matrix := make([][]int, height)
	for i, _ := range matrix {
		matrix[i] = make([]int, width)
	}

	for y:=0;y<ce.yChunkSize;y++ {
		for x:=0;x<ce.xChunkSize;x++ {
			matrix[y][x] = ce.original[y][x]
		}
	}
	for ver :=0; ver <5; ver++ {
		if ver > 0 {
			ce.createCavePartToBottom(0, (ver-1) * ce.yChunkSize, matrix)
		}
		for hor:=0;hor<4;hor++ {
			ce.createCavePartToRight(hor * ce.xChunkSize,ver * ce.yChunkSize, matrix)
		}
	}
	return matrix
}

func (ce *CaveExtrapolater) createCavePartToRight(sourceStartX int, sourceStartY int,matrix [][]int)  {
	targetX := sourceStartX + ce.xChunkSize
	targetY := sourceStartY
	for y:=sourceStartY;y<(sourceStartY + ce.yChunkSize);y++ {
		targetX = sourceStartX + ce.xChunkSize
		for x:=sourceStartX;x<(sourceStartX + ce.xChunkSize);x++ {
			value := matrix[y][x] + 1
			if value > 9 {
				value = 1
			}
			matrix[targetY][targetX] = value
			targetX++
		}
		targetY++
	}
}

func (ce *CaveExtrapolater) createCavePartToBottom(sourceStartX int, sourceStartY int, matrix [][]int) {
	targetX := sourceStartX
	targetY := sourceStartY + ce.yChunkSize

	for y:=sourceStartY;y<(sourceStartY + ce.yChunkSize);y++ {
		targetX = sourceStartX
		for x:=sourceStartX;x<(sourceStartX + ce.xChunkSize);x++ {
			value := matrix[y][x] + 1
			if value > 9 {
				value = 1
			}
			matrix[targetY][targetX] = value
			targetX++
		}
		targetY++
	}
}
