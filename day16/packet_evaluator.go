package main

import "math"

type PacketEvaluator struct {

}

func (pe *PacketEvaluator) Evaluate(packet *Packet) int {
	return pe.evaluatePacket(packet)
}

func (pe *PacketEvaluator) evaluatePacket(packet *Packet) int {
	if packet.typ == 4 {
		return packet.value
	}
	var operands []int
	for _, subPacket := range packet.subPackets {
		operands = append(operands, pe.evaluatePacket(subPacket))
	}
	switch packet.typ {
	case 0:
		sum := 0
		for _, operand := range operands {
			sum += operand
		}
		return sum
	case 1:
		product := 1
		for _, operand := range operands {
			product *= operand
		}
		return product
	case 2:
		min := math.MaxInt
		for _, operand := range operands {
			if operand < min {
				min = operand
			}
		}
		return min
	case 3:
		max := math.MinInt
		for _, operand := range operands {
			if operand > max {
				max = operand
			}
		}
		return max
	case 5:
		if operands[0] > operands[1] {
			return 1
		}
		return 0
	case 6:
		if operands[0] < operands[1] {
			return 1
		}
		return 0
	case 7:
		if operands[0] == operands[1] {
			return 1
		}
		return 0
	}
	return -1
}
