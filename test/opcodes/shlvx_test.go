package test

import (
	"emulator/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHLVX_VF_1(t *testing.T) {
	var c = cpu.CPU{}

	c.V[3] = 0xFF
	c.SHLVX(0x0300)

	assert.Equal(t, uint8(0xFE), c.V[3])
	assert.Equal(t, uint8(1), c.V[0xF])
}

func TestSHLVX_VF_0(t *testing.T) {
	var c = cpu.CPU{}

	c.V[3] = 0x0F
	c.SHLVX(0x0300)

	assert.Equal(t, uint8(0x1E), c.V[3])
	assert.Equal(t, uint8(0), c.V[0xF])
}
