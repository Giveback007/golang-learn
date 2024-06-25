package main

import (
	"github.com/eiannone/keyboard"
)

func KeyListener(funct func(rune)) {
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}

		funct(event.Rune)
	}
}
