package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPacketEvaluator_Evaluate(t *testing.T) {
	parser := BitsParser{}
	packet := parser.ParseInput("9C0141080250320F1802104A08")
	assert.NotNil(t, packet)

	evaluator := PacketEvaluator{}
	value := evaluator.Evaluate(packet)
	assert.Equal(t, 1, value)
}
