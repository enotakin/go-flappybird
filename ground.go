package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Ground struct {
	head *GroundSegment
	tail *GroundSegment
}

type GroundSegment struct {
	posX, posY float64
	next       *GroundSegment
}

func NewGround() *Ground {
	g := &Ground{}
	g.AddSegment()
	g.AddSegment()
	return g
}

func (g *Ground) Draw(screen *ebiten.Image, scale float64) {
	for s := g.head; s != nil; s = s.next {
		if s.IsInSight() {
			opts := &ebiten.DrawImageOptions{}

			opts.GeoM.Scale(scale, scale)
			opts.GeoM.Translate(s.posX*scale, s.posY*scale)

			screen.DrawImage(GroundTexture, opts)
		}
	}
}

func (g *Ground) Update(delta float64) {
	for s := g.head; s != nil; s = s.next {
		s.posX -= 30 * delta
	}

	if g.head != nil && !g.head.IsInSight() {
		oldHead := g.head
		oldTail := g.tail

		g.head = oldHead.next
		g.tail.next = oldHead
		g.tail = oldHead
		g.tail.next = nil

		oldHead.posX = oldTail.posX + float64(GroundTexture.Bounds().Dx())
		oldHead.posY = oldTail.posY
	}
}

func (s *GroundSegment) IsInSight() bool {
	bounds := GroundTexture.Bounds()

	minX := s.posX
	minY := s.posY
	maxX := minX + float64(bounds.Dx())
	maxY := minY + float64(bounds.Dy())

	return minX <= screenWidth && maxX >= 0 && minY <= screenHeight && maxY >= 0
}

func (g *Ground) AddSegment() {
	newSegment := &GroundSegment{posY: 200}

	if g.head == nil {
		g.head = newSegment
		g.tail = newSegment
	} else {
		newSegment.posX = g.tail.posX + float64(GroundTexture.Bounds().Dx())

		g.tail.next = newSegment
		g.tail = newSegment
	}
}
