package main

type TargetArea struct {
	Xmin int
	Xmax int
	Ymin int
	Ymax int
}

func (t *TargetArea) contains(x int, y int) bool {
	return (x >= t.Xmin && x <= t.Xmax) && (y >= t.Ymin && y <= t.Ymax)
}

func (t *TargetArea) overshot(x int, y int) bool {
	return !t.contains(x, y)  && (x > t.Xmax || y < t.Ymin)
}