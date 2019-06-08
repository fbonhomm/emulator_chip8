package test

import (
	"github.com/fbonhomm/emulator_chip8/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSNEVXB_PC_2(t *testing.T) {
	var c = cpu.CPU{}

	c.V[8] = 0xDD
	c.SNEVXB(0x08ED)

	assert.Equal(t, uint16(2), c.Pc)
}

func TestSNEVXB_PC_0(t *testing.T) {
	var c = cpu.CPU{}

	c.V[8] = 0xED
	c.SNEVXB(0x08ED)

	assert.Equal(t, uint16(0), c.Pc)
}
