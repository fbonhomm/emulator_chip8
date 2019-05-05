package main

import (
	"emulator/src/cpu"
	"emulator/src/keyboard"
	"emulator/src/screen"
	"flag"
	"io/ioutil"
	"log"
)

func main() {
	var running = true
	var s = screen.SCREEN{}
	var c = cpu.CPU{Display: &s}
	var gameName = flag.String("game")

	if gameName == nil {
		log.Fatal("Load rom with -game [game name]")
	}

	content, err := ioutil.ReadFile(gameName)
	if err != nil {
		log.Fatal(err)
	}

	s.Initialize()
	c.Initialize(content)
	for running {
		running = keyboard.Event()
		keyboard.Delay()
	}
	s.Destroy()
}
