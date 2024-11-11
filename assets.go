package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
)

var (
	BackgroundTexture = loadTexture("assets/background.png")
	GroundTexture     = loadTexture("assets/ground.png")
	BirdTextures      = loadVertTextures("assets/bird.png", 4, 17, 12)
)

func loadTexture(path string) *ebiten.Image {
	tex, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
	}
	return tex
}

func loadVertTextures(path string, amount, width, height int) []*ebiten.Image {
	texture := loadTexture(path)

	textures := make([]*ebiten.Image, amount)
	for index := 0; index < amount; index++ {
		yOffset := index * height
		frame := texture.SubImage(image.Rect(0, yOffset, width, yOffset+height)).(*ebiten.Image)
		textures[index] = frame
	}

	return textures
}

func loadHorizTextures(path string, amount, width, height int) []*ebiten.Image {
	texture := loadTexture(path)

	textures := make([]*ebiten.Image, amount)
	for index := 0; index < amount; index++ {
		xOffset := index * width
		frame := texture.SubImage(image.Rect(xOffset, 0, xOffset+width, height)).(*ebiten.Image)
		textures[index] = frame
	}

	return textures
}
