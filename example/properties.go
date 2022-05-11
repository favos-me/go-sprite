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

type game struct {
	sprites [7]*sprite.Sprite
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	for i := 0; i < len(g.sprites); i++ {
		g.sprites[i].Draw(screen)
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func main() {
	var sprites [7]*sprite.Sprite
	x := 1.0
	y := 1.0
	for i := 0; i < len(sprites); i++ {
		sprites[i] = sprite.NewSprite()
		sprites[i].CenterCoordonnates = true
		sprites[i].AddAnimation(sprite.DefaultAnimationLabel, "gfx/som_girl_stand_down.png", 1, 1)
		sprites[i].Position(windowWidth/4*x, windowHeight/4*y)
		sprites[i].Start()

		x++
		if x > 3 {
			x = 1
			y++
		}
	}

	i := 0

	sprites[i].Zoom(2, 1.5) // set the multiplier ZoomX and ZoomY
	i++

	sprites[i].Skew(30, 10) // set SkewX and SkewY in degres
	i++

	sprites[i].Rotate(45) // set Angle in degres
	i++

	sprites[i].Alpha = 0.5 // Between 0 and 1
	i++

	sprites[i].Red = 5 // Multiplier
	i++

	sprites[i].Red = 3 // become Yellow
	sprites[i].Green = 3
	i++

	sprites[i].Borders = true // Debug Borders
	i++

	g := &game{sprites: sprites}
	ebiten.SetWindowTitle("Sprite demo")
	// infinite loop
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
