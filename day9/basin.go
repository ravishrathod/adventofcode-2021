package main

import (
	"strconv"
)

type Basin struct {
	Points   map[string]bool
	LowPoint LowPoint
}

func (this *Basin) HasPoint(x int, y int) bool {
	formatted := this.formatPoint(x, y)
	return this.Points[formatted]
}

func (this *Basin) AddPoint(x int, y int) {
	formatted := this.formatPoint(x, y)
	this.Points[formatted] = true
}

func (this *Basin) formatPoint(x, y int) string {
	formatted := strconv.Itoa(x) + ", " + strconv.Itoa(y)
	return formatted
}
