package main

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image         *ebiten.Image
	Options       *ebiten.DrawImageOptions
	Pos           Vec2
	Size          iVec2
	Rot           float64
	Anim          map[string]Animation
	AnimPlaying   string
	CurFrame      int
	PrevFrameTime time.Time
	FrameImage    *ebiten.Image
}

type Animation struct {
	h      int // Height slot in sprite
	Length int
	Speed  int
}

func NewSprite(image *ebiten.Image, pos Vec2, size iVec2, rot float64) *Sprite {
	sprite := new(Sprite)
	sprite.Image = image
	sprite.Options = &ebiten.DrawImageOptions{}
	sprite.Pos, sprite.Size, sprite.Rot = pos, size, rot
	sprite.Anim = make(map[string]Animation)
	sprite.CurFrame = 0
	sprite.PrevFrameTime = time.Now()

	return sprite
}

func (s *Sprite) AddAnimation(name string, h int, length int, speed int) {
	s.Anim[name] = Animation{h, length, speed}
}

func (s *Sprite) PlayAnimation(name string) {
	s.CurFrame = 0
	s.AnimPlaying = name
}

func (s *Sprite) ResetOptions() {
	s.Options.GeoM.Reset()
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	if len(s.Anim) != 0 && time.Since(s.PrevFrameTime) >= time.Duration(s.Anim[s.AnimPlaying].Speed)*time.Millisecond {
		s.PrevFrameTime = time.Now()
		s.CurFrame++
		s.CurFrame = s.CurFrame % s.Anim[s.AnimPlaying].Length
		sw, sh := s.CurFrame*s.Size.x, s.Anim[s.AnimPlaying].h*s.Size.y
		ew, eh := sw+s.Size.x, sh+s.Size.y
		s.FrameImage = s.Image.SubImage(image.Rect(sw, sh, ew, eh)).(*ebiten.Image)
	}

	// Draw
	s.Options.GeoM.Rotate(ToRadian(s.Rot))
	s.Options.GeoM.Translate(s.Pos.x, s.Pos.y)
	if len(s.Anim) != 0 {
		screen.DrawImage(s.FrameImage, s.Options)
	} else {
		screen.DrawImage(s.Image, s.Options)
	}
}
