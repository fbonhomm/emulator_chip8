package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDDTVX(t *testing.T) {
	var c = cpu.CPU{}

	c.V[2] = 134
	opcodes.LDDTVX(&c, 0x0200)

	assert.Equal(t, uint8(134), c.SystemTimer)
}
