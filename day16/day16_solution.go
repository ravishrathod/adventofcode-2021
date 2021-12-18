package main

import "adventoccode2021/commons"

func main() {
	lines, err := commons.ReadFile("input/day16.txt")
	if err != nil {
		panic(err)
	}
	bitParser := BitsParser{}
	packet := bitParser.ParseInput(lines[0])
	println(addVersions(0, *packet))
}

func addVersions(sum int , packet Packet) int {
	for _, subPacket := range packet.subPackets {
		sum += addVersions(sum, *subPacket)
	}
	sum += packet.version
	return sum
}
