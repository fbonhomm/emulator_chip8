package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRET(t *testing.T) {
	var c = cpu.CPU{}

	c.Sp = 2
	c.Stack[2] = 432
	opcodes.RET(&c)

	assert.Equal(t, uint16(432), c.Pc)
}
