package main

import "github.com/hajimehoshi/ebiten/v2"

type Pipe struct {
	// Physics
	posX, posY     float64
	scaleX, scaleY float64
}

type Barrier struct {
	up   Pipe
	down Pipe
	next *Barrier
}

func NewBarrier() *Barrier {
	return &Barrier{up: Pipe{posX: 0, posY: 50, scaleX: 1, scaleY: -1}, down: Pipe{posX: 0, posY: 100, scaleX: 1, scaleY: 1}}
}

func (p *Pipe) Draw(screen *ebiten.Image, scale float64) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(p.scaleX*scale, p.scaleY*scale)
	opts.GeoM.Translate(p.posX*scale, p.posY*scale)

	screen.DrawImage(PipeTexture, opts)
}

func (b *Barrier) Draw(screen *ebiten.Image, scale float64) {
	b.up.Draw(screen, scale)
	b.down.Draw(screen, scale)
}

func (b *Barrier) Update(delta float64) {
	// TODO: полностью переработать логику
	b.up.posX -= speed * delta
	b.down.posX -= speed * delta

	if b.up.posX < -float64(PipeTexture.Bounds().Dx()) {
		b.up.posX = screenWidth
	}

	if b.down.posX < -float64(PipeTexture.Bounds().Dx()) {
		b.down.posX = screenWidth
	}
}
