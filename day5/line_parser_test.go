package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseLines(t *testing.T) {
	textLines := []string {"599,531 -> 599,32", "845,552 -> 596,801", "535,556 -> 349,556"}
	lines := parseLines(textLines)
	assert.Equal(t, 3, len(lines))

	firstLine := lines[0]
	assert.Equal(t, 599, firstLine.start.X)
	assert.Equal(t, 531, firstLine.start.Y)

	assert.Equal(t, 599, firstLine.end.X)
	assert.Equal(t, 32, firstLine.end.Y)
}
