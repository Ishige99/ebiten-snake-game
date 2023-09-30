package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/rand"
	"log"
	"time"
)

const (
	// game screen size: pixel
	screenWidth  = 320
	screenHeight = 240

	tileSize = 5
)

type Point struct {
	// x and y coordinates struct
	X int
	Y int
}

type Snake struct {
	// snake struct
	Body        []Point
	Direction   Point
	GrowCounter int
}

func main() {
	// generate seed values to add game randomness
	rand.Seed(uint64(time.Now().UnixNano()))

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

func NewSnake() *Snake {
	// initialize snake
	return &Snake{
		// set window center position
		Body: []Point{
			{
				X: screenWidth / tileSize / 2,
				Y: screenHeight / tileSize / 2,
			},
		},
		// set the snake's direction of movement to the right
		Direction: Point{
			X: 1,
			Y: 0,
		},
	}
}

func (s *Snake) Move() {
	// 新しい進行方向の座標を作成
	newHead := Point{
		X: s.Body[0].X + s.Direction.X,
		Y: s.Body[0].Y + s.Direction.Y,
	}
	// Bodyの先頭にnewHeadを追加
	// sliceTricks: https://github.com/golang/go/wiki/SliceTricks
	s.Body = append([]Point{newHead}, s.Body...)

	// GrowCounterが1であれば追加なのでBodyの変更は無し
	// GrowCounterが0の場合は追加無しなので、蛇の長さを1つ減らす（リセット）
	if s.GrowCounter > 0 {
		s.GrowCounter--
	} else {
		s.Body = s.Body[:len(s.Body)-1]
	}
}
