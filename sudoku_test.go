package sudoku

import (
	"fmt"
	"testing"
)

func TestBoard_New(t *testing.T) {
	t.Run("emptyBoardIsCreated", func(t *testing.T) {
		var board = NewBoard()
		for i := 0; i < 81; i++ {
			if board.cells[i] > 0 {
				t.Fail()
			}
		}
	})
}

func TestBoard_IsValid(t *testing.T) {
	t.Run("emptyBoardIsNotValid", func(t *testing.T) {
		var board = NewBoard()
		if board.IsValid() {
			t.Fail()
		}
	})
}

func TestFromString(t *testing.T) {
	t.Run("simpleBoard", func(t *testing.T) {
		input := "111111111\n222222222\n333333333\n444444444\n555555555\n666666666\n777777777\n888888888\n999999999"
		board := FromString(input)
		if !checkCounts(board.cells, 9) {
			t.Fail()
		}
	})
}

func TestCountValues(t *testing.T) {
	var values = []int{1, 1, 2, 2, 2, 3, 3, 3, 3}
	var expected = map[int]int{1: 2, 2: 3, 3: 4}
	var actual = CountValues(values)
	for index, value := range expected {
		if actual[index] != value {
			fmt.Println("Got ", actual[index], " instead of expected ", value)
			t.Fail()
		}
	}
}
