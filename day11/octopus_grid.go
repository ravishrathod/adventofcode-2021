package main

import (
	"fmt"
	"strconv"
)

type OctopusGrid struct {
	grid                      [][]int
	ticker                    int
	flashes                   int
	height                    int
	width                     int
	flashedByCoords           map[string]bool
}

func (this *OctopusGrid) Tick(count int) {
	this.height = len(this.grid)
	this.width = len(this.grid[0])
	for i := this.ticker; i < (this.ticker + count); i++ {
		this.flashedByCoords = make(map[string]bool)
		for y := 0; y < this.height; y++ {
			for x := 0; x < this.width; x++ {
				this.grid[y][x] = this.grid[y][x] + 1
			}
		}
		for y := 0; y < this.height; y++ {
			for x := 0; x < this.width; x++ {
				this.handleFlash(x, y)
			}
		}
	}
}

func (this *OctopusGrid) FirstSynchronousFlash() int {
	this.height = len(this.grid)
	this.width = len(this.grid[0])
	for i := 0; ; i++ {
		this.flashedByCoords = make(map[string]bool)
		for y := 0; y < this.height; y++ {
			for x := 0; x < this.width; x++ {
				this.grid[y][x] = this.grid[y][x] + 1
			}
		}
		for y := 0; y < this.height; y++ {
			for x := 0; x < this.width; x++ {
				this.handleFlash(x, y)
			}
		}
		if len(this.flashedByCoords) >= this.width * this.height {
			return i+1
		}
	}
}

func (this *OctopusGrid) flash(x int, y int) {
	this.flashes++
	this.grid[y][x] = 0
	this.markFlashed(x, y)
	if x > 0 {
		this.incrementEnergy(x-1, y)
		if y > 0 {
			this.incrementEnergy(x-1, y-1)
		}
		if y < this.height-1 {
			this.incrementEnergy(x-1, y+1)
		}
	}
	if x < this.width-1 {
		this.incrementEnergy(x+1, y)
		if y > 0 {
			this.incrementEnergy(x+1, y-1)
		}
		if y < this.height-1 {
			this.incrementEnergy(x+1, y+1)
		}
	}
	if y > 0 {
		this.incrementEnergy(x, y-1)
	}
	if y < this.height-1 {
		this.incrementEnergy(x, y+1)
	}
}

func (this *OctopusGrid) incrementEnergy(x int, y int) {
	if this.isAlreadyFlashed(x, y) {
		return
	}
	this.grid[y][x] = this.grid[y][x] + 1
	if this.grid[y][x] > 9 && !this.isAlreadyFlashed(x, y) {
		this.flash(x, y)
	}
}

func (this *OctopusGrid) handleFlash(x int, y int) {
	energyLevel := this.grid[y][x]
	if energyLevel > 9 && !this.isAlreadyFlashed(x, y) {
		this.flash(x, y)
	}
}

func (this *OctopusGrid) markFlashed(x int, y int) {
	this.flashedByCoords[toStringCoords(x, y)] = true
}

func (this *OctopusGrid) isAlreadyFlashed(x int, y int) bool {
	return this.flashedByCoords[toStringCoords(x, y)]
}

func toStringCoords(x int, y int) string {
	return strconv.Itoa(x)+","+strconv.Itoa(y)
}

func (this *OctopusGrid) GetFlashes() int {
	return this.flashes
}

func (this *OctopusGrid) Print() {
	for i := 0; i < this.width; i++ {
		for j := 0; j < this.height; j++ {
			fmt.Printf(strconv.Itoa(this.grid[i][j]) + " ")
		}
		fmt.Print("\n")
	}
}
