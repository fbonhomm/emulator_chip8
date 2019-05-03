package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSNEVXVY_PC_2(t *testing.T) {
	var c = cpu.CPU{}

	c.V[8] = 0xDD
	c.V[3] = 0xED
	opcodes.SNEVXVY(&c, 0x0830)

	assert.Equal(t, uint16(2), c.Pc)
}

func TestSNEVXVY_PC_0(t *testing.T) {
	var c = cpu.CPU{}

	c.V[8] = 0xED
	c.V[3] = 0xED
	opcodes.SNEVXVY(&c, 0x0830)

	assert.Equal(t, uint16(0), c.Pc)
}
