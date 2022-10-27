package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WIDTH  = 800
	HEIGHT = 600
)

var (
	prevDeltaTime = time.Now()
	player        = NewPlayer()
	bullets       []*Bullet
)

type Game struct{}

func (g *Game) Update() error {
	deltaTime := float64(time.Since(prevDeltaTime))
	prevDeltaTime = time.Now()

	player.Update(deltaTime)
	for _, bullet := range bullets {
		bullet.Update(deltaTime)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, bullet := range bullets {
		bullet.Draw(screen)
	}
	player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH / 2, HEIGHT / 2
	//return WIDTH, HEIGHT
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
