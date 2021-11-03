package main

import (
	"flag"
	"time"
)

func main() {
	rowFlag := flag.Int("rows", 15, "Specify the number of rows on the board")
	colFlag := flag.Int("cols", 15, "Specify the number of cols on the board")
	flag.Parse()

	b := Board{rowCount: *rowFlag, colCount: *colFlag}
	b.SeedBoard()
	for {
		b.Run()
		DrawBoard(&b)
		time.Sleep(500 * time.Millisecond)
	}
}
