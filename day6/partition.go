package main

type Bounds struct {
	Low int
	High int
}

func partition(totalSize int, partitionSize int) []Bounds {
	var allBounds []Bounds
	if partitionSize >= totalSize {
		bound := Bounds{
			Low: 0,
			High: totalSize-1,
		}
		allBounds = append(allBounds, bound)
		return allBounds
	}
	lowerBound := 0
	higherBound := partitionSize-1
	for remainingElements := totalSize; remainingElements > 0; {
		bound := Bounds{
			Low: lowerBound,
			High: higherBound,
		}
		remainingElements -= partitionSize
		lowerBound += partitionSize
		higherBound += partitionSize
		if higherBound >= totalSize {
			higherBound = totalSize - 1
		}
		allBounds = append(allBounds, bound)
	}
	return allBounds
}

func partitionList(fishes []*LanternFish, partitionSize int) [][]*LanternFish {
	var partitions [][]*LanternFish
	boundsList := partition(len(fishes), partitionSize)
	for _, bounds := range boundsList {
		subList := fishes[bounds.Low:bounds.High+1]
		partitions = append(partitions, subList)
	}
	return partitions
}