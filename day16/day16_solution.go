package main

import (
	"adventoccode2021/commons"
)

func main() {
	lines, err := commons.ReadFile("input/day16.txt")
	if err != nil {
		panic(err)
	}
	bitParser := BitsParser{}
	packet := bitParser.ParseInput(lines[0])
	counter := &Counter{}
	addVersions(counter, *packet)
	println(counter.ticker)
}

func addVersions(counter *Counter , packet Packet) {
	for _, subPacket := range packet.subPackets {
		addVersions(counter, *subPacket)
	}
	counter.Add(packet.version)
}

type Counter struct {
	ticker int
}

func (c *Counter) Add(value int) {
	c.ticker += value
}
