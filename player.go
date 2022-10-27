package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Sprite *Sprite
	Speed  float64
}

func NewPlayer() *Player {
	player := new(Player)
	img, _, _ := ebitenutil.NewImageFromFile("assets/player.png")
	pos := Vec2{150, 150}
	size := iVec2{32, 32}
	player.Sprite = NewSprite(img, pos, size, 0)
	player.Sprite.AddAnimation("idle", 0, 5, 125)
	player.Sprite.AddAnimation("run", 1, 8, 75)
	player.Sprite.PlayAnimation("idle")
	player.Speed = 0.0000002
	return player
}

func (p *Player) Update(deltaTime float64) {
	// Rotation
	cx, cy := ebiten.CursorPosition()
	rot := GetAngleBetween(p.Sprite.Pos.x, p.Sprite.Pos.y, float64(cx), float64(cy))
	p.Sprite.Rot = rot

	// Movement
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Sprite.Pos.y += p.Speed * deltaTime
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Sprite.Pos.y -= p.Speed * deltaTime
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Sprite.Pos.x -= p.Speed * deltaTime
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Sprite.Pos.x += p.Speed * deltaTime
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		b := NewBullet(p.Sprite.Pos, rot)
		bullets = append(bullets, b)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.Sprite.ResetOptions()
	p.Sprite.Options.GeoM.Translate(-float64(p.Sprite.Size.x)/2, -float64(p.Sprite.Size.y)/2) // Rotate axis in center
	p.Sprite.Draw(screen)

	// Debug
	cx, cy := ebiten.CursorPosition()
	prnt := fmt.Sprintf("[x:%.2f y:%.2f r:%.2f]\n[x:%d y:%d]\n[FPS: %.2f]",
		p.Sprite.Pos.x, p.Sprite.Pos.y, p.Sprite.Rot, cx, cy, ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, prnt)
}
