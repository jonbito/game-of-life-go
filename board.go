package main

import (
	"math/rand"
	"time"
)

type Board struct {
	rowCount int
	colCount int
	cells    [][]Cell
}

func (b *Board) SeedBoard() {
	nextCells := [][]Cell{}
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for y := 0; y < b.rowCount; y++ {
		row := []Cell{}
		for x := 0; x < b.colCount; x++ {
			newCell := Cell{}
			newCell.Randomize(r)
			row = append(row, newCell)
		}
		nextCells = append(nextCells, row)
	}

	b.cells = nextCells
}

func (b *Board) Run() {
	nextCells := [][]Cell{}

	for y := range b.cells {
		for x := range b.cells[y] {
			b.cells[y][x].NextGeneration(b.getNeighborsOf(x, y))
		}
	}
	for y := 0; y < b.rowCount; y++ {
		row := []Cell{}
		for x := 0; x < b.colCount; x++ {
			newCell := Cell{}
			newCell.NextGeneration(b.getNeighborsOf(x, y))
			row = append(row, newCell)
		}
		nextCells = append(nextCells, row)
	}

	b.cells = nextCells
}

func (b *Board) getNeighborsOf(x int, y int) []Cell {
	neighbors := []Cell{}
	for i := -1; i <= 1; i++ {
		if y+i <= -1 || b.rowCount <= y+i {
			continue
		}
		for j := -1; j <= 1; j++ {
			if (x+j <= -1 || b.colCount <= x+j) || (i == 0 && j == 0) {
				continue
			}
			neighbors = append(neighbors, b.cells[y+i][x+j])
		}
	}
	return neighbors
}
