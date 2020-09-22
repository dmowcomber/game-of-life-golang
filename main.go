package main

import (
	"math/rand"
	"time"

	"github.com/dmowcomber/game-of-life-golang/gol"
)

func main() {
	board := gol.NewBoard()
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 2410; i++ {
		x := rand.Intn(130)
		y := rand.Intn(38)
		board.SetAlive(gol.NewPoint(x, y))
	}

	for {
		board.Print(130, 38)
		time.Sleep(10 * time.Millisecond)
		board = board.Next()
	}
}
