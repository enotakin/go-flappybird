package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	screenWidth  = 160
	screenHeight = 240
	scale        = 3
	speed        = 30
)

var (
	deltaTime = 1 / float64(ebiten.TPS())
)

func main() {
	ebiten.SetWindowSize(screenWidth*scale, screenHeight*scale)
	ebiten.SetWindowTitle("FlappyBird")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
