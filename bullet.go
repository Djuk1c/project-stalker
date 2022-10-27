package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	Sprite *Sprite
	Speed  float64
}

func NewBullet(pos Vec2, rot float64) *Bullet {
	bullet := new(Bullet)
	size := iVec2{2, 4}
	img := ebiten.NewImage(size.x, size.y)
	img.Fill(color.RGBA{255, 0, 0, 255})
	bullet.Sprite = NewSprite(img, pos, size, rot)
	bullet.Speed = 0.000001
	return bullet
}

func (b *Bullet) Update(deltaTime float64) {
	rot := b.Sprite.Rot - 90
	b.Sprite.Pos.x += math.Cos(ToRadian(rot)) * b.Speed * deltaTime
	b.Sprite.Pos.y += math.Sin(ToRadian(rot)) * b.Speed * deltaTime
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	b.Sprite.ResetOptions()
	b.Sprite.Options.GeoM.Translate(-float64(b.Sprite.Size.x)/2, -float64(b.Sprite.Size.y)/2) // Rotate axis in center
	b.Sprite.Draw(screen)
}
