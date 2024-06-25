package main

/*
 0: Neither player won

 1: Plr-1 "X" won

 2: Plr-2 "O" won
*/
func Winner(game [9]SqrSt) SqrSt {
	var winCond [8][3]int = [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for _, w := range winCond {
		a, b, c := game[w[0]], game[w[1]], game[w[2]]

		if a > 0 && a == b && a == c {
			return a
		}
	}

	return S
}

func scoreMove(gameState GameState, char SqrSt, moveIdx int) Scr {
	if gameState[moveIdx] != 0 {
		return I // Can't `Sim-move`
	}

	gameState[moveIdx] = char // Sim-move

	if Winner(gameState) == char {
		return W // Sim-move: `win`
	}

	opnt := O // Opnt-char
	if char == O {
		opnt = X
	}

	// Simulate opponent moves
	opntScrs := [3]int{}
	for i, _ := range gameState {
		opntScr := scoreMove(gameState, opnt, i)

		if opntScr == W {
			return L // if opnt `win` Sim-move: `loss`
		}

		// count scores
		opntScrs[opntScr]++
	}

	if opntScrs[T] > 0 {
		return T // if opnt `tie` Sim-move: `tie`
	}

	if opntScrs[L] > 0 {
		return W // if opnt `loss` Sim-move: `win`
	}

	return T // default `tie`
}

func MinMaxScores(gameState GameState, char SqrSt) Scores {
	scores := Scores{}
	for i := 0; i < len(gameState); i++ {
		scores[i] = scoreMove(gameState, char, i)
	}

	return scores
}
