package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCA(t *testing.T) {
	var c = cpu.CPU{}

	c.Pc = 2
	c.Sp = 4
	opcodes.CA(&c, 0x0972)

	assert.Equal(t, uint16(2), c.Stack[5])
	assert.Equal(t, uint16(0x0970), c.Pc)
}
