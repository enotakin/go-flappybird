package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	posX, posY float64
}

func NewBackground() *Background {
	return &Background{}
}

func (b Background) Draw(screen *ebiten.Image, scale float64) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(b.posX, b.posY)
	opts.GeoM.Scale(scale, scale)

	screen.DrawImage(BackgroundTexture, opts)
}

func (b Background) Update(delta float64) {
}
