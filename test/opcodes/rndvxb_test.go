package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRNDVXB(t *testing.T) {
	var c = cpu.CPU{}

	opcodes.RNDVXB(&c, 0x03EC)

	assert.NotEqual(t, uint8(0), c.V[3])
}
