package main

func main() {

	targetArea := TargetArea{
		Xmin: 241,
		Xmax: 273,
		Ymin: -97,
		Ymax: -63,
	}
	part1(targetArea)
	part2(targetArea)
}

func part1(targetArea TargetArea) {
	yMax := -1

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			success, highestY := reachesTarget(i, j, targetArea)
			if success {
				if highestY > yMax {
					yMax = highestY
				}
			}
		}
	}

	println(yMax)
}

func part2(targetArea TargetArea) {
	successCount := 0
	for i := 0; i < 1000; i++ {
		for j := -1000; j < 1000; j++ {
			success, _ := reachesTarget(i, j, targetArea)
			if success {
				successCount++
			}
		}
	}
	println(successCount)
}

func reachesTarget(velx int, vely int, targetArea TargetArea) (bool, int)  {
	x := 0
	y := 0
	success := false
	highestY := -1
	for ;; {
		x += velx
		y += vely
		if y > highestY {
			highestY = y
		}
		if targetArea.contains(x, y)  {
			success = true
			break
		}
		if targetArea.overshot(x, y) {
			break
		}
		vely--
		if velx > 0 {
			velx--
		} else if velx < 0 {
			velx++
		}
	}
	return success, highestY
}