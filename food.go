package main

import "math/rand"

type Food struct {
	Position Point
}

func NewFood() *Food {
	// ランダムな箇所にfoodを作成
	var food Food
	for {
		food.Position.X = rand.Intn(screenWidth / tileSize)
		food.Position.Y = rand.Intn(screenHeight / tileSize)

		// 画面一番端であれば再度生成
		if food.Position.X != 0 &&
			food.Position.Y != 0 &&
			food.Position.X != (screenWidth/tileSize-1) &&
			food.Position.Y != (screenWidth/tileSize-1) {
			break
		}
	}

	return &food
}
