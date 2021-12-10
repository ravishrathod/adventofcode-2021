package main

type Point struct {
	X int
	Y int
}

type Line struct {
	start     Point
	end       Point
	slope     int
	intercept int
}

func CreateLine(start Point, end Point) *Line {
	slope := 1
	intercept := -1
	if start.Y == end.Y { //horizontal line
		slope = 0
		intercept = start.Y
	} else if start.X == end.X { //verticle line
		slope = 0
	} else if (start.X + start.Y) == (end.X + end.Y) {
		slope = -1
		intercept = start.X + start.Y
	} else {
		slope = 1
		intercept = start.Y - start.X
	}
	return &Line{
		start:     start,
		end:       end,
		slope:     slope,
		intercept: intercept,
	}
}

func (this *Line) IsVertical() bool {
	return this.start.X == this.end.X
}

func (this *Line) IsHorizontal() bool {
	return this.start.Y == this.end.Y
}

func (this *Line) getPoints() []Point {
	var points []Point
	if this.IsVertical() {
		var minY, maxY int
		if this.start.Y > this.end.Y {
			maxY = this.start.Y
			minY = this.end.Y
		} else {
			maxY = this.end.Y
			minY = this.start.Y
		}
		for y := minY; y <= maxY; y++ {
			point := Point{
				X: this.start.X,
				Y: y,
			}
			points = append(points, point)
		}
	} else {
		var minX, maxX int
		if this.start.X < this.end.X {
			minX = this.start.X
			maxX = this.end.X
		} else {
			minX = this.end.X
			maxX = this.start.X
		}
		for x := minX; x <= maxX; x++ {
			point := Point{
				X: x,
				Y: this.slope*x + this.intercept,
			}
			points = append(points, point)
		}
	}
	return points
}
