package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 160
	screenHeight = 240
	scale        = 3
	speed        = 30
)

type Game struct {
	background *ebiten.Image
	barrier    *Barrier
	bird       *Bird
	ground     *Ground

	running        bool
	lastUpdateTime time.Time
}

func (g *Game) Update() error {
	if g.running {
		deltaTime := 1 / float64(ebiten.TPS())

		g.ground.Update(deltaTime)
		g.barrier.Update(deltaTime)
		g.bird.Update(deltaTime)

		fmt.Println(deltaTime)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render Layers:
	//	- Background
	//	- Pipe
	//	- Bird
	// 	- Ground

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scale, scale)
	screen.DrawImage(BackgroundTexture, opts)

	g.barrier.Draw(screen, scale)
	g.bird.Draw(screen, scale)
	g.ground.Draw(screen, scale)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth * scale, screenHeight * scale
}

func NewGame() *Game {
	game := &Game{running: true, lastUpdateTime: time.Now()}
	game.ground = NewGround()
	game.barrier = NewBarrier()
	game.bird = &Bird{posX: screenWidth * 0.235, posY: screenHeight * 0.5, acceleration: 460, radius: 6}
	return game
}

func main() {
	ebiten.SetWindowSize(screenWidth*scale, screenHeight*scale)
	ebiten.SetWindowTitle("FlappyBird")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
