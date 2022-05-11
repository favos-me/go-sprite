package main

import (
	"log"

	"github.com/favos-me/go-sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowWidth  = 320 // Width of the window
	windowHeight = 240 // Height of the window
)

// var (
// 	explosion1, explosion2, explosion3, explosion4, zoom1, rotate1, skew1 *sprite.Sprite
// )

type game struct {
	explosion1 *sprite.Sprite
	explosion2 *sprite.Sprite
	explosion3 *sprite.Sprite
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.explosion1.Draw(screen)
	g.explosion2.Draw(screen)
	g.explosion3.Draw(screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func main() {

	explosionDuration := 500
	// create some explosions
	explosion1 := sprite.NewSprite()
	explosion1.AddAnimation(sprite.DefaultAnimationLabel, "gfx/explosion1.png", explosionDuration, 5)
	explosion1.Position(10, windowHeight/3*2)
	explosion1.Start()

	explosion2 := sprite.NewSprite()
	explosion2.AddAnimation(sprite.DefaultAnimationLabel, "gfx/explosion2.png", explosionDuration, 7)
	explosion2.Position(windowWidth/2-24, windowHeight/3*2)
	explosion2.Start()

	explosion3 := sprite.NewSprite()
	explosion3.AddAnimation(sprite.DefaultAnimationLabel, "gfx/explosion3.png", explosionDuration, 9)
	explosion3.Position(windowWidth-10-48, windowHeight/3*2)
	explosion3.Start()

	g := &game{
		explosion1: explosion1,
		explosion2: explosion2,
		explosion3: explosion3,
	}

	ebiten.SetWindowTitle("Sprite demo")
	// infinite loop
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
