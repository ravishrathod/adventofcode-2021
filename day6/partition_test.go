package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_partition(t *testing.T) {
	boundsList := partition(10, 3)

	assert.Equal(t, 4, len(boundsList))
	assert.Equal(t, createBounds(0,2), boundsList[0])
	assert.Equal(t, createBounds(3,5), boundsList[1])
	assert.Equal(t, createBounds(6,8), boundsList[2])
	assert.Equal(t, createBounds(9,9), boundsList[3])
}

func Test_partition_biggerThanListSize(t *testing.T) {
	boundsList := partition(10, 11)

	assert.Equal(t, 1, len(boundsList))
	assert.Equal(t, createBounds(0,9), boundsList[0])
}

func Test_partitionList(t *testing.T) {
	list := []*LanternFish{
		CreateFish(3),
		CreateFish(4),
		CreateFish(3),
		CreateFish(1),
		CreateFish(2),
		}
	partition := partitionList(list, 1000)
	assert.Equal(t, 1, len(partition))
	assert.Equal(t, 5, len(partition[0]))
}

func Test_partitionListWithMoreThanOneSubList(t *testing.T) {
	list := []*LanternFish{
		CreateFish(3),
		CreateFish(4),
		CreateFish(3),
		CreateFish(1),
		CreateFish(2),
		CreateFish(5),
		CreateFish(6),
		CreateFish(1),
	}
	partition := partitionList(list, 3)
	assert.Equal(t, 3, len(partition))
	assert.Equal(t, 3, len(partition[0]))
	assert.Equal(t, 3, len(partition[1]))
	assert.Equal(t, 2, len(partition[2]))
}

func createBounds(lower int, higher int) Bounds {
	return Bounds{
		Low: lower,
		High: higher,
	}
}
