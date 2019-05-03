package test

import (
	"emulator/src/cpu"
	"emulator/src/opcodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJPA(t *testing.T) {
	var c = cpu.CPU{}

	opcodes.JPA(&c, 0x0234)

	assert.Equal(t, uint16(0x0232), c.Pc)
}
