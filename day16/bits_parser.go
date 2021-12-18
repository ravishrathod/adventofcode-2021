package main

import (
	"strconv"
	"strings"
)

var hexToBinaryMap = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

type BitsParser struct {

}

func (bp *BitsParser) ParseInput(input string) *Packet  {
	inputArray := bp.stringToArray(input)
	binaryArray := bp.hexToBinary(inputArray)
	packet := &Packet{
		version: 0,
		typ: -1,
		subPackets: make([]*Packet, 0),
	}
	bp.parseSubPacket(binaryArray, packet)
	return packet.subPackets[0]
}

func (bp *BitsParser) parseSubPacket(input []string, parent *Packet) []string {
	if !bp.hasMore(input) {
		return []string{}
	}
	packet := bp.createPacket(input)
	parent.subPackets = append(parent.subPackets, packet)
	if packet.typ == 4 {
		value, index, _ := bp.parseLiteralPacketValue(input)
		packet.value = value
		input = input[index:]
		return input
	}
	lengthTypeId := input[6]
	readIndex := 6
	if lengthTypeId == "0" {
		lengthBits := input[readIndex+1:(readIndex+15+1)]
		readIndex += 15+1
		subPacketLength := bp.binaryArrayToInt(lengthBits)
		subPacketArray := input[readIndex:(readIndex+subPacketLength)]
		readIndex += subPacketLength
		for ;len(subPacketArray) > 0; {
			subPacketArray = bp.parseSubPacket(subPacketArray, packet)
		}
		return input[readIndex:]
	} else {
		lengthBits := input[readIndex+1:(readIndex+11+1)]
		readIndex += 11+1
		subStackCount := bp.binaryArrayToInt(lengthBits)
		remaining := input[readIndex:]
		for i:=0;i<subStackCount;i++ {
			remaining = bp.parseSubPacket(remaining, packet)
		}
		return remaining
	}
}

func (bp *BitsParser) hasMore(input []string) bool {
	if len(input) < 7 {
		return false
	}
	intVal := bp.binaryArrayToInt(input)
	return intVal > 0
}

func (bp *BitsParser) parseLiteralPacketValue(input []string) (int, int, bool) {
	var bits []string
	length := len(input)
	index :=6
	hasMore := false
	for ;index<length; {
		group := input[index: index+5]
		for i, bit := range group {
			if i != 0 {
				bits = append(bits, bit)
			}
		}
		index += 5
		if group[0] == "0" {
			break
		}
	}
	for i:=index;i<len(input);i++ {
		if input[i] == "1" {
			hasMore = true
			break
		}
	}
	return bp.binaryArrayToInt(bits), index, hasMore
}

func (bp *BitsParser) createPacket(input []string) *Packet {
	version := bp.getVersion(input)
	typ := bp.getType(input)
	return &Packet{
		version:    version,
		typ:        typ,
		subPackets: make([]*Packet, 0),
	}
}

func (bp *BitsParser) getVersion(input []string) int  {
	firstThreeBits := input[:3]
	return bp.binaryArrayToInt(firstThreeBits)
}

func (bp *BitsParser) getType(input []string) int {
	bits := input[3:6]
	return bp.binaryArrayToInt(bits)
}

func (bp *BitsParser) binaryArrayToInt(bits []string) int {
	binary := strings.Join(bits, "")
	decimalValue, _ := strconv.ParseInt(binary, 2, 0)
	return int(decimalValue)
}

func (bp *BitsParser) hexToBinary(input []string) []string {
	var binary []string
	for _, hex := range input {
		bits := hexToBinaryMap[hex]
		for _, bit := range []rune(bits) {
			binary = append(binary, string(bit))
		}
	}
	return binary
}

func (bp *BitsParser) stringToArray(input string) []string {
	var array []string
	for _, char := range []rune(input) {
		array = append(array, string(char))
	}
	return array
}

type Packet struct {
	version    int
	typ        int
	value      int
	subPackets []*Packet
}
