package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDVXVY(t *testing.T) {
	var c = cpu.CPU{}

	c.V[2] = 15
	opcodes.LDVXVY(&c, 0x0320)

	assert.Equal(t, uint8(15), c.V[3])
}
