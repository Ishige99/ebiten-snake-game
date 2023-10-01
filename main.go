package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
	"image/color"
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

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// ゲームオーバー状態を処理
	if g.gameOver {
		// Rキーが押されたかどうかを判定し、押されていればrestart
		// `IsKeyJustPressed()`: https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2/inpututil#IsKeyJustPressed
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.restart()
		}
		return nil
	}

	// ゲームの更新カウンターを制御。
	// フレームごとに実行をして、速さよりもカウンターが小さければ早期リターンをして、以後の処理をスキップをする。（蛇の速度を制御する目的）
	g.updateCounter++
	if g.updateCounter < g.speed {
		return nil
	}
	g.updateCounter = 0

	// 蛇の移動
	g.snake.Move()

	// キーボード入力を監視して蛇の移動方向を変更
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

	// 蛇の頭部の位置が画面外に出た場合、ゲームオーバーとなる
	head := g.snake.Body[0]
	if head.X < 0 || head.Y < 0 || head.X >= screenWidth/tileSize || head.Y >= screenHeight/tileSize {
		g.gameOver = true
		g.speed = 10
	}

	// 蛇が自分自身と衝突した場合、ゲームオーバーとなる
	for _, part := range g.snake.Body[1:] {
		if head.X == part.X && head.Y == part.Y {
			g.gameOver = true
			g.speed = 10
		}
	}

	// 蛇が食べ物を食べた場合の処理
	if head.X == g.food.Position.X && head.Y == g.food.Position.Y {
		g.score++                // スコア+1
		g.snake.GrowCounter += 1 // 蛇の長さ+1
		g.food = NewFood()       // 新しいfoodの生成

		// 蛇の速度を上げる（早くなりすぎないようには制御：8回まで）
		if g.speed > 2 {
			g.speed--
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ゲーム背景を黒で塗りつぶし
	screen.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 255})

	// 蛇の体を描画
	for _, p := range g.snake.Body {
		// 蛇の1つ1つのセグメントを緑で塗りつぶし
		vector.DrawFilledRect(
			screen,
			float32(p.X*tileSize),
			float32(p.Y*tileSize),
			tileSize,
			tileSize,
			color.RGBA{R: 0, G: 255, B: 0, A: 255},
			true,
		)
	}

	// 食べ物を描画（赤色）
	vector.DrawFilledRect(
		screen,
		float32(g.food.Position.X*tileSize),
		float32(g.food.Position.Y*tileSize),
		tileSize,
		tileSize,
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
		true,
	)

	// ゲームオーバー画面の表示
	face := basicfont.Face7x13
	if g.gameOver {
		text.Draw(screen, "Game Over", face, screenWidth/2-40, screenHeight/2, color.White)
		text.Draw(screen, "Press 'R' to restart", face, screenWidth/2-60, screenHeight/2+16, color.White)
	}

	// スコアの表示
	scoreText := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreText, face, 5, screenHeight-5, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
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

func (g *Game) restart() {
	g.snake = NewSnake()
	g.score = 0
	g.gameOver = false
	g.food = NewFood()
	g.speed = 10
}
