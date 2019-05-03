package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDSTVX(t *testing.T) {
	var c = cpu.CPU{}

	c.V[3] = uint8(2)
	opcodes.LDSTVX(&c, 0x0300)

	assert.Equal(t, uint8(2), c.SoundTimer)
}
