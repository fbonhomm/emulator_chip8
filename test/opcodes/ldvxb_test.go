package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDVXB(t *testing.T) {
	var c = cpu.CPU{}

	opcodes.LDVXB(&c, 0x0322)

	assert.Equal(t, uint8(0x22), c.V[3])
}
