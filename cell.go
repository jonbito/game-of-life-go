package main

import (
	"math/rand"
)

type Cell struct {
	isAlive bool
}

func (c *Cell) NextGeneration(neighbors []Cell) {
	aliveNeighbors := []Cell{}
	for i := range neighbors {
		if neighbors[i].isAlive {
			aliveNeighbors = append(aliveNeighbors, neighbors[i])
		}
	}

	if len(aliveNeighbors) == 3 || (c.isAlive && len(aliveNeighbors) == 2) {
		c.isAlive = true
	} else {
		c.isAlive = false
	}
}

func (c *Cell) Randomize(r *rand.Rand) {
	if r.Float32() < .5 {
		c.isAlive = false
	} else {
		c.isAlive = true
	}
}
