package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLDIA(t *testing.T) {
	var c = cpu.CPU{}

	opcodes.LDIA(&c, 0x0234)

	assert.Equal(t, uint16(0x0234), c.I)
}
