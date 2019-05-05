package test

import (
	"emulator/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHRVX_VF_1(t *testing.T) {
	var c = cpu.CPU{}

	c.V[3] = 0xFF
	c.SHRVX(0x0300)

	assert.Equal(t, uint8(0x7F), c.V[3])
	assert.Equal(t, uint8(1), c.V[0xF])
}

func TestSHRVX_VF_0(t *testing.T) {
	var c = cpu.CPU{}

	c.V[3] = 0xF0
	c.SHRVX(0x0300)

	assert.Equal(t, uint8(0x78), c.V[3])
	assert.Equal(t, uint8(0), c.V[0xF])
}
