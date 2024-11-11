package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
)

var (
	BackgroundTexture = loadTexture("assets/background.png")
	GroundTexture     = loadTexture("assets/ground.png")
	BirdTextures      = loadTextures("assets/bird.png", 4, 17, 12)
)

func loadTexture(path string) *ebiten.Image {
	tex, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
	}
	return tex
}

func loadTextures(path string, amount, width, height int) []*ebiten.Image {
	texture := loadTexture(path)

	textures := make([]*ebiten.Image, amount)
	for index := 0; index < amount; index++ {
		y := index * height
		frame := texture.SubImage(image.Rect(0, y, width, y+height)).(*ebiten.Image)
		textures[index] = frame
	}

	return textures
}
