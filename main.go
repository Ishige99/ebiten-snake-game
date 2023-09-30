package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

func (g *Game) Update() error {
	// リスタート
	if g.gameOver {
		// ユーザーの押したキーを判定
		// `IsKeyJustPressed()`: https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2/inpututil#IsKeyJustPressed
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.restart()
		}
		return nil
	}

	//
	g.updateCounter++
	if g.updateCounter < g.speed {
		return nil
	}
	g.updateCounter = 0

	g.snake.Move()

	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.snake.Direction.X == 0 {
		g.snake.Direction = Point{X: -1, Y: 0}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && g.snake.Direction.X == 0 {
		g.snake.Direction = Point{X: 1, Y: 0}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.snake.Direction.Y == 0 {
		g.snake.Direction = Point{X: 0, Y: -1}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && g.snake.Direction.Y == 0 {
		g.snake.Direction = Point{X: 0, Y: 1}
	}

	head := g.snake.Body[0]
	if head.X < 0 || head.Y < 0 || head.X >= screenWidth/tileSize || head.Y >= screenHeight/tileSize {
		g.gameOver = true
		g.speed = 10
	}

	for _, part := range g.snake.Body[1:] {
		if head.X == part.X && head.Y == part.Y {
			g.gameOver = true
			g.speed = 10
		}
	}

	if head.X == g.food.Position.X && head.Y == g.food.Position.Y {
		g.score++
		g.snake.GrowCounter += 1
		g.food = NewFood()
		g.score++
		g.food = NewFood()
		g.snake.GrowCounter = 1

		if g.speed > 2 {
			g.speed--
		}
	}

	return nil
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
