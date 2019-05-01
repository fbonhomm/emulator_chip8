package main

import (
	cpu "emulator/src"
	screen "emulator/src"
)

func main() {
	var s = screen.SCREEN{}
	var c = cpu.CPU{display: s}

	s.Initialize()
	c.SetScreen(s)

	running := true
	for running {
		running = s.Event()
	}
}
