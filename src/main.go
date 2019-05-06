package main

import (
	"emulator/src/cpu"
	"emulator/src/keyboard"
	"emulator/src/screen"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var opBySec = keyboard.GetOpBySec()
	var running = true
	var s = screen.SCREEN{}
	var c = cpu.CPU{}
	// var gameName = flag.String("game", "", "-game [game name]")
	//
	// flag.Parse()
	// if gameName == nil || *gameName == "" {
	// 	log.Fatal("Load rom with -game [game name]")
	// }
	if len(os.Args) != 2 {
		panic("Usage: ./program [roms]")
	}

	gameName := os.Args[1]

	content, err := ioutil.ReadFile(gameName)
	if err != nil {
		log.Fatal(err)
	}

	s.Initialize()
	c.Initialize(content)
	fmt.Println(c.Memory)
	for running {
		running = keyboard.Event()
		for idx := uint32(0); idx < opBySec; idx++ {
			c.InterpreterOpcode(&s)
		}
		s.Apply()
		c.DecreaseTimer()
		keyboard.Delay()
	}
	s.Destroy()
}
