package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSNEVXB_PC_2(t *testing.T) {
	var c = cpu.CPU{}

	c.V[8] = 0xDD
	opcodes.SNEVXB(&c, 0x08ED)

	assert.Equal(t, uint16(2), c.Pc)
}

func TestSNEVXB_PC_0(t *testing.T) {
	var c = cpu.CPU{}

	c.V[8] = 0xED
	opcodes.SNEVXB(&c, 0x08ED)

	assert.Equal(t, uint16(0), c.Pc)
}
