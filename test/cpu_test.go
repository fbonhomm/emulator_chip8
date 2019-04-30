package test

import (
	cpu "emulator/src"
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	// var c cpu.CPU
	var c = cpu.CPU{}

	result := c.identifyOpcode()

	fmt.Printf("%d\n\n", result)
}
