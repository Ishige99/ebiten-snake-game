package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math/rand"
	"time"
)

const (
	// game screen's pixel size
	screenWidth  = 320
	screenHeight = 240

	// 1 tile's pixel size
	tileSize = 5
)

type Game struct {
	snake         *Snake
	food          *Food
	score         int
	gameOver      bool
	ticks         int
	updateCounter int
	speed         int
}

func main() {
	// generate seed values to add game randomness
	rand.NewSource(time.Now().UnixNano())

	// init game struct
	game := &Game{
		snake:    NewSnake(),
		food:     NewFood(),
		gameOver: false,
		ticks:    0,
		speed:    10,
	}

	// set game screen size
	// `func SetWindowSize(width, height int)`: https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#SetWindowSize
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)

	// set game screen title
	// `func SetWindowTitle(title string)`: https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#SetWindowTitle
	ebiten.SetWindowTitle("Snake Game")

	// start game
	// `func RunGame(game Game) error`: https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#RunGame
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
