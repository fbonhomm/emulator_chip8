package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSEVXB_SET_PC(t *testing.T) {
	var c = cpu.CPU{}

	c.V[3] = 0xEC
	opcodes.SEVXB(&c, 0x03EC)

	assert.NotEqual(t, uint8(1), c.Pc)
}

func TestSEVXB_NOT_SET_PC(t *testing.T) {
	var c = cpu.CPU{}

	c.V[3] = 0xEB
	opcodes.SEVXB(&c, 0x03EC)

	assert.NotEqual(t, uint8(0), c.Pc)
}
