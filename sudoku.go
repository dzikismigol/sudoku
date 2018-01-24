package sudoku

import (
	"strconv"
	"strings"
)

type board struct {
	cells []int
}

func NewBoard() *board {
	var board = new(board)
	for i := 0; i < 81; i++ {
		board.cells = append(board.cells, 0)
	}
	return board
}

func FromString(input string) *board {
	var board = new(board)
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		for i := 0; i < 9; i++ {
			strconv.Atoi(string(row[i]))
		}
	}
	return board
}

func (b *board) IsValid() bool {
	return checkRows(*b)
}

func checkRows(b board) bool {
	for i := 0; i < 9; i++ {
		if !checkCounts(b.GetRow(i), 1) {
			return false
		}
	}
	return true
}

func checkCounts(array []int, expected int) bool {
	for _, counted := range CountValues(array) {
		if counted != expected {
			return false
		}
	}
	return true
}

func (b *board) GetRow(index int) []int {
	row := make([]int, 0, 9)
	startIndex := index * 9
	for i := 0; i < 9; i ++ {
		row = append(row, b.cells[startIndex+i])
	}
	return row
}

func CountValues(array []int) map[int]int {
	values := make(map[int]int)
	for _, number := range array {
		if values[int(number)] != 0 {
			values[int(number)]++
		} else {
			values[int(number)] = 1
		}
	}

	return values
}
