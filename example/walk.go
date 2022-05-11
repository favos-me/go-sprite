package main

import (
	"fmt"
	"log"

	sprite "github.com/favos-me/go-sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	windowWidth         = 320 // Width of the window
	windowHeight        = 240 // Height of the window
	girlSpeed           = windowWidth / 160
	AnimationStandLeft  = 1
	AnimationStandRight = 2
	AnimationWalkleft   = 3
	AnimationWalkRight  = 4
	AnimationStandUp    = 5
	AnimationWalkUp     = 6
	AnimationStandDown  = 7
	AnimationWalkDown   = 8
)

type game struct {
	girl *sprite.Sprite
}

func (g *game) Update() error {
	binding(g)
	// reset position if outside of the screen
	if g.girl.X > windowWidth {
		g.girl.X = 0 - g.girl.GetWidth()
	}
	if g.girl.X+g.girl.GetWidth() < 0 {
		g.girl.X = windowWidth
	}
	if g.girl.Y+g.girl.GetHeight() < 0 {
		g.girl.Y = windowHeight + 2*g.girl.GetHeight()
	}
	if g.girl.Y-2*g.girl.GetHeight() > windowHeight {
		g.girl.Y = 0 - g.girl.GetHeight()
	}

	// frame skip
	// if ebiten.IsDrawingSkipped() {
	// 	return nil
	// }

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.girl.Draw(screen)
	drawFPS(g, screen)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func main() {
	// create new sprite and load animations
	girl := sprite.NewSprite()
	girl.AddAnimation(AnimationStandRight, "gfx/som_girl_stand_right.png", 0, 1)
	girl.AddAnimation(AnimationWalkRight, "gfx/som_girl_walk_right.png", 700, 6)
	girl.AddAnimation(AnimationStandLeft, "gfx/som_girl_stand_left.png", 0, 1)
	girl.AddAnimation(AnimationWalkleft, "gfx/som_girl_walk_left.png", 700, 6)
	girl.AddAnimation(AnimationStandUp, "gfx/som_girl_stand_up.png", 0, 1)
	girl.AddAnimation(AnimationWalkUp, "gfx/som_girl_walk_up.png", 500, 4)
	girl.AddAnimation(AnimationStandDown, "gfx/som_girl_stand_down.png", 0, 1)
	girl.AddAnimation(AnimationWalkDown, "gfx/som_girl_walk_down.png", 500, 4)

	// set position and first animation
	girl.Position(windowWidth/2, windowHeight/2)
	girl.CurrentAnimation = AnimationStandRight
	girl.Start()

	g := &game{
		girl: girl,
	}
	ebiten.SetWindowTitle("girl walk")

	// infinite loop
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

// display some stuff on the screen
func drawFPS(g *game, surface *ebiten.Image) {
	ebitenutil.DebugPrint(surface,
		fmt.Sprintf("FPS:%0.1f  X:%d Y:%d %s\nLeft:%v Right:%v Up:%v Down:%v",
			ebiten.CurrentFPS(),
			int(g.girl.X), int(g.girl.Y),
			g.girl.CurrentAnimation,
			ebiten.IsKeyPressed(ebiten.KeyLeft),
			ebiten.IsKeyPressed(ebiten.KeyRight),
			ebiten.IsKeyPressed(ebiten.KeyUp),
			ebiten.IsKeyPressed(ebiten.KeyDown),
		))
}

func binding(g *game) {

	//////////////////////////// GO THE RIGHT
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {

		if ebiten.IsKeyPressed(ebiten.KeyUp) { // Right+Up
			g.girl.Direction = 45
			g.girl.Speed = girlSpeed + 1
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) { // Right+Down
			g.girl.Direction = -45
			g.girl.Speed = girlSpeed + 1
		} else { // Right
			g.girl.Direction = 0
			g.girl.Speed = girlSpeed
		}
		g.girl.CurrentAnimation = AnimationWalkRight
		g.girl.Start() // Show, Reset, Resume
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyRight) {
		g.girl.Speed = 0
		g.girl.CurrentAnimation = AnimationStandRight
	}

	//////////////////////////// GO THE LEFT
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {

		if ebiten.IsKeyPressed(ebiten.KeyUp) { // Left+Up
			g.girl.Direction = 135
			g.girl.Speed = girlSpeed + 1
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) { // Left+Down
			g.girl.Direction = 225
			g.girl.Speed = girlSpeed + 1
		} else { // Left
			g.girl.Speed = girlSpeed
			g.girl.Direction = 180
		}

		g.girl.CurrentAnimation = AnimationWalkleft
		g.girl.Start() // Show, Reset, Resume
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyLeft) {
		g.girl.Speed = 0
		g.girl.CurrentAnimation = AnimationStandLeft
	}

	//////////////////////////// GO THE TOP
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {

		if ebiten.IsKeyPressed(ebiten.KeyRight) { // Up+Right
			g.girl.Direction = 45
			g.girl.Speed = girlSpeed + 1
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) { // Up+Left
			g.girl.Direction = 135
			g.girl.Speed = girlSpeed + 1
		} else { // Up
			g.girl.Direction = 90
			g.girl.Speed = girlSpeed
		}

		g.girl.CurrentAnimation = AnimationWalkUp
		g.girl.Start() // Show, Reset, Resume
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		g.girl.Speed = 0
		g.girl.CurrentAnimation = AnimationStandUp
	}

	//////////////////////////// GO THE BOTTOM
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {

		if ebiten.IsKeyPressed(ebiten.KeyRight) { // Down+Right
			g.girl.Direction = -45
			g.girl.Speed = girlSpeed + 1
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) { // Down+Left
			g.girl.Direction = 225
			g.girl.Speed = girlSpeed + 1
		} else { // Down
			g.girl.Speed = girlSpeed
			g.girl.Direction = 270
		}

		g.girl.CurrentAnimation = AnimationWalkDown
		g.girl.Start() // Show, Reset, Resume
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		g.girl.Speed = 0
		g.girl.CurrentAnimation = AnimationStandDown
	}
}
