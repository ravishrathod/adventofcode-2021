package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLine_GetPointsForVerticalLine(t *testing.T) {
	line := Line{
		start: Point{
			X: 1,
			Y: 2,
		},
		end: Point{
			X: 1,
			Y: 5,
		},
	}

	points := line.getPoints()
	assert.Equal(t, 4, len(points))
	assert.Equal(t, 1, points[1].X)
	assert.Equal(t, 3, points[1].Y)
}

func TestLine_ContainsCoordsForHorizontalLine(t *testing.T) {
	line := CreateLine(
		Point{
			X: 1,
			Y: 2,
		},
		Point{
			X: 4,
			Y: 2,
		},
	)

	points := line.getPoints()
	assert.Equal(t, 4, len(points))
	assert.Equal(t, 2, points[1].X)
	assert.Equal(t, 2, points[1].Y)
}

func TestLine_ContainsCoordsForPositiveSlope(t *testing.T) {
	line := CreateLine(Point{
		X: 310,
		Y: 539,
	}, Point{
		X: 602,
		Y: 831,
	})

	assert.Equal(t, 1, line.slope)
	assert.Equal(t, 229, line.intercept)

	points := line.getPoints()
	assert.Equal(t, 293, len(points))
	assert.Equal(t, 311, points[1].X)
	assert.Equal(t, 540, points[1].Y)
}

func TestLine_ContainsCoordsForNegativeSlope(t *testing.T) {
	line := CreateLine(Point{
		X: 780,
		Y: 295,
	}, Point{
		X: 179,
		Y: 896,
	})

	assert.Equal(t, -1, line.slope)
	assert.Equal(t, 1075, line.intercept)

	points := line.getPoints()
	assert.Equal(t, 602, len(points))
	assert.Equal(t, 180, points[1].X)
	assert.Equal(t, 895, points[1].Y)
}

func Test_GetPoints(t *testing.T) {
	line := CreateLine(Point{
		X: 6,
		Y: 4,
	}, Point{
		X: 2,
		Y: 0,
	})

	points := line.getPoints()
	assert.Equal(t, 5, len(points))
	assert.Equal(t, 6, points[4].X)
	assert.Equal(t, 4, points[4].Y)

}
