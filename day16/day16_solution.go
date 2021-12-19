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
	println("Sum of Versions: ", counter.ticker)

	evaluator := PacketEvaluator{}
	value := evaluator.evaluatePacket(packet)
	println("Eval result ", value)
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
