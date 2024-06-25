package main

import (
	"errors"
	"fmt"
)

var gridStrRender string = `
 7    |8    |9    
   %s  |  %s  |  %s  
 _____|_____|_____
 4    |5    |6    
   %s  |  %s  |  %s  
 _____|_____|_____
 1    |2    |3    
   %s  |  %s  |  %s  
      |     |     
`

func clearScreen() {
	fmt.Printf(" ")
	fmt.Print("\033[H\033[2J") // Clear the screen
}

func render(gridStr *GameState, scores *Scores) {
	gridInterface := make([]interface{}, len(gridStr))
	for i, v := range gridStr {
		if v == 0 {
			scr := scores[i]
			gridInterface[i] = ScoreDict[scr]
		} else {
			gridInterface[i] = Chars[v]
		}
	}

	clearScreen()
	fmt.Printf(gridStrRender, gridInterface...)
}

func keyHandler(r rune) (int, error) {
	if r == 0 {
		return -99, nil
	}

	n := int(r-'0') - 1
	if n < 0 || n > 8 {
		return 0, errors.New("invalid Idx")
	}

	x := [9]int{6, 7, 8, 3, 4, 5, 0, 1, 2}[n]
	return x, nil
}

func resetGrid(char *SqrSt, gameState *GameState, scores *Scores) {
	*char = X
	*gameState = GameState{}
	*scores = MinMaxScores(*gameState, *char)
	render(gameState, scores)
}

func main() {
	var char SqrSt = X
	var gameState GameState = GameState{}
	var scores Scores = MinMaxScores(gameState, char)

	render(&gameState, &scores)

	KeyListener(func(r rune) {
		i, err := keyHandler(r)
		if i == -99 {
			resetGrid(&char, &gameState, &scores)
			return
		}

		if err != nil || gameState[i] != 0 {
			return
		}

		gameState[i] = char

		if char == X {
			char = O
		} else {
			char = X
		}

		scores = MinMaxScores(gameState, char)
		render(&gameState, &scores)
	})
}
