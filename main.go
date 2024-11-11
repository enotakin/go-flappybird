package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 160
	screenHeight = 240
	scale        = 3
	groundSpeed  = 0.5
)

type Game struct {
	background *ebiten.Image
	ground     *Ground
	bird       *Bird

	running        bool
	lastUpdateTime time.Time
}

func (g *Game) Update() error {
	if g.running {
		deltaTime := time.Since(g.lastUpdateTime)
		g.lastUpdateTime = time.Now()
		delta := deltaTime.Seconds()

		g.ground.Update(delta)
		g.bird.Update(delta)
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

	// TODO pipe
	g.bird.Draw(screen, scale)
	g.ground.Draw(screen, scale)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth * scale, screenHeight * scale
}

func NewGame() *Game {
	game := &Game{running: true, lastUpdateTime: time.Now()}
	game.ground = NewGround()
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
