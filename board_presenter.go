package main

import (
	"fmt"
	"strings"
)

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func DrawBoard(board *Board) {
	ClearTerminal()
	output := ""
	for y := range board.cells {
		row := ""
		for x := range board.cells[y] {
			if board.cells[y][x].isAlive {
				row += "■"
			} else {
				row += "□"
			}
			row += " "
		}
		row = strings.Trim(row, " ")
		output += (row + "\n")
	}
	fmt.Println(output)
}
