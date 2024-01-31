package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Char struct {
	x  int
	y  int
	vx int
	vy int
}

const (
	groundY = 360
	unit    = 10
)

func (c *Char) tryJump() {
	if c.y == groundY*unit {
		c.vy = -10 * unit
	}
}

func (c *Char) update() {
	c.x += c.vx
	c.y += c.vy

	if c.y > groundY*unit {
		c.y = groundY * unit
	}
	if c.vx > 0 {
		c.vx -= 2
	} else if c.vx < 0 {
		c.vx += 2
	}
	if c.vy < 20*unit {
		c.vy += 8
	}
}

type Player struct {
	c   *Char
	cam camera
}

func (p *Player) Update() error {
	if p.c == nil {
		p.c = &Char{x: 50 * unit, y: groundY * unit}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.c.vx = -5 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.c.vx = 5 * unit
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.c.tryJump()
	}

	p.cam.x = p.c.x

	p.c.update()
	return nil
}
