package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	input1   GameState
	input2   SqrSt
	expected Scores
}

var testCases = []testCase{
	{
		[9]SqrSt{},
		X,
		[9]Scr{2, 2, 2, 2, 2, 2, 2, 2, 2},
	},

	{
		[9]SqrSt{0, 0, 1, 0, 2, 0, 1, 0, 2},
		X,
		[9]Scr{3, 1, 0, 1, 0, 1, 0, 1, 0},
	},

	{ //
		[9]SqrSt{0, 0, 0, 0, 0, 0, 1, 0, 2},
		X,
		[9]Scr{3, 2, 3, 3, 2, 2, 0, 1, 0},
	},

	{
		[9]SqrSt{0, 0, 1, 2, 0, 0, 1, 0, 0},
		O,
		[9]Scr{1, 1, 0, 0, 2, 1, 0, 1, 1},
	},

	{ //
		[9]SqrSt{1, 2, 1, 0, 0, 2, 1, 2, 1},
		O,
		[9]Scr{0, 0, 0, 1, 3, 0, 0, 0, 0},
	},

	{
		[9]SqrSt{2, 0, 1, 0, 0, 0, 1, 0, 2},
		X,
		[9]Scr{0, 1, 0, 1, 3, 1, 0, 1, 0},
	},
}

func TestMinMax(t *testing.T) {
	pass := 0
	fail := 0

	for _, tc := range testCases {
		result := MinMaxScores(tc.input1, tc.input2)

		if result != tc.expected {
			fail++

			t.Error()
			fmt.Printf("\nINPT: %v: %v", Chars[tc.input2], tc.input1)
			fmt.Printf("\nEXPC: %v", tc.expected)
			fmt.Printf("\nOUTP: %v", result)
			fmt.Printf("\n")
		} else {
			pass++
		}
	}

	fmt.Printf("\n [Pass: %v] :: [Fail: %v]\n", pass, fail)
}
