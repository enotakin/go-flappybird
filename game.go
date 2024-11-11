package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameElement interface {
	Draw(screen *ebiten.Image, scale float64)

	Update(delta float64)
}

type Game struct {
	elements []GameElement

	running bool
}

func NewGame() *Game {
	game := &Game{running: true}
	game.elements = append(game.elements, NewBackground(), NewBarrier(), NewBird(), NewGround())
	return game
}

func (g *Game) Update() error {
	if g.running {
		for _, e := range g.elements {
			e.Update(deltaTime)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, e := range g.elements {
		e.Draw(screen, scale)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
