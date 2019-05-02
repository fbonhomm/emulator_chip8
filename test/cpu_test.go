package cpu_test

import (
	"emulator/src/cpu"
	"testing"

	"github.com/stretchr/testify/assert"
)

// const OPCODE [] uint16 {
//   0x0NNN, 0x00E0, 0x00EE, 0x1NNN, 0x2NNN, 0x3XKK, 0x4XKK, 0x5XY0, 0x6XKK,
//   0x7XKK, 0x8XY0, 0x8XY1, 0x8XY2, 0x8XY3, 0x8XY4, 0x8XY5, 0x8XY6, 0x8XY7,
//   0x8XYE, 0x9XY0, 0xANNN, 0xBNNN, 0xCXKK, 0xDXYN, 0xEX9E, 0xEXA1, 0xFX07,
//   0xFX0A, 0xFX15, 0xFX18, 0xFX1E, 0xFX29, 0xFX33, 0xFX55, 0xFX65,
// }

var opcodeID = []uint16{
	0x0FFF, 0x00E0, 0x00EE, 0x1aaa, 0x2aaa, 0x3aaa, 0x4aaa, 0x5aa0, 0x6aaa,
	0x7aaa, 0x8aa0, 0x8aa1, 0x8aa2, 0x8aa3, 0x8aa4, 0x8aa5, 0x8aa6, 0x8aa7,
	0x8aaE, 0x9aa0, 0xAaaa, 0xBaaa, 0xCaaa, 0xDaaa, 0xEa9E, 0xEaA1, 0xFa07,
	0xFa0A, 0xFa15, 0xFa18, 0xFa1E, 0xFa29, 0xFa33, 0xFa55, 0xFa65,
}

func TestIdentifyOpcode(t *testing.T) {
	var c = cpu.CPU{}

	for idx, op := range opcodeID {
		result := c.IdentifyOpcode(op)
		if op == 0x0FFF {
			idx = 35
		}
		assert.Equal(t, result, uint8(idx))
	}
}
