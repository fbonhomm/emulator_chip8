package cpu

import (
	"emulator/src/keyboard"
)

// SKPVX Ex9E - Skip next instruction if key with the value of Vx is pressed.
func (c *CPU) SKPVX(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	if keyboard.CheckKey(x) == 1 {
		c.Pc += 2
	}
}
