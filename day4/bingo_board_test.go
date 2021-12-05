package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBingoBoardStruct_ColumnSolved(t *testing.T) {
	lines := []string{"25 29 78 57 69", "47 51  7 21 82", "61 81 99 53 30", "50 80 41 94 46", " 9 37 48 71 91"}
	board := CreateBoard(lines)

	assert.False(t, board.mark(78))
	assert.False(t, board.mark(25))
	assert.False(t, board.mark(7))
	assert.False(t, board.mark(99))
	assert.False(t, board.mark(48))
	assert.True(t, board.mark(41))

}

func TestBingoBoardStruct_RowSolved(t *testing.T) {
	lines := []string{"25 29 78 57 69", "47 51  7 21 82", "61 81 99 53 30", "50 80 41 94 46", " 9 37 48 71 91"}
	board := CreateBoard(lines)

	assert.False(t, board.mark(78))
	assert.False(t, board.mark(25))
	assert.False(t, board.mark(7))
	assert.False(t, board.mark(21))
	assert.False(t, board.mark(47))
	assert.False(t, board.mark(82))
	assert.True(t, board.mark(51))
}

func TestBingoBoardStruct_Score(t *testing.T) {
	lines := []string{"25 29 78 57 69", "47 51  7 21 82", "61 81 99 53 30", "50 80 41 94 46", " 9 37 48 71 91"}
	board := CreateBoard(lines)

	board.mark(78)
	board.mark(25)
	board.mark(7)
	board.mark(99)
	board.mark(48)
	board.mark(41)

	assert.Equal(t, 43419, board.score())

}
