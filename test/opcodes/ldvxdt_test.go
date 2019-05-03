package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDVXDT(t *testing.T) {
	var c = cpu.CPU{}

	c.SystemTimer = 23
	opcodes.LDVXDT(&c, 0x0300)

	assert.Equal(t, uint8(23), c.V[3])
}
