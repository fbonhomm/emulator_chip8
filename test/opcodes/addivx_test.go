package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestADDIVX(t *testing.T) {
	var c = cpu.CPU{}

	c.V[2] = 2
	opcodes.ADDIVX(&c, 0x0200)

	assert.Equal(t, uint16(2), c.I)
}
