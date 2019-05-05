package test

import (
	"emulator/src/cpu"
	
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDBVX(t *testing.T) {
	var c = cpu.CPU{}

	c.V[2] = 134
	c.I = 200
	c.LDBVX(0x0200)

	assert.Equal(t, uint8(1), c.Memory[200])
	assert.Equal(t, uint8(3), c.Memory[201])
	assert.Equal(t, uint8(4), c.Memory[202])
}
