package main

var Chars = [3]string{"-", "X", "O"}
var ScoreDict = [4]string{" ", "-", " ", "+"}

// var ScoreDict = map[int]string{-1: "l", 0: "", 1: "w"}

// -- GridState -- //
type SqrSt int

const (
	S SqrSt = 0 // Space
	X SqrSt = 1 // X: Player-1
	O SqrSt = 2 // O: Player-2
)

type GameState [9]SqrSt

// -- Scores -- //
type Scr int

const (
	I Scr = 0 // Impassable
	L Scr = 1 // Loss
	T Scr = 2 // Tie
	W Scr = 3 // Win
)

type Scores [9]Scr
