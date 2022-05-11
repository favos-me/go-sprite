Go-sprite
=======

forked from https://github.com/ryosama/go-sprite and upgrade to ebiten v2, change label from string to int

A simple library for playing with sprites and animations

It use the [Ebiten](https://github.com/hajimehoshi/ebiten/v2) library for the 2D graphics engine

Install
=======

```bash
$ go get -u github.com/hajimehoshi/ebiten/v2
$ go get -u github.com/favos-me/go-sprite
```

Screenshot
===========

![Screenshot](https://github.com/favos-me/go-sprite/raw/master/screenshot1.png "Screenshot")

Quick Start
===========

```Go
import "github.com/favos-me/go-sprite"

mySprite = sprite.NewSprite()
mySprite.AddAnimation(1,	"walk_right.png", 700, 6)
mySprite.Position(WINDOW_WIDTH/2, WINDOW_HEIGHT/2)
mySprite.CurrentAnimation = 1
mySprite.Speed = 2
mySprite.Start()
```

Documentation
=============

The documentation can be found here : https://godoc.org/github.com/favos-me/go-sprite

Or export with this command

```bash
$ godoc github.com/favos-me/go-sprite
```

TODO
====

- Add a video