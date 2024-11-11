package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bird struct {
	// Physics
	posX, posY   float64
	radius       float64
	velocity     float64
	rotation     float64
	acceleration float64

	dead bool

	// Graphics
	frame            int
	animationCounter int
}

func (b *Bird) Draw(screen *ebiten.Image, scale float64) error {
	b.animationCounter++
	const framesPerTexture = 15
	if b.animationCounter >= framesPerTexture {
		b.animationCounter = 0                      // сброс счётчика
		b.frame = (b.frame + 1) % len(BirdTextures) // переключение кадра
	}

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(-b.radius, -b.radius)
	opts.GeoM.Rotate(Radians(b.rotation))
	opts.GeoM.Scale(scale, scale)
	opts.GeoM.Translate(b.posX*scale, b.posY*scale)

	screen.DrawImage(BirdTextures[b.frame], opts)
	return nil
}

func (b *Bird) Update(deltaTime float64) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		b.velocity = -100
	}

	b.velocity += b.acceleration * deltaTime
	if b.velocity > 200 {
		b.velocity = 200
	}

	b.posY += b.velocity * deltaTime
	if b.posY > 195 {
		b.posY = 195
	}

	if b.velocity < 0 {
		b.rotation -= 600 * deltaTime
		if b.rotation < -20 {
			b.rotation = -20
		}
	} else if b.velocity > 110 || b.dead {
		b.rotation += 480 * deltaTime
		if b.rotation > 90 {
			b.rotation = 90
		}
	}

	fmt.Println(b.velocity)
}
