package cpu

import (
	"github.com/fbonhomm/emulator_chip8/src/keyboard"
)

// SKNPVX ExA1 - Skip next instruction if key with the value of Vx is not pressed.
func (c *CPU) SKNPVX(opcode uint16) {
	x := uint8((opcode & 0x0F00) >> 8)

	if keyboard.CheckKey(c.V[x]) == 0 {
		c.Pc += 2
	}
}
