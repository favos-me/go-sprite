package main

import (
	"log"

	"github.com/favos-me/go-sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	windowWidth  = 320 // Width of the window
	windowHeight = 240 // Height of the window
)

type game struct {
	explosion *sprite.Sprite
}

func (g *game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.explosion.RunOnce(afterRunOnce)
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	// draw sprite
	g.explosion.Draw(screen)

	ebitenutil.DebugPrint(screen, "Press 'Space' to restart animation")
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func main() {
	explosion := sprite.NewSprite()
	explosion.AddAnimation(sprite.DefaultAnimationLabel, "gfx/explosion3.png", 500, 9)
	explosion.Position(windowWidth/2, windowHeight/2)
	explosion.RunOnce(afterRunOnce)

	g := &game{
		explosion: explosion,
	}
	ebiten.SetWindowTitle("Sprite demo")
	// infinite loop
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func afterRunOnce(s *sprite.Sprite) {
	print("Execute after run once\n")
}
