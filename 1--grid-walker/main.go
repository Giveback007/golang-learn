package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var prtLn = fmt.Println

const (
	width  = 20
	height = 10
)

type GameGrid [height][width]int

type Loc struct {
	cordY int
	cordX int
}

func moveCursorTo(y, x int) {
	fmt.Printf("\033[%d;%dH", y+1, x+1)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J") // Clear the screen
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

func draw(grid GameGrid) {
	for y := 0; y < height; y++ {
		moveCursorTo(y, 0)
		drawLine(grid[y])
	}
}

func drawLine(line [width]int) {
	for x := 0; x < len(line); x++ {
		if line[x] == 1 {
			fmt.Print("X")
		} else {
			fmt.Print("-")
		}
	}

	prtLn()
}

func main() {
	var player = Loc{1, 1}
	var grid = GameGrid{}

	clearScreen()
	hideCursor()
	defer showCursor()

	fmt.Println("Press ESC to quit")

	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		grid = GameGrid{} // Clear the grid
		grid[player.cordY][player.cordX] = 1

		moveCursorTo(0, 0)
		draw(grid)

		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}

		switch event.Rune {
		case 'w':
			if player.cordY > 0 {
				player.cordY--
			}
		case 's':
			if player.cordY < height-1 {
				player.cordY++
			}
		case 'd':
			if player.cordX < width-1 {
				player.cordX++
			}
		case 'a':
			if player.cordX > 0 {
				player.cordX--
			}
		case rune(0):
			{
				clearScreen()
				showCursor()
				prtLn("Game EXIT")
				return
			}
		}
	}
}
