package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitsParser_ParseTypeAndVersion(t *testing.T) {
	parser := BitsParser{}
	packet := parser.ParseInput("D2FE28")

	assert.Equal(t, 6, packet.version)
	assert.Equal(t, 4, packet.typ)
}

func TestBitsParser_ParseValueForLiteralPacket(t *testing.T) {
	parser := BitsParser{}
	packet := parser.ParseInput("D2FE28")

	assert.Equal(t, 2021, packet.value)
}

func TestBitsParser_ParseTypeZero(t *testing.T) {
	parser := BitsParser{}
	packet := parser.ParseInput("38006F45291200")
	assert.NotNil(t, packet)
	assert.Equal(t, 1, packet.version)
	assert.Equal(t, 6, packet.typ)

	assert.Equal(t, 2, len(packet.subPackets))
	assert.Equal(t, 10, packet.subPackets[0].value)
	assert.Equal(t, 20, packet.subPackets[1].value)
}

func TestBitsParser_ParseTypeOne(t *testing.T) {
	parser := BitsParser{}
	packet := parser.ParseInput("EE00D40C823060")
	assert.NotNil(t, packet)
	assert.Equal(t, 7, packet.version)
	assert.Equal(t, 3, packet.typ)

	assert.Equal(t, 3, len(packet.subPackets))
	assert.Equal(t, 1, packet.subPackets[0].value)
	assert.Equal(t, 2, packet.subPackets[1].value)
	assert.Equal(t, 3, packet.subPackets[2].value)
}

func TestBitsParser_test(t *testing.T) {
	parser := BitsParser{}
	packet := parser.ParseInput("8A004A801A8002F478")
	assert.NotNil(t, packet)

	assert.Equal(t, 2, packet.typ)
	assert.Equal(t, 1, len(packet.subPackets))
	child := packet.subPackets[0]
	assert.Equal(t, 2, child.typ)
	assert.Equal(t, 1, len(child.subPackets))
	child = child.subPackets[0]
	assert.Equal(t, 2, child.typ)
	assert.Equal(t, 1, len(child.subPackets))
	child = child.subPackets[0]
	assert.Equal(t, 4, child.typ)
	assert.Equal(t, 0, len(child.subPackets))

	sum := addVersions(0, *packet)
	assert.Equal(t, 16, sum)
}

func TestBitsParser_test1(t *testing.T) {
	parser := BitsParser{}
	packet := parser.ParseInput("620080001611562C8802118E34")
	assert.NotNil(t, packet)

	assert.Equal(t, 2, packet.typ)
	assert.Equal(t, 2, len(packet.subPackets))
	//child := packet.subPackets[0]
	//assert.Equal(t, 2, child.typ)
	//assert.Equal(t, 1, len(child.subPackets))
	//child = child.subPackets[0]
	//assert.Equal(t, 2, child.typ)
	//assert.Equal(t, 1, len(child.subPackets))
	//child = child.subPackets[0]
	//assert.Equal(t, 4, child.typ)
	//assert.Equal(t, 0, len(child.subPackets))

	sum := addVersions(0, *packet)
	assert.Equal(t, 12, sum)
}

