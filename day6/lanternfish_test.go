package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DayPassed(t *testing.T) {
	fish := CreateFish(2)
	newFish := fish.DayPassed()
	assert.Nil(t, newFish)
	assert.Equal(t, 1, fish.GetCounter())

	newFish = fish.DayPassed()
	assert.Nil(t, newFish)
	assert.Equal(t, 0, fish.GetCounter())

	newFish = fish.DayPassed()
	assert.NotNil(t, newFish)
	assert.Equal(t, 6, fish.GetCounter())

	assert.Equal(t, 8, newFish.GetCounter())
}
